package execute

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/trace"
)

// QueryType is used as the trace Event name
type QueryType int

const (
	QTUnknown QueryType = iota
	QTRow
	QTRows
	QTExec
)

func (qt QueryType) String() string {
	switch qt {
	case QTRow:
		return "QueryRow"
	case QTRows:
		return "QueryRows"
	case QTExec:
		return "QueryExec"
	}
	return "UnknownQuery"
}

// TraceQuery will create a struct to hold query execution information. Any members will directly modify the returned
// pointer. Methods can be chained. End is called to log all query details to End's ctx. View End for more details.
// If the ctx passed in here is not recording, a nil TraceQueryInfo is returned. All methods are nil safe.
func TraceQuery(ctx context.Context, qType QueryType, query string, args []interface{}) *TraceQueryInfo {
	// If we are not recording, why bother doing all this work?
	if !trace.SpanFromContext(ctx).IsRecording() {
		return nil
	}
	return &TraceQueryInfo{
		qType: qType,
		query: query,
		args:  args,
	}
}

// TraceQueryInfo holds information about a single query execution. The only thing it does not (and should not)
// track is the time it takes to do anything. That should be calculated depending on what information people want from
// their traces / metrics.
type TraceQueryInfo struct {
	finish sync.Once
	qType  QueryType
	count  int64
	query  string
	args   []interface{}
	err    error
}

// Err will store the error on Ends trace. Can be called multiple times to overwrite.
func (q *TraceQueryInfo) Err(err error) *TraceQueryInfo {
	if q != nil {
		q.err = err
	}
	return q
}

// RowCount will store the number of rows on Ends trace. Can be called multiple times to overwrite.
func (q *TraceQueryInfo) RowCount(rowCount int64) *TraceQueryInfo {
	if q != nil {
		q.count = rowCount
	}
	return q
}

// End will log the trace info to the span within ctx. If the TraceQueryInfo is valid, it appends all information it
// has collected into an event named by the QueryType. If an error occurred, it is logged after the QueryType event.
func (q *TraceQueryInfo) End(ctx context.Context) {
	if q == nil {
		return // Nothing to do
	}

	// We actually have data to trace. Only do this all finish!
	q.finish.Do(func() {
		span := trace.SpanFromContext(ctx)
		// If we somehow were recording before, and now we are not, there is nowhere for us to put our data!
		// Why bother serializing everything? Just exit early.
		if !span.IsRecording() {
			return
		}

		sqlQuery := core.KeyValue{
			Key:   "SQL",
			Value: core.String(q.query),
		}
		rowCountAttr := core.KeyValue{
			Key:   "RowCount",
			Value: core.Int64(q.count),
		}

		// Adding +2 so it all can be logged in one call, for:
		//  (1) the SQLExp query string Field itself
		//  (2) the rowsRetrieved int64 Field itself
		fields := make([]core.KeyValue, 0, len(q.args)+2)
		fields = append(appendArgs(fields, q.args), sqlQuery, rowCountAttr)
		span.AddEvent(ctx, q.qType.String(), fields...)

		// If we had an error, log it!
		if q.err != nil {
			span.RecordError(ctx, q.err)
		}
	})
}

// traceErr records the error on the span and returns the same error. NOTE: Only to be used in places where there is
// not no TraceQueryInfo, like building SQL using a dialect, transaction errors, etc.
func traceErr(ctx context.Context, errToRecord error) error {
	span := trace.SpanFromContext(ctx)
	span.RecordError(ctx, errToRecord)
	return errToRecord
}

// appendArgs will convert each arg to an OpenTelemetry core.Values and append the args to attrs,
// returning the resulting slice.
func appendArgs(attrs []core.KeyValue, args []interface{}) []core.KeyValue {
	for idx, arg := range args {
		logName := fmt.Sprintf("Arg[%d]", idx)
		attrs = append(attrs, core.KeyValue{
			Key:   core.Key(logName),
			Value: toTraceValue(arg),
		})
	}
	return attrs
}

//revive:disable:cyclomatic - toTraceValue has a high complexity score of 16, and cannot find a great way to break
// this up into smaller chunks. 🤷‍♂️ Shrug.

func toTraceValue(arg interface{}) core.Value {
	var v core.Value
	switch casted := arg.(type) {
	case []byte:
		v = core.String(fmt.Sprintf("%x", casted))
	case [][]byte:
		values := make([]string, len(casted))
		for vi, castedVal := range casted {
			values[vi] = fmt.Sprintf("%x", castedVal)
		}
		v = core.String(strings.Join(values, ","))
	case string:
		v = core.String(casted)
	case bool:
		v = core.Bool(casted)
	case int8:
		v = core.Int(int(casted))
	case int16:
		v = core.Int(int(casted))
	case int:
		v = core.Int(casted)
	case int32:
		v = core.Int32(casted)
	case int64:
		v = core.Int64(casted)
	case uint8:
		v = core.Uint(uint(casted))
	case uint16:
		v = core.Uint(uint(casted))
	case uint32:
		v = core.Uint32(casted)
	case uint64:
		v = core.Uint64(casted)
	case float32:
		v = core.Float32(casted)
	case float64:
		v = core.Float64(casted)
	case time.Time:
		v = core.String(casted.String())
	case nil:
		v = core.String("nil")
	default:
		v = core.String(fmt.Sprintf("%s", arg))
	}
	return v
}

//revive:enable:cyclomatic

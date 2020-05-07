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

// queryType is used as the trace Event name
type queryType int

const (
	queryTypeUnknown queryType = iota
	queryTypeRow
	queryTypeRows
	queryTypeExec
)

func (qt queryType) String() string {
	switch qt {
	case queryTypeUnknown:
		return "UnknownQuery"
	case queryTypeRow:
		return "QueryRow"
	case queryTypeRows:
		return "QueryRows"
	case queryTypeExec:
		return "QueryExec"
	}
	return queryTypeUnknown.String()
}

// traceQuery will create a struct to hold query execution information. Any members will directly modify the returned
// pointer. Methods can be chained. End is called to log all query details to End's ctx. View End for more details.
// If the ctx passed in here is not recording, a nil traceQueryInfo is returned. All methods are nil safe.
func traceQuery(ctx context.Context, qType queryType, query string, args []interface{}) *traceQueryInfo {
	span := trace.SpanFromContext(ctx)

	// If we are not recording, why bother doing all this work?
	if !span.IsRecording() {
		return nil
	}
	return &traceQueryInfo{
		qType: qType,
		query: query,
		args:  args,
	}
}

// traceQueryInfo holds information about a single query execution. The only thing it does not (and should not)
// track is the time it takes to do anything. That should be calculated depending on what information people want from
// their traces / metrics.
type traceQueryInfo struct {
	qType queryType
	count int64
	query string
	args  []interface{}
	err   error
	once  sync.Once
}

// Err will store the error on Ends trace. Can be called multiple times to overwrite.
func (q *traceQueryInfo) Err(err error) *traceQueryInfo {
	if q != nil {
		q.err = err
	}
	return q
}

// RowCount will store the number of rows on Ends trace. Can be called multiple times to overwrite.
func (q *traceQueryInfo) RowCount(rowCount int64) *traceQueryInfo {
	if q != nil {
		q.count = rowCount
	}
	return q
}

// End will log the trace info to the span within ctx. If the traceQueryInfo is nil, the span within ctx has an
// queryTypeUnknown event is attempted to be logged. If the traceQueryInfo is valid, it appends all information it
// has collected into an event named by the queryType. If an error occurred, it is logged after the queryType event.
func (q *traceQueryInfo) End(ctx context.Context) {
	span := trace.SpanFromContext(ctx)

	// Still add the event if they are doing traces in this context. They might have not created this
	// cursor with span recording, but this context might, do it just to register a query occurred.
	if q == nil {
		span.AddEvent(ctx, queryTypeUnknown.String())
		return
	}

	// We actually have data to trace. Only do this all once!
	q.once.Do(func() {
		// Adding +2 so it all can be logged in one call, for:
		//  (1) the SQLExp query string Field itself
		//  (2) the rowsRetrieved int64 Field itself
		fields := make([]core.KeyValue, 0, len(q.args)+2)
		fields = append(appendArgs(fields, q.args), sqlStrToAttr(q.query), rowCountToAttr(q.count))
		span.AddEvent(ctx, q.qType.String(), fields...)

		// If we had an error, log it!
		if q.err != nil {
			span.RecordError(ctx, q.err)
		}
	})
}

//revive:disable:cyclomatic - TraceQuery has a high complexity score of 16, and ScanCursor cannot find a great
// way to break this up into smaller chunks. 🤷‍♂️ Shrug.

// appendArgs will convert each arg to an OpenTelemetry core.Values and append the args to attrs,
// returning the resulting slice.
func appendArgs(attrs []core.KeyValue, args []interface{}) []core.KeyValue {
	for idx, arg := range args {
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
			v = core.String(casted.UTC().String())
		default:
			v = core.String(fmt.Sprintf("%s", arg))
		}

		logName := fmt.Sprintf("Arg[%d]", idx)
		attrs = append(attrs, core.KeyValue{
			Key:   core.Key(logName),
			Value: v,
		})
	}
	return attrs
}

//revive:enable:cyclomatic

func sqlStrToAttr(query string) core.KeyValue {
	return core.KeyValue{
		Key:   "SQLExp",
		Value: core.String(query),
	}
}

func rowCountToAttr(count int64) core.KeyValue {
	return core.KeyValue{
		Key:   "RowCount",
		Value: core.Int64(count),
	}
}

// traceErr records the error on the span and returns the same error.
func traceErr(ctx context.Context, errToRecord error) error {
	span := trace.SpanFromContext(ctx)
	span.RecordError(ctx, errToRecord)
	return errToRecord
}

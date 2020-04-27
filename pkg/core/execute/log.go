package execute

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/trace"
)

// TODO fix this comment
// TraceQuery will log the original query string to the span as "sql", and all arguments as "arg[idx]" where idx
// is the numerical order of the argument.
//
// NOTE: This will not properly print everything like how it would be used by SQL. For example, the SQL driver might
// replace time.Time with a int64 before sending it to the database.
func (e execSQL) TraceQuery(ctx context.Context, rowCount uint64, query string, args []interface{}) {
	span := trace.SpanFromContext(ctx)

	// If we are not recording, why bother doing all this work?
	if !span.IsRecording() {
		return
	}


	// Adding +2 so it all can be logged in one call, for:
	//  (1) the SQL query string Field itself
	//  (2) the rowsRetrieved int64 Field itself
	fields := make([]core.KeyValue, 0, len(args)+2)
	argsToAttrs(fields, args)
	fields = append(fields, sqlStrToAttr(query), rowCountToAttr(rowCount))

	span.AddEvent(ctx, "query", fields...)
}

func (e execSQL) TraceQueryStart(ctx context.Context, query string, args[]interface{}) {
	span := trace.SpanFromContext(ctx)

	// If we are not recording, why bother doing all this work?
	if !span.IsRecording() {
		return
	}

	// Adding +1 so it all can be logged in one call the SQL query string Field itself
	fields := make([]core.KeyValue, 0, len(args)+1)
	argsToAttrs(fields, args)
	fields = append(fields, sqlStrToAttr(query))

	span.AddEvent(ctx, "query_start", fields...)
}

func (e execSQL) TraceQueryEnd(ctx context.Context, rowCount uint64) {
	span := trace.SpanFromContext(ctx)

	// If we are not recording, why bother doing all this work?
	if !span.IsRecording() {
		return
	}

	span.AddEvent(ctx, "query_end", rowCountToAttr(rowCount))
}

//revive:disable:cyclomatic - TraceQuery has a high complexity score of 16, and ScanCursor cannot find a great
// way to break this up into smaller chunks. 🤷‍♂️ Shrug.

func argsToAttrs(attrs []core.KeyValue, args[]interface{}) {
	for idx, arg := range args {
		var v core.Value
		switch casted := arg.(type) {
		case string:
			v = core.String(casted)
		case bool:
			v = core.Bool(casted)
		case int:
			v = core.Int(casted)
		case int32:
			v = core.Int32(casted)
		case int64:
			v = core.Int64(casted)
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
		case []byte:
			v = core.String(fmt.Sprintf("%x", casted))
		case [][]byte:
			values := make([]string, len(casted))
			for vi, castedVal := range casted {
				values[vi] = fmt.Sprintf("%x", castedVal)
			}
			v = core.String(strings.Join(values, ","))
		default:
			v = core.String(fmt.Sprintf("%s", casted))
		}

		logName := fmt.Sprintf("arg[%d]", idx)
		attrs = append(attrs, core.KeyValue{
			Key:   core.Key(logName),
			Value: v,
		})
	}
}
//revive:enable:cyclomatic

func sqlStrToAttr(query string) core.KeyValue {
	return core.KeyValue{
		Key: "SQL",
		Value: core.String(query),
	}
}

func rowCountToAttr(count uint64) core.KeyValue {
	return core.KeyValue{
		Key: "rows",
		Value: core.Uint64(count),
	}
}

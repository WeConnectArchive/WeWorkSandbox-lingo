package execute

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/trace"
)

//revive:disable:cyclomatic - LogSQLAndArgs has a high complexity score of 16, and I cannot find a great
// way to break this up into smaller chunks. 🤷‍♂️ Shrug.

// LogSQLAndArgs will log the original query string to the span as "sql", and all arguments as "arg[idx]" where idx
// is the numerical order of the argument.
//
// NOTE: This will not properly print everything like how it would be used by SQL. For example, the SQL driver might
// replace time.Time with a int64 before sending it to the database.
func (e execSQL) LogSQLAndArgs(ctx context.Context, s trace.Span, query string, args ...interface{}) {
	if !s.IsRecording() {
		return
	}

	// +1 for the SQL string Field itself so it all can be logged in one call
	fields := make([]core.KeyValue, len(args)+1)
	for i, arg := range args {
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

		logName := fmt.Sprintf("arg[%d]", i)
		fields = append(fields, core.KeyValue{
			Key:   core.Key(logName),
			Value: v,
		})
	}
	fields[len(args)] = core.KeyValue{
		Key: "SQL",
		Value: core.String(query),
	}
	s.AddEvent(ctx, "query", fields...)
}
//revive:enable:cyclomatic


package execute

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/metric"
	"go.opentelemetry.io/otel/api/trace"

	"github.com/weworksandbox/lingo/internal"
	"github.com/weworksandbox/lingo/pkg/core"
)

type SQL interface {
	QueryRow(pCtx context.Context, exp core.Expression, valuePtrs... interface{}) error
}

func NewSQL(db *sql.DB, d core.Dialect) SQL {
	const t = "/pkg/core/execute.SQL"
	return execSQL{
		db: db,
		d: d,
		tr: global.Tracer(internal.ModuleName + t),
		met: global.Meter(internal.ModuleName + t),
	}
}

type execSQL struct {
	db *sql.DB
	d core.Dialect

	tr trace.Tracer
	met metric.Meter
}

func (e execSQL) QueryRow(pCtx context.Context, exp core.Expression, valuePtrs... interface{}) error {
	ctx, span := e.tr.Start(pCtx, "QueryRow")
	defer span.End()

	for idx, value := range valuePtrs {
		v := reflect.ValueOf(value)
		if !isPtr(v) {
			return fmt.Errorf("value at %d is not a pointer but was %s", idx, v.String())
		}
	}

	lSQL, err := exp.GetSQL(e.d)
	if err != nil {
		span.RecordError(ctx, err)
		return err
	}

	if span.IsRecording() {
		e.LogSQLAndArgs(ctx, span, lSQL.String(), lSQL.Values()...)
	}

	row := e.db.QueryRowContext(ctx, lSQL.String(), lSQL.Values()...)
	err = row.Scan(valuePtrs...)
	if err != nil {
		span.RecordError(ctx, err)
		return err
	}
	return nil
}

func isPtr(v reflect.Value) bool {
	return v.Type().Kind() == reflect.Ptr
}

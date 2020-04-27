package execute

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"sync/atomic"

	otelcore "go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/metric"
	"go.opentelemetry.io/otel/api/trace"

	"github.com/weworksandbox/lingo/internal"
	"github.com/weworksandbox/lingo/pkg/core"
)

type SQL interface {
	Query(pCtx context.Context, exp core.Expression) (RowScanner, error)
	QueryRow(pCtx context.Context, exp core.Expression, valuePtrs ...interface{}) error
}

func NewSQL(db *sql.DB, d core.Dialect) SQL {
	const t = "/pkg/core/execute.SQL"
	return execSQL{
		db:  db,
		d:   d,
		tr:  global.Tracer(internal.ModuleName + t),
		met: global.Meter(internal.ModuleName + t),
	}
}

type execSQL struct {
	db *sql.DB
	d  core.Dialect

	tr  trace.Tracer
	met metric.Meter
}

func (e execSQL) QueryRow(pCtx context.Context, exp core.Expression, queryIntoPtrs ...interface{}) error {
	ctx, span := e.tr.Start(pCtx, "QueryRow")
	defer span.End()

	if err := e.validatePtrs(queryIntoPtrs); err != nil {
		return err
	}

	lSQL, err := exp.GetSQL(e.d)
	if err != nil {
		return err
	}

	tSQL := lSQL.String()
	sVals := lSQL.Values()

	var rowCount uint64
	defer func() {
		e.TraceQuery(ctx, rowCount, tSQL, sVals...)
	}()

	err = span.Tracer().WithSpan(ctx, "sql.DB.QueryRowContext", func(ctx context.Context) error {
		return e.db.QueryRowContext(ctx, tSQL, sVals...).Scan(queryIntoPtrs...)
	})
	if err != nil {
		return err
	}

	// If we got here there was at least 1 row read. If not, err would be sql.ErrNoRows.
	rowCount = 1
	return nil
}

type RowScanner struct {
	rowCount uint64
	rows     *sql.Rows
	err      error
}

func (r RowScanner) ScanRow(vals ...interface{}) bool {
	if !r.rows.Next() {
		r.err = getDoneErr(r.rows)
		return false
	}

	if scanErr := r.rows.Scan(vals...); scanErr != nil {
		r.err = scanErr
		return false
	}

	atomic.AddUint64(&r.rowCount, 1)
	return true
}

func (r RowScanner) Err() error {
	return getDoneErr(r.rows)
}

func (e execSQL) Query(pCtx context.Context, exp core.Expression) (RowScanner, error) {
	ctx, span := e.tr.Start(pCtx, "Query")
	defer span.End()

	lSQL, err := exp.GetSQL(e.d)
	if err != nil {
		return RowScanner{}, err
	}

	tSQL := lSQL.String()
	sVals := lSQL.Values()

	e.TraceQueryStart(ctx, tSQL, sVals)
	rows, err := e.db.QueryContext(ctx, tSQL, sVals...)
	if err != nil {
		return RowScanner{}, err
	}

	return RowScanner{
		rows: rows,
	}, nil
}

type Scan = func(vals ...interface{}) error

type ScanOrError struct {

}

func (e execSQL) QueryFunc(pCtx context.Context, exp core.Expression) (Scan, error) {
	ctx, span := e.tr.Start(pCtx, "Query")
	defer span.End()

	lSQL, err := exp.GetSQL(e.d)
	if err != nil {
		return nil, err
	}

	tSQL := lSQL.String()
	sVals := lSQL.Values()

	e.TraceQueryStart(ctx, tSQL, sVals)
	rows, err := e.db.QueryContext(ctx, tSQL, sVals...)
	if err != nil {
		return nil, err
	}

	f := func(vals ...interface{}) error {
		defer func() {
			// Allowed to close more than once, ensure all resources are released
			_ = rows.Close()
		}()

		return nil
	}
	return f, nil
}


func getDoneErr(rows *sql.Rows) error {
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rCloseErr := rows.Close()
	if rCloseErr != nil {
		return rCloseErr
	}

	// Rows.Err will report the last error encountered by Rows.ScanFunc.
	if rErr := rows.Err(); rErr != nil {
		return rErr
	}
	return nil
}

func (e execSQL) validatePtrs(valuePtrs []interface{}) error {
	for idx, value := range valuePtrs {
		v := reflect.ValueOf(value)
		if !isPtr(v) {
			return fmt.Errorf("value at %d is not a pointer but was %s", idx, v.String())
		}
	}
	return nil
}

func isPtr(v reflect.Value) bool {
	return v.Type().Kind() == reflect.Ptr
}

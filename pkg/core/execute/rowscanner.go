package execute

import (
	"context"
	"database/sql"
	"fmt"
)

type RowScanner interface {
	ScanRow(vals ...interface{}) bool
	Err(ctx context.Context) error
	Close(ctx context.Context)
}

type rowScanner struct {
	rowCount int64
	err      error

	rows       *sql.Rows
	queryTrace *TraceQueryInfo
}

// ScanRow will check to see if there is another row to scan, calls rows.Scan, and returns true if successful. If
// false, use Err to determine if an error occurred or not. If not, that was the end of the result set.
func (r *rowScanner) ScanRow(vals ...interface{}) bool {
	if !r.rows.Next() {
		// Cannot recover from not having anything else.
		r.closeAndSetErr(r.rows.Err())
		return false
	}

	if scanErr := r.rows.Scan(vals...); scanErr != nil {
		r.closeAndSetErr(scanErr)
		return false
	}

	r.rowCount++
	return true
}

// Err will closeAndSetErr the result set (if not closed already) and return whatever error occurred.
func (r *rowScanner) Err(ctx context.Context) error {
	r.traceQueryEnd(ctx)
	r.closeAndSetErr(nil)
	return r.err
}

// Close calls close and sets the error if one occurs. This is really only used to defer
// to ensure the db.Rows are closed. Multiple calls allowed.
func (r *rowScanner) Close(ctx context.Context) {
	r.traceQueryEnd(ctx)
	r.closeAndSetErr(nil)
}

func (r *rowScanner) traceQueryEnd(ctx context.Context) {
	r.queryTrace.RowCount(r.rowCount).End(ctx)
}

func (r *rowScanner) closeAndSetErr(otherErr error) {
	closeErr := r.rows.Close()

	var result error
	switch {
	// Do not replace an existing error
	case r.err != nil:
		result = r.err
	case closeErr != nil && otherErr != nil:
		result = fmt.Errorf("%s: %w", closeErr, otherErr)
	case otherErr != nil:
		result = otherErr
	case closeErr != nil:
		result = closeErr
	}
	r.err = result
}

package execute

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type ExecSQLInTx = func(ctx context.Context, s SQLQuery) error

type SQL interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (TxSQL, error)
	InTx(ctx context.Context, opts *sql.TxOptions, todo ExecSQLInTx) error
	SQLQuery
}

type TxSQL interface {
	CommitOrRollback(ctx context.Context, err error) error
	Rollback(ctx context.Context, err error) error
	SQLQuery
}

type SQLQuery interface {
	Exec(ctx context.Context, tSQL string, args ...interface{}) (sql.Result, error)
	Query(ctx context.Context, tSQL string, args ...interface{}) (RowScanner, error)
	QueryRow(ctx context.Context, tSQL string, args []interface{}, valuePtrs ...interface{}) error
}

func NewSQL(db TxDBQuery) SQL {
	return sqlDBExec{
		db: db,
	}
}

type sqlDBExec struct {
	db TxDBQuery
}

func (s sqlDBExec) BeginTx(ctx context.Context, opts *sql.TxOptions) (TxSQL, error) {
	tx, err := s.db.BeginTx(ctx, opts)
	if err != nil {
		return nil, traceErr(ctx, err)
	}
	return NewSQLTx(tx), nil
}

func (s sqlDBExec) Exec(ctx context.Context, tSQL string, sVals ...interface{}) (sql.Result, error) {
	return exec(ctx, s.db, tSQL, sVals)
}

func (s sqlDBExec) QueryRow(ctx context.Context, tSQL string, sVals []interface{}, queryIntoPtrs ...interface{}) error {
	return queryRow(ctx, s.db, tSQL, sVals, queryIntoPtrs)
}

func (s sqlDBExec) Query(ctx context.Context, tSQL string, sVals ...interface{}) (RowScanner, error) {
	return query(ctx, s.db, tSQL, sVals)
}

func (s sqlDBExec) InTx(ctx context.Context, opts *sql.TxOptions, execThis ExecSQLInTx) (err error) {
	var txSQL TxSQL

	panicked := true
	defer func() {
		r := recover()
		if r != nil || panicked { // Workaround for if someone throws `nil`
			panicked = true // Set it to true regardless because we checked r, or panicked was already true.
			err = fmt.Errorf("panicked with %v", r)
		}
		if txSQL != nil {
			err = txSQL.CommitOrRollback(ctx, err)
		}

		if panicked {
			panic(r) // Throw the same thing we caught. Do not change the type, value, etc.
		}
	}()

	txSQL, err = s.BeginTx(ctx, opts)
	if err != nil {
		panicked = false // Normal error condition short circuit, no panic happened
		return err       // Already Traced
	}
	err = execThis(ctx, txSQL)
	panicked = false
	return
}

type ActiveDBTx interface {
	DBQuery
	driver.Tx
}

func NewSQLTx(db ActiveDBTx) TxSQL {
	return sqlActiveTxExec{
		db: db,
	}
}

type sqlActiveTxExec struct {
	sqlDBExec
	db ActiveDBTx
}

func (s sqlActiveTxExec) CommitOrRollback(ctx context.Context, err error) error {
	// If there was an error, we need to rollback
	if err != nil {
		return s.Rollback(ctx, err)
	}

	// err is nil; if Commit returns error update err
	if err = s.db.Commit(); err != nil {
		return traceErr(ctx, err)
	}
	return nil
}

func (s sqlActiveTxExec) Rollback(ctx context.Context, err error) error {
	if rollbackErr := s.db.Rollback(); rollbackErr != nil {
		if err != nil {
			err = fmt.Errorf("%s: %w", rollbackErr, err)
		} else {
			err = rollbackErr
		}
		return traceErr(ctx, err)
	}
	return err
}

func (s sqlActiveTxExec) Exec(ctx context.Context, tSQL string, sVals ...interface{}) (sql.Result, error) {
	return exec(ctx, s.db, tSQL, sVals)
}

func (s sqlActiveTxExec) QueryRow(
	ctx context.Context, tSQL string, sVals []interface{}, queryIntoPtrs ...interface{},
) error {
	return queryRow(ctx, s.db, tSQL, sVals, queryIntoPtrs)
}

func (s sqlActiveTxExec) Query(ctx context.Context, tSQL string, sVals ...interface{}) (RowScanner, error) {
	return query(ctx, s.db, tSQL, sVals)
}

func exec(ctx context.Context, db DBQuery, tSQL string, sVals []interface{}) (result sql.Result, err error) {
	var rowCount int64
	queryTrace := traceOrNil(ctx, QTExec, tSQL, sVals)
	defer func() {
		queryTrace.RowCount(rowCount).Err(err).End(ctx)
	}()

	result, err = db.ExecContext(ctx, tSQL, sVals...)
	if err != nil {
		return nil, err
	}

	// If we have a rowsAffected, trace / log it.
	rowCount, _ = result.RowsAffected()

	// Not going to log the LastInsertID... potential security risk.
	return result, nil
}

func queryRow(
	ctx context.Context, db DBQuery, tSQL string, sVals []interface{}, queryIntoPtrs []interface{},
) (err error) {
	var rowCount int64
	queryTrace := traceOrNil(ctx, QTRow, tSQL, sVals)
	defer func() {
		queryTrace.RowCount(rowCount).Err(err).End(ctx)
	}()

	if err = db.QueryRowContext(ctx, tSQL, sVals...).Scan(queryIntoPtrs...); err != nil {
		return err
	}

	// If we got here there was at least 1 row read. If not, err would be sql.ErrNoRows.
	rowCount = 1
	return nil
}

func query(ctx context.Context, db DBQuery, tSQL string, sVals []interface{}) (RowScanner, error) {
	queryTrace := traceOrNil(ctx, QTRows, tSQL, sVals)

	var err error
	defer func() {
		// Only End the trace if we have an error, if we dont eror,
		// the RowScanner will do that work.
		if err != nil {
			queryTrace.Err(err).End(ctx)
		}
	}()

	var rows *sql.Rows
	rows, err = db.QueryContext(ctx, tSQL, sVals...)
	if err != nil {
		queryTrace.Err(err).End(ctx)
		return nil, err
	}

	return &rowScanner{
		rows:       rows,
		queryTrace: queryTrace,
	}, nil
}

func traceOrNil(ctx context.Context, qType QueryType, tSQL string, sVals []interface{}) *TraceQueryInfo {
	return TraceQuery(ctx, qType, tSQL, sVals)
}

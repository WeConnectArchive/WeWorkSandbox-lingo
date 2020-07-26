package parse

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	// Include MySQL driver in order to connect to it in NewMySQL
	_ "github.com/go-sql-driver/mysql"

	"github.com/weworksandbox/lingo/internal/generate"
)

func NewMySQL(ctx context.Context, dsn string) (*MySQL, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return &MySQL{db: db}, nil
}

type MySQL struct {
	db *sql.DB
}

func (MySQL) DBTypesToPaths() map[string]generate.PathPackageToType {
	const pkgCorePath = "github.com/weworksandbox/lingo/expr/path"
	// TODO - Need to do further changes to Paths. Right now, every Path can have nullable operations against it.
	//  We may want to create a `Int64NullPath` vs `Int64Path` for example. In that case, `Int64NullPath` just extends
	//  and adds the nullable methods? https://github.com/go-sql-driver/mysql/blob/master/fields.go
	// Note: For `decimal`, we create our own, but there is no 'decimal' type in Go
	// besides `math/big/decimal.go` which is binary anyway...
	return map[string]generate.PathPackageToType{
		"BIGINT":    {pkgCorePath, "Int64"},
		"BINARY":    {pkgCorePath, "Binary"},
		"CHAR":      {pkgCorePath, "String"},
		"DATETIME":  {pkgCorePath, "Time"},
		"DECIMAL":   {pkgCorePath, "Binary"}, // See note above.
		"DOUBLE":    {pkgCorePath, "Float64"},
		"FLOAT":     {pkgCorePath, "Float32"},
		"INT":       {pkgCorePath, "Int"},
		"JSON":      {pkgCorePath, "JSON"},
		"MEDIUMINT": {pkgCorePath, "Int32"},
		"SMALLINT":  {pkgCorePath, "Int16"},
		"TEXT":      {pkgCorePath, "String"},
		"TINYINT":   {pkgCorePath, "Int8"},
		"TIMESTAMP": {pkgCorePath, "Time"},
		"VARCHAR":   {pkgCorePath, "String"},
	}
}

func (m MySQL) Tables(ctx context.Context, schema string) ([]string, error) {
	const selectQuery = "SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = ?"
	sqlStmt, prepareErr := m.db.PrepareContext(ctx, selectQuery)
	if prepareErr != nil {
		return nil, prepareErr
	}

	defer func() {
		if closeErr := sqlStmt.Close(); closeErr != nil {
			log.Printf("unable to close `findTables` query: %v", closeErr)
		}
	}()

	rows, queryErr := sqlStmt.QueryContext(ctx, schema)
	if queryErr != nil {
		return nil, queryErr
	}

	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("unable to close `rows` during `findTables` query: %v", closeErr)
		}
	}()

	var tables []string
	for rows.Next() {
		var tableName string
		if scanErr := rows.Scan(&tableName); scanErr != nil {
			return nil, scanErr
		}
		tables = append(tables, tableName)
	}
	return tables, nil
}

func (m MySQL) Columns(ctx context.Context, schema, table string) ([]generate.Column, error) {
	sqlStr := fmt.Sprintf("SELECT * FROM %s.%s LIMIT 0", schema, table)
	sqlStmt, prepareErr := m.db.PrepareContext(ctx, sqlStr)
	if prepareErr != nil {
		return nil, prepareErr
	}

	defer func() {
		if closeErr := sqlStmt.Close(); closeErr != nil {
			log.Printf("unable to close `findColumns` query: %v", closeErr)
		}
	}()

	rows, queryErr := sqlStmt.QueryContext(ctx)
	if queryErr != nil {
		return nil, queryErr
	}

	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			log.Printf("unable to close `rows` during `findColumns` query: %v", closeErr)
		}
	}()

	columnTypes, typesErr := rows.ColumnTypes()
	if typesErr != nil {
		return nil, typesErr
	}

	var columns = make([]generate.Column, 0, len(columnTypes))
	for _, col := range columnTypes {
		var columnInfo = Column{
			table:      table,
			columnType: col,
		}
		columns = append(columns, columnInfo)
	}
	return columns, nil
}

package parse

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	// Include MySQL driver in order to connect to it in NewMySQL
	_ "github.com/go-sql-driver/mysql"

	"github.com/weworksandbox/lingo/dialect"
	"github.com/weworksandbox/lingo/internal/generator"
)

func NewMySQL(dsn string) (*MySQL, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &MySQL{db: db}, nil
}

type MySQL struct {
	dialect.MySQL
	db *sql.DB
}

func (MySQL) DBTypesToPaths() map[string]generator.PathPackageToType {
	const pkgCorePath = "github.com/weworksandbox/lingo/expr/path"
	// TODO - Need to do further changes to Paths. Right now, every Path can have nullable operations against it.
	//  We may want to create a `Int64NullPath` vs `Int64Path` for example. In that case, `Int64NullPath` just extends
	//  and adds the nullable methods? https://github.com/go-sql-driver/mysql/blob/master/fields.go
	// Note: For `decimal`, we create our own, but there is no 'decimal' type in Go
	// besides `math/big/decimal.go` which is binary anyway...
	return map[string]generator.PathPackageToType{
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

func (m MySQL) Tables(ctx context.Context, schema string) (<-chan string, <-chan error) {
	var tables = make(chan string)
	var errors = make(chan error)

	go func() {
		defer close(tables)
		defer close(errors)

		const selectQuery = "SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = ?"
		sqlStmt, prepareErr := m.db.PrepareContext(ctx, selectQuery)
		if prepareErr != nil {
			errors <- prepareErr
			return
		}

		defer func() {
			if closeErr := sqlStmt.Close(); closeErr != nil {
				log.Printf("unable to close `findTables` query: %v", closeErr)
			}
		}()

		rows, queryErr := sqlStmt.QueryContext(ctx, schema)
		if queryErr != nil {
			errors <- queryErr
			return
		}

		defer func() {
			if closeErr := rows.Close(); closeErr != nil {
				log.Printf("unable to close `rows` during `findTables` query: %v", closeErr)
			}
		}()

		for rows.Next() {
			var tableName string
			if scanErr := rows.Scan(&tableName); scanErr != nil {
				errors <- scanErr
				return
			}
			tables <- tableName
		}
	}()
	return tables, errors
}

func (m MySQL) Columns(ctx context.Context, schema, table string) (<-chan generator.Column, <-chan error) {
	var columns = make(chan generator.Column)
	var errors = make(chan error)

	go func() {
		defer close(columns)
		defer close(errors)

		sqlStr := fmt.Sprintf("SELECT * FROM %s.%s LIMIT 0", schema, table)
		sqlStmt, prepareErr := m.db.PrepareContext(ctx, sqlStr)
		if prepareErr != nil {
			errors <- prepareErr
			return
		}

		defer func() {
			if closeErr := sqlStmt.Close(); closeErr != nil {
				log.Printf("unable to close `findColumns` query: %v", closeErr)
			}
		}()

		rows, queryErr := sqlStmt.QueryContext(ctx)
		if queryErr != nil {
			errors <- queryErr
			return
		}

		defer func() {
			if closeErr := rows.Close(); closeErr != nil {
				log.Printf("unable to close `rows` during `findColumns` query: %v", closeErr)
			}
		}()

		columnTypes, typesErr := rows.ColumnTypes()
		if typesErr != nil {
			errors <- typesErr
			return
		}

		for _, col := range columnTypes {
			var columnInfo = Column{
				table:      table,
				columnType: col,
			}
			columns <- &columnInfo
		}
	}()
	return columns, errors
}

func (m MySQL) ForeignKeys(ctx context.Context, schema, table string) (<-chan generator.ForeignKey, <-chan error) {
	var foreignKeys = make(chan generator.ForeignKey)
	var errors = make(chan error)

	go func() {
		defer close(foreignKeys)
		defer close(errors)

		sqlStr := `
			SELECT 
				cu.CONSTRAINT_NAME, cu.COLUMN_NAME, cu.ORDINAL_POSITION, 
				cu.REFERENCED_TABLE_SCHEMA, cu.REFERENCED_TABLE_NAME, cu.REFERENCED_COLUMN_NAME
			FROM information_schema.KEY_COLUMN_USAGE cu
			LEFT JOIN information_schema.TABLE_CONSTRAINTS tc on cu.CONSTRAINT_NAME = tc.CONSTRAINT_NAME
			WHERE cu.TABLE_SCHEMA = ? AND cu.TABLE_NAME = ? AND tc.CONSTRAINT_TYPE = 'FOREIGN KEY'
			ORDER BY cu.ORDINAL_POSITION;
		`

		sqlStmt, prepareErr := m.db.PrepareContext(ctx, sqlStr)
		if prepareErr != nil {
			errors <- prepareErr
			return
		}
		defer func() {
			if closeErr := sqlStmt.Close(); closeErr != nil {
				log.Printf("unable to close `findForeignConstraints` query: %v", closeErr)
			}
		}()

		rows, queryErr := sqlStmt.QueryContext(ctx, schema, table)
		if queryErr != nil {
			errors <- queryErr
			return
		}

		defer func() {
			if closeErr := rows.Close(); closeErr != nil {
				log.Printf("unable to close `rows` during `findForeignConstraints` query: %v", closeErr)
			}
		}()

		var fKey *ForeignKey
		for rows.Next() {
			// Prob refactor this out?
			var constraintName string
			var columnName string
			var ordinalPosition int
			var referencedTableSchema string
			var referencedTableName string
			var referencedTableColumnName string

			scanErr := rows.Scan(
				&constraintName,
				&columnName,
				&ordinalPosition,
				&referencedTableSchema,
				&referencedTableName,
				&referencedTableColumnName)
			if scanErr != nil {
				errors <- scanErr
			}

			if fKey == nil {
				fKey = &ForeignKey{
					name: constraintName,
				}
			} else if fKey.Name() != constraintName {
				foreignKeys <- fKey

				fKey = &ForeignKey{
					name: constraintName,
				}
			}
			// Append Column
		}

		if fKey != nil {
			foreignKeys <- fKey
		}
	}()
	return foreignKeys, errors
}

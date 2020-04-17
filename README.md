![Go](https://github.com/WeWorkSandbox/lingo/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/WeWorkSandbox/lingo/branch/master/graph/badge.svg)](https://codecov.io/gh/WeWorkSandbox/lingo)

# Lingo <!-- omit in toc -->

Lingo is used to dynamically create type safe SQL queries in Go. 
It exists to help ease the pain of manually working with SQL.

# Table of Contents <!-- omit in toc -->
- [What is Lingo](#what-is-lingo)
- [Setup](#setup)
  - [Database Support](#database-support)
  - [Generating Table Files](#generating-table-files)
    - [Generate Command](#generate-command)
      - [Config File](#config-file)
      - [Generated Output](#generated-output)
        - [Exported.go](#exportedgo)
        - [Table.go](#tablego)
- [Queries](#queries)
  - [Column Type Safety](#column-type-safety)
    - [MySQL Mappings](#mysql-mappings)
- [Generating SQL to use with Go's SQL Package](#generating-sql-to-use-with-gos-sql-package)
  - [Using Generated SQL](#using-generated-sql)
  - [Execution Example](#execution-example)
  - [Query Dialects](#query-dialects)
- [Select Examples](#select-examples)
  - [Select All from Table](#select-all-from-table)
  - [Select Columns from Table](#select-columns-from-table)
  - [Select with multiple Where Clauses](#select-with-multiple-where-clauses)
  - [Select with nested Where](#select-with-nested-where)
  - [Select with Subquery](#select-with-subquery)
- [Insert Examples](#insert-examples)
  - [Insert Into with Go Types](#insert-into-with-go-types)
  - [Insert Into with Expressions](#insert-into-with-expressions)
  - [Insert Into with Go Types & Expressions](#insert-into-with-go-types--expressions)
  - [Insert Into with Select](#insert-into-with-select)
- [Update Examples](#update-examples)
  - [Update Set All Rows](#update-set-all-rows)
  - [Update Set with Where](#update-set-with-where)
- [Delete Examples](#delete-examples)
  - [Delete All Rows](#delete-all-rows)
  - [Delete with Where](#delete-with-where)
  - [Delete Left Join Where](#delete-left-join-where)

# What is Lingo
Lingo at it's most basic is a type-safe query builder that ensures your SQL queries and commands are accurate and safe. Lingo is a command line tool and Go Module that uses code generation to create type safe table representations based on your actual database schema. These generated structs, when used with the various Lingo packages, ensure that you create valid SQL queries.

With Lingo you can interleave business logic with your query building. Need to filter a SELECT query by different columns based on different request parameters? Lingo lets you do that easily, with no string concatenation required!

Lingo was inspird by [Querydsl](http://www.querydsl.com) and [jOOQ](https://www.jooq.org).

# Setup
Run `mage build` in the root directory of this project to locally build the `lingo` command line tool.

## Database Support
Currently, only MySQL table structure parsing is supported. Note that you can build your own table files by hand using
the existing interfaces which can be used with queries.

## Generating Table Files
In order to use the library in conjunction with a given database schema, you must
upgrade your schema to the latest version, and then run the `lingo generate` command.

### Generate Command
You can include this `lingo generate` command as either a `go:generate` or a bash command.

```.text
Generate entity table and columns from an existing database schema

Usage:
  lingo generate [flags]

Flags:
  -d, --dir string       directory where generated file structure should go (default "./db")
      --driver string    driver name used to initialize the SQL driver (default "mysql")
      --dsn string       data source connection string
  -h, --help             help for generate
  -s, --schema strings   schema name to generate for

Global Flags:
      --config string   configuration file
```

#### Config File
A configuration file can be used to setup each CLI argument. Each config file key is
the same as each CLI command argument.

```yaml
dir: "db/mysql"
schema:
  - "information_schema"
driver_name: "mysql"
dsn: "root:P@ssw0rd@/?maxAllowedPacket=0&parseTime=true"
```

#### Generated Output
Every schema or table is generated with an upper case `Q` for types, and a lower case `q` for packages.

Each schema that the generator is ran for will create a schema package (`q[schemaname]`) and its corresponding tables
under that package (`q[tablename]`).

In each table's `q` package, two files should exist. One of them is named `exported.go` and the other is `table.go`.

##### Exported.go
The file `exported.go` is a purpose built `package` shortcut. It helps in quick query building without needing to hold
on to an instance of the table itself.
For example:
```go
query.DeleteFrom(qcharactersets.Q()).Where(qcharactersets.Maxlen().LT(10)).GetSQL(mysql)
```

##### Table.go
The corresponding `table.go` is the main reference code to each table. It contains methods to give an alias to this table
instance (useful during joins), and each column as its own method. Each column has its own set of methods which help build
out your query. Some of the column methods are specific to the type of column, while others are generic across all column
types.

# Queries
The 4 basic query types exist in the library; `Select`, `Insert`, `Update`, `Delete`. Each query is in the package
`pkg/core/query`.
:warning: _**PLEASE NOTE!!!**_ None of the queries below make logical sense. They are just used to explain how to use the Lingo API
to get the SQL you want.

## Column Type Safety
All of the column types that are generated are configurable via the config file. The column types are only used during query building. 
You do _not_ need to have an exact match of DB type to a Type Safe Column type.

Out of the box, the following Type Safe columns are supported:
- BinaryPath
- BoolPath
- EntityPath
- IntPath
- Int64Path
- JSONPath (small wrapper of StringPath with JSON methods)
- StringPath
- TimePath
- ColumnPath (a generic string defined column)

### MySQL Mappings
When generating, the MySQL type is converted to a Type Safe Column. The default mappings are below:

| MySQL Type | Type Safe Column |
| --- | --- |
| BIGINT | Int64Path |
| BINARY | BinaryPath |
| DATETIME | TimePath |
| INT | IntPath |
| JSON | JSONPath |
| TEXT | StringPath |
| TINYINT | BoolPath |
| TIMESTAMP | TimePath |
| VARCHAR | StringPath |


# Generating SQL to use with Go's SQL Package
All of the examples above are useless until you use the `GetSQL()` method of each `Query` type. `GetSQL()` requires
a Query `Dialect` to be passed in. Doing so will either (1) generate a SQL query string and arguments or (2) return an
error because your syntax or construction of the SQL query is incorrect.

## Using Generated SQL
When `GetSQL()` is called and no error occurs, you get a `core.SQL` object back. This has two methods which are useful
for querying a database. The first is `String() string` which returns a SQL string with `?` for parameters. The second
is an array of empty interface (`interface{}`) types.

## Execution Example
The following is an example showing how to create a Lingo Query, create a DB connection, prepare the statement with the
Lingo SQL generated, and then pass the arguments of your Query into `QueryRow`.

```go
func getCount() int {
    var d core.Dialect = dialect.MySQL{}

    gT := qcharactersets.As("gT")
    var getCountByNameAndDescQuery = Select(Count(Star())).From(gT).Where(gT.DefaultCollateName().IsNull().And(gT.Description().Eq("uuidByte")))
    querySql, err1 := getCountByNameAndDescQuery.GetSQL(d)
    Expect(err1).ShouldNot(HaveOccurred())

    db, dbErr := sql.Open("mysql", "DSN")
    Expect(dbErr).ShouldNot(HaveOccurred())

    ctx, cancelFunc := context.WithTimeout(context.Background(), 200*time.Millisecond)
    stmt, prepareErr := db.PrepareContext(ctx, querySql.String())
    cancelFunc()
    Expect(prepareErr).ShouldNot(HaveOccurred())

    ctx, cancelFunc = context.WithTimeout(context.Background(), 200*time.Millisecond)
    row := stmt.QueryRowContext(ctx, querySql.Values()...)
    cancelFunc()

    var count int
    scanErr := row.Scan(&count)
    Expect(scanErr).ShouldNot(HaveOccurred())

    return count
}
```

## Query Dialects
Every time one attempts to serialize a query using `GetSQL(core.Dialect) (core.SQL, error)`, you are required to pass
in a dialect. These dialects aid in building the SQL properly for the systems the query is to run against.

Right now, only `dialect.MySQL` is built / supported, however, anyone can build their own dialect.


# Select Examples

## Select All from Table
```go
query.SelectFrom(qcharactersets.Q())
```

```mysql
SELECT CHARACTER_SETS.CHARACTER_SET_NAME,
       CHARACTER_SETS.DEFAULT_COLLATE_NAME,
       CHARACTER_SETS.DESCRIPTION,
       CHARACTER_SETS.MAXLEN
FROM   information_schema.CHARACTER_SETS
```
Values: `[]`

## Select Columns from Table
```go
query.Select(qcharactersets.Description(), qcharactersets.Maxlen()).From(qcharactersets.Q())
```

```mysql
SELECT CHARACTER_SETS.DESCRIPTION,
       CHARACTER_SETS.MAXLEN
FROM   information_schema.CHARACTER_SETS
```
Values: `[]`


## Select with multiple Where Clauses
```go
cs := qcharactersets.As("cs")
q := query.SelectFrom(cs)
q = q.Where(cs.Description().Like("utf%"))
q = q.Where(cs.Maxlen().GT(10))
```

```mysql
SELECT cs.CHARACTER_SET_NAME,
       cs.DEFAULT_COLLATE_NAME,
       cs.DESCRIPTION,
       cs.MAXLEN
FROM   information_schema.CHARACTER_SETS AS cs
WHERE  ( cs.DESCRIPTION LIKE ?
         AND cs.MAXLEN > ? )
```
Values: `["utf%", 10]`

## Select with nested Where

In this case, note the ordering of when the `Or` / `And` are called.
```go
cs := qcharactersets.As("cs")
query.SelectFrom(cs).Where(cs.Description().IsNull().Or(cs.CharacterSetName().Like("utf%").Or(cs.Description().Eq("other"))).And(cs.Description().In("desc1", "desc2")))
```

```mysql
SELECT cs.CHARACTER_SET_NAME,
       cs.DEFAULT_COLLATE_NAME,
       cs.DESCRIPTION,
       cs.MAXLEN
FROM   information_schema.CHARACTER_SETS AS cs
WHERE  ( ( cs.DESCRIPTION IS NULL
            OR ( cs.CHARACTER_SET_NAME LIKE ?
                  OR cs.DESCRIPTION = ? ) )
         AND cs.DESCRIPTION IN ( ?, ? ) )
```
Values: `["utf%", "other", "desc1", "desc2"],`

## Select with Subquery
```go
cs := qcharactersets.As("cs")
subQuery := query.Select(cs.Description()).From(cs).Where(cs.Maxlen().GT(50))
query.SelectFrom(cs).Where(cs.Description().InPaths(subQuery))
```

```mysql
SELECT cs.CHARACTER_SET_NAME,
       cs.DEFAULT_COLLATE_NAME,
       cs.DESCRIPTION,
       cs.MAXLEN
FROM   information_schema.CHARACTER_SETS AS cs
WHERE  cs.DESCRIPTION IN (SELECT cs.DESCRIPTION
                          FROM   information_schema.CHARACTER_SETS AS cs
                          WHERE  cs.MAXLEN > ?)
```
Values: `[50]`

# Insert Examples
For all insert statements, you must provide the table to insert into, and then the columns (in order) to insert into,
followed by the values you want to insert. If you have auto generated columns, simply do not specify them in your
`Columns()` functions.

For the values to insert, you can do one of the following to specify values:
1. Use `ValuesConstants()` passing in Go types (which will be converted to Lingo `Expression`s)
2. Use `Values()` passing in other columns or aggregate `Expression`s
3. Use `Values()` with the previous bullet point, but also constants using `expression.NewValue()` / `expression.NewValues()`
to specify which are constants.

## Insert Into with Go Types
```go
charSets := qcharactersets.Q()
query.InsertInto(charSets).Columns(charSets.CharacterSetName(), charSets.DefaultCollateName(), charSets.Description(), charSets.Maxlen(),
).ValuesConstants("char_set_name", "default_collate", "description", 10)
```

```mysql
INSERT INTO information_schema.character_sets
            (character_set_name,
             default_collate_name,
             description,
             maxlen)
VALUES      (?,
             ?,
             ?,
             ?)
```
Values: `["char_set_name", "default_collate", "description", 10]`

## Insert Into with Expressions

```go
query.InsertInto(qcharactersets.Q()).Columns(qcharactersets.Maxlen()).Values(expressions.Count(expressions.Star()))
```

```mysql
INSERT INTO information_schema.character_sets
            (maxlen)
VALUES      (Count(*))
```
Values: `[]`

## Insert Into with Go Types & Expressions

```go
query.InsertInto(qcharactersets.Q()).Columns(
    qcharactersets.Maxlen(), qcharactersets.CharacterSetName(),
).Values(expressions.Count(expressions.Star()), expression.NewValue("default_name"))
```

```mysql
INSERT INTO information_schema.character_sets
            (maxlen,
             character_set_name)
VALUES      (Count(*),
             ?)
```
Values: `["default_name"]`

## Insert Into with Select

```go
charSelect := qcharactersets.As("char_select")
selQuery := query.Select(charSelect.Maxlen(), charSelect.CharacterSetName()).From(charSelect).Where(charSelect.Maxlen().GT(5))

query.InsertInto(qcharactersets.Q()).Columns(qcharactersets.Maxlen(), qcharactersets.CharacterSetName()).Select(selQuery)
```

```mysql
INSERT INTO information_schema.character_sets
            (maxlen,
             character_set_name)
SELECT char_select.maxlen,
       char_select.character_set_name
FROM   information_schema.character_sets AS char_select
WHERE  char_select.maxlen > ?
```
Values: `[5]`

# Update Examples

## Update Set All Rows

```go
query.Update(qcharactersets.Q()).Set(
	qcharactersets.Description().To("new_desc"),
	qcharactersets.DefaultCollateName().To("new_collate_name"),
)
```

```mysql
UPDATE information_schema.character_sets
SET    character_sets.description = ?,
       character_sets.default_collate_name = ?
```

## Update Set with Where

```go
query.Update(qcharactersets.Q()).Set(
	qcharactersets.Description().To("new_desc"),
	qcharactersets.DefaultCollateName().To("new_collate_name"),
).Where(qcharactersets.CharacterSetName().Eq("identity"))
```

```mysql
UPDATE information_schema.character_sets
SET    character_sets.description = ?,
       character_sets.default_collate_name = ?
WHERE  character_sets.character_set_name = ?
```
Values: `["new_desc", "new_collate_name", "identity"]`

# Delete Examples

## Delete All Rows

```go
query.Delete(qcharactersets.Q())
```

```mysql
DELETE FROM information_schema.CHARACTER_SETS
```
Values: `[]`

## Delete with Where

```go
query.Delete(qcharactersets.Q()).Where(qcharactersets.Maxlen().LTOrEq(14))
```

```mysql
DELETE FROM information_schema.character_sets
WHERE  character_sets.maxlen <= ?
```
Values: `[14]`

## Delete Left Join Where

```go
query.Delete(qcharactersets.Q()).Join(
    qcollations.Q(),
    expression.LeftJoin,
    qcharactersets.CharacterSetName().EqPath(qcollations.CharacterSetName()),
).Where(qcharactersets.Maxlen().LTOrEq(14))
```

```mysql
DELETE
FROM      information_schema.character_sets
LEFT JOIN information_schema.collations
ON        character_sets.character_set_name = collations.character_set_name
WHERE     character_sets.maxlen <= ?
```
Values: `[14]`

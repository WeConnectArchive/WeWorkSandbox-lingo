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


# Queries
The 4 basic query types exist in the library; `Select`, `Insert`, `Update`, `Delete`.

## Column Type Safety

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

## Using Generated SQL

## Query Dialects

# Select Examples

## Select All from Table
```go
```

```mysql
```
Values: `[]`

## Select Columns from Table
```go
```

```mysql
```
Values: `[]`


## Select with multiple Where Clauses
```go
```

```mysql
```
Values: `[]`

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
Values: `[],`

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

## Insert Into with Go Types
```go
```

```mysql
```
Values: `[]`

## Insert Into with Expressions

```go
```

```mysql
```
Values: `[]`

## Insert Into with Go Types & Expressions

```go
```

```mysql
```
Values: `[]`

## Insert Into with Select

```go
```

```mysql
```
Values: `[]`

# Update Examples

## Update Set All Rows

```go
```

```mysql
```

## Update Set with Where

```go
```

```mysql
```
Values: `[]`

# Delete Examples

## Delete All Rows

```go
```

```mysql
```
Values: `[]`

## Delete with Where

```go
```

```mysql
```
Values: `[]`

## Delete Left Join Where

```go
```

```mysql
```
Values: `[]`

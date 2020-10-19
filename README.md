![Go Build](https://github.com/WeWorkSandbox/lingo/workflows/Go%20Build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/weworksandbox/lingo)](https://goreportcard.com/report/github.com/weworksandbox/lingo)
[![codecov](https://codecov.io/gh/WeWorkSandbox/lingo/branch/master/graph/badge.svg)](https://codecov.io/gh/WeWorkSandbox/lingo)

# Lingo <!-- omit in toc -->

Lingo at it's most basic is a type-safe query builder and bundled execution framework that ensures your SQL queries
 and commands are accurate and safe, and checks all the Go SQL errors for you.

The idea is not to write an ORM but instead to help (1) build dynamic queries, and (2) quickly write
code to execute those queries against a `sql.DB`. The frameworks for the two concepts can be used together, separately,
or you can use either or. It is up to you!

With Lingos' Query Building, you can interleave business logic with your query building. Need to filter a `SELECT`
query by different columns based on different request parameters? Lingo lets you do that easily, no manual `string`
concatenation / `[]interface{}` appending required!

Lingo can also be used as a command line tool which uses code generation to create type safe table representations
of each table by querying your actual database schema. These generated structs, when used with the various Lingo
packages, ensure that you create valid SQL queries. Those same structs can be created **manually**, if you don't feel
like running a code generator in your build process.

With Lingos' Execution Framework, an `execute` package wraps a native `sql.DB`, and provides simple methods to query
a single row, multiple rows, or executing commands. The `execute` types accept interfaces of the `sql` package types,
facilitating your custom transactional logic or frameworks. The Query Building and Execution Frameworks can be used
completely separately.

Lingo was inspired by [Querydsl](http://www.querydsl.com) and [jOOQ](https://www.jooq.org).

<!-- TODO - Everything below should be refactored to go docs with examples!!! -->

# Table of Contents <!-- omit in toc -->
- [Setup](#setup)
- [Contributors](#contributors)
  - [Developer Setup](#developer-setup)
  - [Licenses](#licenses)

# Setup

Run `go get github.com/weworksandbox/lingo` in the project root. Follow the instructions for [Generate Command](#generate-command).

# Contributors
Adhering to the code of conduct and license, all help is welcome.

## Developer Setup

**Dependencies**
- One time install of [`mage` tool](https://github.com/magefile/mage#installation).
    - You can also install `mage` via `brew install mage` on OSX.
- Ensure [Docker Compose >= v1.25.5](https://github.com/docker/compose/releases/) is installed to use the 3.8 file format.
- Ensure [Docker / Docker Engine >= v19.03.0](https://docs.docker.com/compose/compose-file/) is installed to use the 3.8 file format.

Build `lingo` locally by running `mage build` in the root directory of this project.

Run `mage -v` in the root directory to show commands available to run like tests, linters, databases etc.

## Licenses
This library is a port from [QueryDSL](https://github.com/querydsl/querydsl) into the Go language. The basic blocks of QueryDSL
(Expressions, Operations, Paths) were lifted and translated for Go Generics. Both of these projects use Apache 2.0 license.

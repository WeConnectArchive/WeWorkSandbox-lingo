package core

type Name interface {
	GetName() string
}

type Alias interface {
	GetAlias() string
}

type Table interface {
	Expression
	Alias
	Name
	GetColumns() []Column
	GetParent() string
}

type Column interface {
	Expression
	Alias
	Name
	GetParent() Table
}

type Expression interface {
	GetSQL(d Dialect) (SQL, error)
}

type Set interface {
	Expression
}

type OrderBy interface {
	Expression
}

type SQL interface {
	String() string
	Values() []interface{}
	AppendSQL(right SQL) SQL
	AppendSQLWithSpace(right SQL) SQL
	AppendSQLValues(sql SQL) SQL
	AppendString(str string) SQL
	AppendStringWithSpace(str string) SQL
	AppendFormat(format string, values ...interface{}) SQL
	AppendValues(values []interface{}) SQL
	AppendValuesWithFormat(appendValues []interface{}, format string, values ...interface{}) SQL
	CombineWithSeparator(sqls []SQL, separator string) SQL
	CombinePaths(sqls []SQL) SQL
	SurroundWithParens() SQL
	SurroundWith(left string, right string) SQL
}

type ComboExpression interface {
	Expression
	And(Expression) ComboExpression
	Or(Expression) ComboExpression
}

type Dialect interface {
	Name

}

type Execute interface {

}

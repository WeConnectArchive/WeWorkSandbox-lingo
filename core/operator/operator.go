package operator

type Operand int

const (
	Unknown Operand = iota

	And
	Or

	Eq
	NotEq
	Like
	NotLike

	LessThan
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual

	Null
	NotNull

	In
	NotIn

	Between
	NotBetween
)
const _ = Unknown // Just prevents unused warning

var _names = map[Operand]string{
	And:                "AND",
	Or:                 "OR",
	Eq:                 "=",
	NotEq:              "<>",
	LessThan:           "<",
	LessThanOrEqual:    "<=",
	GreaterThan:        ">",
	GreaterThanOrEqual: ">=",
	Like:               "LIKE",
	NotLike:            "NOT LIKE",
	Null:               "IS NULL",
	NotNull:            "IS NOT NULL",
	In:                 "IN",
	NotIn:              "NOT IN",
	Between:            "BETWEEN",
	NotBetween:         "NOT BETWEEN",
}

func (o Operand) IsValidEnum() bool {
	switch o {
	case And, Or, Eq, NotEq, LessThan, LessThanOrEqual, GreaterThan, GreaterThanOrEqual,
	Like, NotLike, Null, NotNull, In, NotIn, Between, NotBetween: return true
	}
	return false
}


func (o Operand) String() string {
	return _names[o]
}

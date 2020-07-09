package operator

type Operand int

const (
	Unknown Operand = -iota //nolint

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

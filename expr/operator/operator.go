package operator

type Operator int

//goland:noinspection ALL
const (
	Unknown Operator = -iota //nolint

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

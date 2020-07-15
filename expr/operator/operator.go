package operator

type Operator int

//goland:noinspection ALL
const (
	OpUnknown Operator = -iota //nolint

	OpAnd
	OpOr

	OpEq
	OpNotEq
	OpLike
	OpNotLike

	OpLessThan
	OpLessThanOrEqual
	OpGreaterThan
	OpGreaterThanOrEqual

	OpNull
	OpNotNull

	OpIn
	OpNotIn

	OpBetween
	OpNotBetween
)

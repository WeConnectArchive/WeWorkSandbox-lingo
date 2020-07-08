package sort

// Direction is used to specify which type of join a query should be created for.
type Direction int

// Feel free to add any additional Direction's in a given dialect.
// Just ensure the int value for Direction is positive as to not conflict
// with these Direction's
const (
	Unknown Direction = -iota // The `-` in front ensures all values are negative

	Ascending
	Descending
)

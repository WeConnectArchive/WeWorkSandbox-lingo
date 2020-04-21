package join

//revive:disable:redefines-builtin-id - "redefinition of the built-in type Type"
// but the built in type is `type` not `Type`.

// Type is used to specify which type of join a query should be created for.
type Type int

// Feel free to add any additional `Type`'s in a given dialect.
// Just ensure the `int` value for `Type` is positive as to not conflict
// with these `Type`s
const (
	Inner Type = -iota // The `-` in front ensures all values are negative, yay C++ macros!
	Outer
	Left
	Right
)

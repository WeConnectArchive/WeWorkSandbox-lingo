package sort

type Direction int

const (
	Unknown Direction = iota

	Ascending
	Descending
)
const _ = Unknown // Just prevents unused warning

var _names = map[Direction]string{
	Ascending:  "ASC",
	Descending: "DESC",
}

func (d Direction) IsValidEnum() bool {
	switch d {
	case Ascending, Descending: return true
	}
	return false
}

func (d Direction) String() string {
	return _names[d]
}

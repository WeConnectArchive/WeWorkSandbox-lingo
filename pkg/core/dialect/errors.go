package dialect

import (
	"fmt"
)

// EnumIsInvalid creates an error that the specified enum and value combination is invalid
func EnumIsInvalid(name string, value interface{}) error {
	return fmt.Errorf("value '%s' for enum '%s' is invalid", value, name)
}

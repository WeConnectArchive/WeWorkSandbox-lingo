package queries_test

// rows is a shortcut so nobody has to write `[][]interface{}{ /*data*/ }` over and over.
func rows(rows ...[]interface{}) [][]interface{} {
	return rows
}

// row is a shortcut so nobody has to write `[]interface{}{ /*data*/ }` over and over.
func row(values ...interface{}) []interface{} {
	return values
}

// Yes, Yes, these functions have their values escaping the stack, but it is also for a test. Meh. Short lived.

// ptrI16 returns a pointer to the integer passed in.
func ptrI16(i int16) *int16 {
	return &i
}

// ptrI32 returns a pointer to the integer passed in.
func ptrI32(i int32) *int32 {
	return &i
}

// ptrStr returns a pointer to the string passed in.
func ptrStr(s string) *string {
	return &s
}

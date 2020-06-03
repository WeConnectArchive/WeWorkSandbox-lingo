package dialect

// Option for dialects can change the behavior of each Dialect
type Option func(opts *options) error

// WithSchemaNameIncluded will include the schema name for `schema.table` or `schema.table AS alias` SQL output.
func WithSchemaNameIncluded(include bool) Option {
	return func(opts *options) error {
		opts.includeSchemaName = include
		return nil
	}
}

type options struct {
	includeSchemaName bool
}

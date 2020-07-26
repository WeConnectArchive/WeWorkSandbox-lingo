package dialect

import (
	"github.com/weworksandbox/lingo/expr"
)

// Option for dialects can change the behavior of each Dialect
type Option func(opts *options) error

// WithEmptyOption can be useful for ensuring parameters are passed. It is worthless beyond that.
func WithEmptyOption() Option {
	return func(*options) error { return nil }
}

// WithDefaultOperandMappings enables the default opMap as to manually create each. It is enabled by default.
func WithDefaultOperandMappings(enabled bool) Option {
	return func(opts *options) error {
		opts.noDefaultMappings = !enabled
		return nil
	}
}

// WithOperationMapping allows one to change how a dialect operates for a given operation
func WithOperandMapping(op expr.Operator, mapping string) Option {
	return func(opts *options) error {
		if opts.opMap == nil {
			opts.opMap = make(opSyntax)
		}
		opts.opMap[op] = mapping
		return nil
	}
}

// WithSchemaNameIncluded will include the schema name for `schema.table` or `schema.table AS alias` SQL output.
func WithSchemaNameIncluded(include bool) Option {
	return func(opts *options) error {
		opts.includeSchemaName = include
		return nil
	}
}

func WithMaxFormatCacheSize(size uint) Option {
	return func(opts *options) error {
		opts.cacheSize = size
		return nil
	}
}

type options struct {
	noDefaultMappings bool
	opMap             opSyntax
	includeSchemaName bool
	cacheSize         uint
}

package query

import (
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

// Modifier has information to alter the query - post statements.
type Modifier interface {
	IsZero() bool
	Limit() (limit uint64, wasSet bool)
	Offset() (offset uint64, wasSet bool)
}

// ModifyDialect is used by the query builders to check if an interface is satisfied for a dialect.
type ModifyDialect interface {
	Modify(modifier Modifier) (sql.Data, error)
}

// Page helps with pagination by utilizing both Limit and Offset
func Page(limit, offset uint64) Modifier {
	return modifier{
		limitSet:  true,
		limit:     limit,
		offsetSet: true,
		offset:    offset,
	}
}

// Limit helps with limiting the number of rows returned
func Limit(limit uint64) Modifier {
	return modifier{
		limit:    limit,
		limitSet: true,
	}
}

// Offset tells a query to start after a specific number of rows
func Offset(offset uint64) Modifier {
	return modifier{
		offset:    offset,
		offsetSet: true,
	}
}

// modifier holds information to alter the query - post statements. A zero value is valid.
type modifier struct {
	// Keep members in this order - padding

	limit, offset       uint64
	limitSet, offsetSet bool
}

// IsZero can be used to determine if limit or offset was set, without having to check both individually.
func (m modifier) IsZero() bool {
	return !m.limitSet && !m.offsetSet
}

// Limit is returned but only if it wasSet
func (m modifier) Limit() (limit uint64, wasSet bool) {
	return m.limit, m.limitSet
}

// Offset is returned but only if it wasSet
func (m modifier) Offset() (offset uint64, wasSet bool) {
	return m.offset, m.offsetSet
}

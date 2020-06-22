package query

import (
	"github.com/weworksandbox/lingo/pkg/core/sql"
)

// Modify is used by the query builders to check if an interface is satisfied for a dialect.
type Modify interface {
	Modify(modifier Modifier) (sql.Data, error)
}

// Page helps with pagination by utilizing both Limit and Offset
func Page(limit, offset int64) Modifier {
	return Modifier{
		limitSet:  true,
		limit:     limit,
		offsetSet: true,
		offset:    offset,
	}
}

// Limit helps with limiting the number of rows returned
func Limit(limit int64) Modifier {
	return Modifier{
		limit:    limit,
		limitSet: true,
	}
}

// Offset tells a query to start after a specific number of rows
func Offset(offset int64) Modifier {
	return Modifier{
		offset: offset,
		offsetSet: true,
	}
}

// Modifier holds information to alter the query - post statements
type Modifier struct {
	// Keep members in this order - padding

	limit, offset  int64
	limitSet, offsetSet bool
}

// Limit is returned but only if it wasSet
func (m Modifier) Limit() (limit int64, wasSet bool) {
	return m.limit, m.limitSet
}

// Offset is returned but only if it wasSet
 func (m Modifier) Offset() (offset int64, wasSet bool) {
 	return m.offset, m.offsetSet
 }




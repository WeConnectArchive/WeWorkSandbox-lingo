package dialect

import (
	"fmt"

	"github.com/weworksandbox/lingo"
)

// NewMySQL takes options to configure a MySQL schema
//
// There are some caveats for MySQL relating to the dialect.Option types.
// If WithSchemaNameIncluded is not true, ensure the database/schema name is in the sql.DB dataSourceName connection
// string. If not, "Error 1046: No database selected" occurs.
func NewMySQL(opts ...Option) (lingo.Dialect, error) {
	dialect, err := NewDefault(opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to create MySQL dialect: %w", err)
	}
	return MySQL{
		Default: dialect,
	}, nil
}

// MySQL schema has extra MySQL specific features like JSONExtract.
type MySQL struct{ Default }

func (MySQL) GetName() string {
	return "MySQL"
}

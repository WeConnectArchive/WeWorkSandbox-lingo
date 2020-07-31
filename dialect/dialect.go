package dialect

import (
	"fmt"
	"strings"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/expr"
	"github.com/weworksandbox/lingo/expr/join"
	"github.com/weworksandbox/lingo/query"
	"github.com/weworksandbox/lingo/query/sort"
	"github.com/weworksandbox/lingo/sql"
)

// NewDialect takes options to configure a Dialect, this stitches together each expression
func NewDialect(opts ...Option) (Dialect, error) {
	var o options
	for idx := range opts {
		if err := opts[idx](&o); err != nil {
			return Dialect{}, fmt.Errorf("unable to create dialect: %w", err)
		}
	}

	opsMap := make(opSyntax)
	if !o.noDefaultMappings {
		opsMap.Merge(defaultSyntax)
	}
	opsMap.Merge(o.opMap)

	if o.schemaName != "" {
		opsMap.Merge(opSyntax{
			expr.OpSchema:     Syntax(o.schemaName),
			expr.OpTable:      "{0}.{1}",
			expr.OpTableAlias: "{0}.{1} AS {2}",
		})
	}

	cacheSize := int(o.cacheSize)
	if cacheSize < minTemplateCacheSize {
		cacheSize = minTemplateCacheSize
	}

	var idxFmts = make([]string, cacheSize)
	copy(idxFmts, minTemplateCache)
	for idx := minTemplateCacheSize; idx < cacheSize; idx++ {
		idxFmts[idx] = fmt.Sprintf("{%d}", idx)
	}

	return Dialect{
		opMap:        opsMap,
		schemaName:   o.schemaName,
		replaceCache: idxFmts,
	}, nil
}

// Dialect schema uses the generic schema methods to work as a basic ANSI schema.
type Dialect struct {
	schemaName string
	opMap      opSyntax

	replaceCache []string
}

func (Dialect) GetName() lingo.Expression {
	return lingo.ExpressionFunc(func(d lingo.Dialect) (sql.Data, error) {
		return sql.String("dialect.Dialect"), nil
	})
}

func (d Dialect) BuildOperator(op expr.Operator, operands ...lingo.Expression) (sql.Data, error) {
	mapping, ok := d.opMap[op]
	if !ok {
		return nil, fmt.Errorf("operation %d not supported", op)
	}

	growTo := len(mapping)                       // Going to be a minimum of the length of the mapping
	idxers := make([]string, 0, len(operands)*2) // *2 for a from and to value for NewReplacer
	sqlDatas := make([]interface{}, 0, len(operands))
	for idx, exp := range operands {
		s, err := exp.ToSQL(d)
		if err != nil {
			return nil, fmt.Errorf("unable to build operation %s: %w", op, err)
		}

		idxStr := d.idxStr(idx)
		idxers = append(idxers, idxStr, s.String())
		sqlDatas = append(sqlDatas, s.Values()...)

		sqlStr := s.String()
		sqlLen := len(sqlStr)
		count := strings.Count(string(mapping), idxStr)
		growTo += sqlLen * count
	}

	b := strings.Builder{}
	b.Grow(growTo)
	_, err := strings.NewReplacer(idxers...).WriteString(&b, string(mapping))
	if err != nil {
		return nil, fmt.Errorf("unable to build replace operator %s format '%s': %w",
			op, mapping, err)
	}
	return sql.New(b.String(), sqlDatas), nil
}

func (d Dialect) idxStr(idx int) string {
	if idx < len(d.replaceCache) {
		return d.replaceCache[idx]
	}
	return fmt.Sprintf("{%d}", idx)
}

func (Dialect) ValueFormat(count int) sql.Data {
	if count == 0 {
		return sql.Empty()
	}

	const (
		qMark = "?"
		comSp = ", " + qMark
	)

	var s strings.Builder

	numCommas := (count - 1) * len(comSp) // Subtract 1 cuz we add the len of the first question mark next
	s.Grow(numCommas + len(qMark))        // Add the first question mark

	_, _ = s.WriteString(qMark)
	for idx := 1; idx < count; idx++ {
		_, _ = s.WriteString(comSp)
	}
	return sql.String(s.String())
}

func (Dialect) SetValueFormat() string {
	return "="
}

func (d Dialect) Value(value []interface{}) (sql.Data, error) {
	return Value(d, value)
}

func (Dialect) Join(left sql.Data, joinType join.Type, on sql.Data) (sql.Data, error) {
	return Join(left, joinType, on)
}

func (Dialect) OrderBy(left sql.Data, direction sort.Direction) (sql.Data, error) {
	return OrderBy(left, direction)
}

// Modify will build: [LIMIT limit] [OFFSET offset]
func (d Dialect) Modify(m query.Modifier) (sql.Data, error) {
	limit, lWasSet := m.Limit()
	offset, oWasSet := m.Offset()

	s := sql.Empty()
	if lWasSet {
		limitSQL, err := d.Value([]interface{}{limit})
		if err != nil {
			return nil, err
		}
		s = sql.String("LIMIT").AppendWithSpace(limitSQL)
	}
	if oWasSet {
		offsetSQL, err := d.Value([]interface{}{offset})
		if err != nil {
			return nil, err
		}
		s = s.AppendWithSpace(sql.String("OFFSET").AppendWithSpace(offsetSQL))
	}
	return s, nil
}

var minTemplateCacheSize = len(minTemplateCache)
var minTemplateCache = []string{
	"{0}", "{1}", "{2}", "{3}", "{4}", "{5}", "{6}", "{7}", "{8}", "{9}",
	"{10}", "{11}", "{12}", "{13}", "{14}", "{15}", "{16}", "{17}", "{18}", "{19}",
	"{20}", "{21}", "{22}", "{23}", "{24}", "{25}", "{26}", "{27}", "{28}", "{29}",
	"{30}", "{31}", "{32}", "{33}", "{34}", "{35}", "{36}", "{37}", "{38}", "{39}",
	"{40}", "{41}", "{42}", "{43}", "{44}", "{45}", "{46}", "{47}", "{48}", "{49}",
	"{50}", "{51}", "{52}", "{53}", "{54}", "{55}", "{56}", "{57}", "{58}", "{59}",
	"{60}", "{61}", "{62}", "{63}", "{64}", "{65}", "{66}", "{67}", "{68}", "{69}",
	"{70}", "{71}", "{72}", "{73}", "{74}", "{75}", "{76}", "{77}", "{78}", "{79}",
	"{80}", "{81}", "{82}", "{83}", "{84}", "{85}", "{86}", "{87}", "{88}", "{89}",
	"{90}", "{91}", "{92}", "{93}", "{94}", "{95}", "{96}", "{97}", "{98}", "{99}",
}

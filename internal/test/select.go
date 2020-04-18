package test

import (
	. "github.com/onsi/gomega"
	"github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core/dialect"

	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/query"
	"github.com/weworksandbox/lingo/pkg/core/sort"
)

var selectQueries = []Query{
	{
		Name: "SelectFrom",
		Benchmark: true,
		Params: Params{
			Dialect: dialect.Default{},
			SQL: query.SelectFrom(cs),
			SQLAssert: ContainSubstring(trimQuery(`
					SELECT cs.CHARACTER_SET_NAME, cs.DEFAULT_COLLATE_NAME, cs.DESCRIPTION, cs.MAXLEN
					FROM information_schema.CHARACTER_SETS AS cs`)),
			ValuesAssert: BeEmpty(),
			ErrAssert: BeNil(),
		},
	},
	{
		Name: "Select_From(columns)",
		Benchmark: true,
		Params: Params{
			Dialect: dialect.Default{},
			SQL: query.Select(cs.Maxlen(), cs.CharacterSetName()).From(cs),
			SQLAssert: ContainSubstring(trimQuery(`
					SELECT cs.MAXLEN, cs.CHARACTER_SET_NAME
					FROM information_schema.CHARACTER_SETS AS cs`)),
			ValuesAssert: BeEmpty(),
			ErrAssert: BeNil(),
		},
	},
	{
		Name: "Select_From_LeftJoin_Where_OrderBy_Where_Where",
		Benchmark: true,
		Params: Params{
			Dialect: dialect.Default{},
			SQL: query.Select(cs.Description(), cs.CharacterSetName()).
				From(cs).
				Join(col, expression.LeftJoin, cs.CharacterSetName().EqPath(col.CharacterSetName())).
				Where(cs.Maxlen().GT(maxLen)).
				OrderBy(cs.Maxlen(), sort.Descending).
				Where(cs.DefaultCollateName().Eq(defCollName)).
				Where(cs.CharacterSetName().In(charSetNameIn...)),
			SQLAssert: ContainSubstring(trimQuery(`
					SELECT cs.DESCRIPTION, cs.CHARACTER_SET_NAME
					FROM information_schema.CHARACTER_SETS AS cs
					LEFT JOIN information_schema.COLLATIONS AS col
					ON cs.CHARACTER_SET_NAME = col.CHARACTER_SET_NAME
					WHERE (cs.MAXLEN > ? AND cs.DEFAULT_COLLATE_NAME = ? AND cs.CHARACTER_SET_NAME IN (?, ?, ?))
					ORDER BY cs.MAXLEN DESC`)),
			ValuesAssert: matchers.AllInSlice(
				BeEquivalentTo(maxLen),
				BeEquivalentTo(defCollName),
				BeEquivalentTo(charSetNameIn[0]),
				BeEquivalentTo(charSetNameIn[1]),
				BeEquivalentTo(charSetNameIn[2]),
			),
			ErrAssert: BeNil(),
		},
	},
}

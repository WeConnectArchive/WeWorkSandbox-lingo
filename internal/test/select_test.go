package test_test

import (
	. "github.com/onsi/gomega"

	. "github.com/weworksandbox/lingo/internal/test/matchers"
	"github.com/weworksandbox/lingo/pkg/core/dialect"
	"github.com/weworksandbox/lingo/pkg/core/expression"
	"github.com/weworksandbox/lingo/pkg/core/query"
	"github.com/weworksandbox/lingo/pkg/core/sort"
)

var selectQueries = []Query{
	{
		Name:      "SelectFrom",
		Benchmark: true,
		Params: Params{
			Dialect: dialect.Default{},
			SQL:     query.SelectFrom(cs),
			SQLAssert: ContainSubstring(trimQuery(`
					SELECT CS.CHARACTER_SET_NAME, CS.DEFAULT_COLLATE_NAME, CS.DESCRIPTION, CS.MAXLEN
					FROM information_schema.CHARACTER_SETS AS CS`)),
			ValuesAssert: BeEmpty(),
			ErrAssert:    BeNil(),
		},
	},
	{
		Name:      "Select_From(columns)",
		Benchmark: true,
		Params: Params{
			Dialect: dialect.Default{},
			SQL:     query.Select(cs.Maxlen(), cs.CharacterSetName()).From(cs),
			SQLAssert: ContainSubstring(trimQuery(`
					SELECT CS.MAXLEN, CS.CHARACTER_SET_NAME
					FROM information_schema.CHARACTER_SETS AS CS`)),
			ValuesAssert: BeEmpty(),
			ErrAssert:    BeNil(),
		},
	},
	{
		Name:      "Select_From_LeftJoin_Where_OrderBy_Where_Where",
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
					SELECT CS.DESCRIPTION, CS.CHARACTER_SET_NAME
					FROM information_schema.CHARACTER_SETS AS CS
					LEFT JOIN information_schema.COLLATIONS AS COL
					ON CS.CHARACTER_SET_NAME = COL.CHARACTER_SET_NAME
					WHERE (CS.MAXLEN > ? AND CS.DEFAULT_COLLATE_NAME = ? AND CS.CHARACTER_SET_NAME IN (?, ?, ?))
					ORDER BY CS.MAXLEN DESC`)),
			ValuesAssert: AllInSlice(
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

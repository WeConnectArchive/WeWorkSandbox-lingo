package query_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/pkg/core/query"
)

var _ = Describe("modifier.go", func() {
	const (
		zero   = uint64(0)
		limit  = uint64(12)
		offset = uint64(56)
	)
	table.DescribeTable(
		"Modifier",
		func(m query.Modifier, l uint64, lSet bool, o uint64, oSet bool) {
			value, ok := m.Limit()
			Expect(ok).To(Equal(lSet))
			Expect(value).To(Equal(l))

			value, ok = m.Offset()
			Expect(ok).To(Equal(oSet))
			Expect(value).To(Equal(o))
		},
		table.Entry("#Page", query.Page(limit, offset), limit, true, offset, true),
		table.Entry("#Offset", query.Offset(offset), zero, false, offset, true),
		table.Entry("#Limit", query.Limit(limit), limit, true, zero, false),
	)
})

package query_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/pkg/core/query"
)

var _ = Describe("modifier.go", func() {
	const (
		zero = int64(0)
		limit = int64(12)
		offset = int64(56)
	)
	table.DescribeTable(
		"Modifier",
		func(m query.Modifier, l int64, lSet bool, o int64, oSet bool) {
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

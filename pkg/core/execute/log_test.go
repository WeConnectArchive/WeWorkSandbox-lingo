package execute_test

import (
	"context"
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/petergtz/pegomock"
	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/trace"

	"github.com/weworksandbox/lingo/pkg/core/execute"
	"github.com/weworksandbox/lingo/pkg/core/execute/matchers"
)

var _ = Describe("log.go", func() {

	Context("#TraceQuery", func() {
		var (
			ctx       context.Context
			qType     execute.QueryType
			queryStr  string
			queryArgs []interface{}

			tqi *execute.TraceQueryInfo

			span trace.Span
		)
		BeforeEach(func() {
			// Behavior
			span = NewMockSpan()
			pegomock.When(span.IsRecording()).ThenReturn(true)

			// Input
			ctx = trace.ContextWithSpan(context.Background(), span)
			qType = execute.QTRows
			queryStr = "SELECT ? WHERE ABC = ?"
			queryArgs = []interface{}{
				"my first column",
				int32(9999),
			}
		})
		JustBeforeEach(func() {
			tqi = execute.TraceQuery(ctx, qType, queryStr, queryArgs)
		})

		It("Returns a pointer to TraceQueryInfo", func() {
			Expect(tqi).ToNot(BeNil())
		})

		Context("#End", func() {
			BeforeEach(func() {
				qType = execute.QTRow
			})
			JustBeforeEach(func() {
				tqi.End(ctx)
			})
			It("Calls span.AddsEvent", func() {
				expectEndCalled(span, qType, queryStr, queryArgs, 0, nil)
			})
		})

		Context("#RowCount, #End", func() {
			BeforeEach(func() {
				qType = execute.QTRows
			})
			JustBeforeEach(func() {
				tqi.RowCount(99).End(ctx)
			})
			It("Calls span.AddsEvent", func() {
				expectEndCalled(span, qType, queryStr, queryArgs, 99, nil)
			})
		})

		Context("#Err, #End", func() {
			var (
				usedErr error
			)
			BeforeEach(func() {
				usedErr = errors.New("my custom error")
				qType = execute.QTExec
			})
			JustBeforeEach(func() {
				tqi.Err(usedErr).End(ctx)
			})
			It("Calls span.AddsEvent", func() {
				expectEndCalled(span, qType, queryStr, queryArgs, 0, usedErr)
			})
		})

		Context("#RowCount, #Err, #End", func() {
			usedErr := errors.New("my custom error")
			JustBeforeEach(func() {
				tqi.RowCount(100).Err(usedErr).End(ctx)
			})
			It("Calls span.AddsEvent", func() {
				expectEndCalled(span, qType, queryStr, queryArgs, 100, usedErr)
			})
		})
	})

	dateTime := time.Time{}.AddDate(10, 5, 5)
	table.DescribeTable("Changing Values",
		func(arg interface{}, matcher types.GomegaMatcher) {
			// Behavior
			span := NewMockSpan()
			pegomock.When(span.IsRecording()).ThenReturn(true)

			ctx := trace.ContextWithSpan(context.Background(), span)
			queryStr := "my query ?"
			rows := int64(55)

			// When
			execute.TraceQuery(ctx, execute.QTExec, queryStr, []interface{}{arg}).RowCount(rows).End(ctx)

			// Then
			inOrderCtx := pegomock.InOrderContext{}
			_, _, eventFields := span.
				VerifyWasCalledInOrder(pegomock.Once(), &inOrderCtx).
				AddEvent(
					matchers.AnyContextContext(),
					pegomock.AnyString(),
					matchers.AnyCoreKeyValue(), matchers.AnyCoreKeyValue(), matchers.AnyCoreKeyValue(),
				).
				GetCapturedArguments()

			Expect(eventFields).To(HaveLen(3))
			Expect(eventFields).To(ContainElement(core.KeyValue{
				Key:   "SQL",
				Value: core.String(queryStr),
			}))
			Expect(eventFields).To(ContainElement(core.KeyValue{
				Key:   "RowCount",
				Value: core.Int64(rows),
			}))
			Expect(eventFields).To(ContainElement(withCoreValue(matcher)))
		},
		table.Entry("nil", nil, Equal(core.String("nil"))),
		table.Entry("[]byte", []byte{0x01, 0x02}, Equal(core.String("0102"))),
		table.Entry("[][]byte", [][]byte{{0x01, 0x02}, {0x03}}, Equal(core.String("0102,03"))),
		table.Entry("string", "my string", Equal(core.String("my string"))),
		table.Entry("bool", true, Equal(core.Bool(true))),
		table.Entry("int8", int8(-5), Equal(core.Int(-5))),
		table.Entry("int16", int16(-5), Equal(core.Int(-5))),
		table.Entry("int", int(-5), Equal(core.Int(-5))),
		table.Entry("int32", int32(-5), Equal(core.Int32(-5))),
		table.Entry("int64", int64(-5), Equal(core.Int64(-5))),
		table.Entry("uint8", uint8(5), Equal(core.Uint(5))),
		table.Entry("uint16", uint16(5), Equal(core.Uint(5))),
		table.Entry("uint32", uint32(5), Equal(core.Uint32(5))),
		table.Entry("uint64", uint64(5), Equal(core.Uint64(5))),
		table.Entry("float32", float32(5.2), Equal(core.Float32(5.2))),
		table.Entry("float64", float64(5.25e40), Equal(core.Float64(5.25e40))),
		table.Entry("time.Time", dateTime, Equal(core.String(dateTime.String()))),
		table.Entry("customType", struct{ Name string }{Name: "my Name"}, Equal(core.String("{my Name}"))),
	)
})

func expectEndCalled(span trace.Span, qType execute.QueryType, sqlStr string, queryArgs []interface{}, rowCount int64, err error) {
	makeKVP := func() []core.KeyValue {
		keyVals := make([]core.KeyValue, len(queryArgs)+2)
		for idx := range keyVals {
			keyVals[idx] = matchers.AnyCoreKeyValue()
		}
		return keyVals
	}

	inOrderCtx := pegomock.InOrderContext{}
	mockSpan := span.(*MockSpan)
	eventCtx, eventQType, eventFields := mockSpan.
		VerifyWasCalledInOrder(pegomock.Once(), &inOrderCtx).
		AddEvent(matchers.AnyContextContext(), pegomock.AnyString(), makeKVP()...).
		GetCapturedArguments()

	ExpectWithOffset(1, eventCtx).ToNot(BeNil())
	ExpectWithOffset(1, eventQType).To(BeEquivalentTo(qType.String()))
	ExpectWithOffset(1, eventFields).To(HaveLen(len(queryArgs) + 2))
	ExpectWithOffset(1, eventFields).To(ContainElement(core.KeyValue{
		Key:   "SQL",
		Value: core.String(sqlStr),
	}))
	ExpectWithOffset(1, eventFields).To(ContainElement(core.KeyValue{
		Key:   "RowCount",
		Value: core.Int64(rowCount),
	}))

	if err != nil {
		_, recordedErr, _ := mockSpan.VerifyWasCalledInOrder(pegomock.Once(), &inOrderCtx).
			RecordError(matchers.AnyContextContext(), AnyError()).
			GetCapturedArguments()
		ExpectWithOffset(1, recordedErr).To(MatchError(err))
	}
}

func withCoreValue(matcher types.GomegaMatcher) types.GomegaMatcher {
	return WithTransform(func(kv core.KeyValue) core.Value { return kv.Value }, matcher)
}

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
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/kv/value"
	"go.opentelemetry.io/otel/api/trace"

	"github.com/weworksandbox/lingo/execute"
	"github.com/weworksandbox/lingo/execute/matchers"
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
					matchers.AnyKvKeyValue(), matchers.AnyKvKeyValue(), matchers.AnyKvKeyValue(),
				).
				GetCapturedArguments()

			Expect(eventFields).To(HaveLen(3))
			Expect(eventFields).To(ContainElement(kv.String("SQL", queryStr)))
			Expect(eventFields).To(ContainElement(kv.Int64("RowCount", rows)))
			Expect(eventFields).To(ContainElement(withKvValue(matcher)))
		},
		table.Entry("nil", nil, Equal(value.String("nil"))),
		table.Entry("[]byte", []byte{0x01, 0x02}, Equal(value.String("0102"))),
		table.Entry("[][]byte", [][]byte{{0x01, 0x02}, {0x03}}, Equal(value.String("0102,03"))),
		table.Entry("string", "my string", Equal(value.String("my string"))),
		table.Entry("bool", true, Equal(value.Bool(true))),
		table.Entry("int8", int8(-5), Equal(value.Int(-5))),
		table.Entry("int16", int16(-5), Equal(value.Int(-5))),
		table.Entry("int", int(-5), Equal(value.Int(-5))),
		table.Entry("int32", int32(-5), Equal(value.Int32(-5))),
		table.Entry("int64", int64(-5), Equal(value.Int64(-5))),
		table.Entry("uint8", uint8(5), Equal(value.Uint(5))),
		table.Entry("uint16", uint16(5), Equal(value.Uint(5))),
		table.Entry("uint32", uint32(5), Equal(value.Uint32(5))),
		table.Entry("uint64", uint64(5), Equal(value.Uint64(5))),
		table.Entry("float32", float32(5.2), Equal(value.Float32(5.2))),
		table.Entry("float64", float64(5.25e40), Equal(value.Float64(5.25e40))),
		table.Entry("time.Time", dateTime, Equal(value.String(dateTime.String()))),
		table.Entry("customType", struct{ Name string }{Name: "my Name"}, Equal(value.String("{my Name}"))),
	)
})

func expectEndCalled(span trace.Span, qType execute.QueryType, sqlStr string, queryArgs []interface{}, rowCount int64, err error) {
	makeKVP := func() []kv.KeyValue {
		keyVals := make([]kv.KeyValue, len(queryArgs)+2)
		for idx := range keyVals {
			keyVals[idx] = matchers.AnyKvKeyValue()
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
	ExpectWithOffset(1, eventFields).To(ContainElement(kv.String("SQL", sqlStr)))
	ExpectWithOffset(1, eventFields).To(ContainElement(kv.Int64("RowCount", rowCount)))

	if err != nil {
		_, recordedErr, _ := mockSpan.VerifyWasCalledInOrder(pegomock.Once(), &inOrderCtx).
			RecordError(matchers.AnyContextContext(), AnyError()).
			GetCapturedArguments()
		ExpectWithOffset(1, recordedErr).To(MatchError(err))
	}
}

func withKvValue(matcher types.GomegaMatcher) types.GomegaMatcher {
	return WithTransform(func(kv kv.KeyValue) value.Value { return kv.Value }, matcher)
}

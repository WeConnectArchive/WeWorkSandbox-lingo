package execute_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
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

		table.DescribeTable("Changing Values",
			func(arg interface{}) {
				queryArgs[0] = arg

			},
		)
	})
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

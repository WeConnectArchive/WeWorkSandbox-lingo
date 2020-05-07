package execute_test

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/execute"
	"github.com/weworksandbox/lingo/pkg/core/execute/matchers"
)

var _ = Describe("sqlexpexec.go", func() {

	Context("#NewSQLExp", func() {
		var (
			// Input
			s execute.SQL
			d core.Dialect

			// Output
			execExp execute.SQLExp
		)
		BeforeEach(func() {
			d = NewMockDialect()
			s = NewMockSQL()
		})
		JustBeforeEach(func() {
			execExp = execute.NewSQLExp(s, d)
		})

		It("Creates a SQLExp", func() {
			Expect(execExp).ToNot(BeNil())
		})

		Context("#Query", func() {
			var (
				ctx context.Context
				exp core.Expression

				rowScanner execute.RowScanner
				err        error

				expSQL core.SQL
				tSQL   string
				sVals  []interface{}
			)
			BeforeEach(func() {
				ctx = context.Background()
				exp = NewMockExpression()

				tSQL = "select stuffs"
				sVals = []interface{}{
					int32(5),
					"string value",
				}

				expSQL = NewMockCoreSQL()
				pegomock.When(expSQL.String()).ThenReturn(tSQL)
				pegomock.When(expSQL.Values()).ThenReturn(sVals)

				pegomock.When(exp.GetSQL(matchers.AnyCoreDialect())).
					ThenReturn(expSQL, nil)

				pegomock.When(s.Query(
					matchers.AnyContextContext(), pegomock.AnyString(), pegomock.AnyInt32(), pegomock.AnyString(),
				)).ThenReturn(NewMockRowScanner(), nil)
			})
			JustBeforeEach(func() {
				rowScanner, err = execExp.Query(ctx, exp)
			})
			It("Returns a RowScanner", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(rowScanner).ToNot(BeNil())
			})

			Context("GetSQL returns an error", func() {
				BeforeEach(func() {
					pegomock.When(exp.GetSQL(matchers.AnyCoreDialect())).
						ThenReturn(nil, errors.New("random error"))
				})
				It("Returns an error", func() {
					Expect(err).To(MatchError("random error"))
					Expect(rowScanner).To(BeNil())
				})
			})
		})

		Context("#QueryRow", func() {
			var (
				ctx context.Context
				exp core.Expression

				iVal int32
				sVal string
				err  error

				expSQL core.SQL
				tSQL   string
				sVals  []interface{}
			)
			BeforeEach(func() {
				ctx = context.Background()
				exp = NewMockExpression()

				tSQL = "select stuffs"
				sVals = []interface{}{
					int32(5),
					"string value",
				}

				expSQL = NewMockCoreSQL()
				pegomock.When(expSQL.String()).ThenReturn(tSQL)
				pegomock.When(expSQL.Values()).ThenReturn(sVals)

				pegomock.When(exp.GetSQL(matchers.AnyCoreDialect())).
					ThenReturn(expSQL, nil)

				pegomock.When(s.QueryRow(
					matchers.AnyContextContext(),
					pegomock.AnyString(),
					pegomock.AnyInterfaceSlice(),
					pegomock.AnyInterface(),
					pegomock.AnyInterface(),
				)).Then(func(params []pegomock.Param) pegomock.ReturnValues {
					params[3] = int32(99)
					params[4] = "my string"
					return nil
				})
			})
			JustBeforeEach(func() {
				err = execExp.QueryRow(ctx, exp, &iVal, &sVal)
			})
			It("Returns no error and sets the out args", func() {
				Expect(err).ToNot(HaveOccurred())

				mockSQL := s.(*MockSQL)
				_, qStr, values, _ := mockSQL.VerifyWasCalledOnce().QueryRow(
					matchers.AnyContextContext(),
					pegomock.AnyString(),
					pegomock.AnyInterfaceSlice(),
					pegomock.AnyInterface(),
					pegomock.AnyInterface(),
				).GetCapturedArguments()

				Expect(qStr).To(BeEquivalentTo(tSQL))
				Expect(values).To(ContainElements(
					sVals[0], sVals[1],
				))
			})

			Context("GetSQL returns an error", func() {
				BeforeEach(func() {
					pegomock.When(exp.GetSQL(matchers.AnyCoreDialect())).
						ThenReturn(nil, errors.New("random error"))
				})
				It("Returns an error", func() {
					Expect(err).To(MatchError("random error"))
				})
			})
		})

		Context("#Exec", func() {

			var (
				ctx context.Context
				exp core.Expression

				result sql.Result
				err    error

				expSQL core.SQL
				tSQL   string
				sVals  []interface{}
			)
			BeforeEach(func() {
				ctx = context.Background()
				exp = NewMockExpression()

				tSQL = "select stuffs"
				sVals = []interface{}{
					int32(5),
					"string value",
				}

				expSQL = NewMockCoreSQL()
				pegomock.When(expSQL.String()).ThenReturn(tSQL)
				pegomock.When(expSQL.Values()).ThenReturn(sVals)

				pegomock.When(exp.GetSQL(matchers.AnyCoreDialect())).
					ThenReturn(expSQL, nil)

				sqlResult := NewMockResult()
				pegomock.When(s.Exec(
					matchers.AnyContextContext(),
					pegomock.AnyString(),
					pegomock.AnyInterface(),
					pegomock.AnyInterface(),
				)).ThenReturn(sqlResult, nil)
			})
			JustBeforeEach(func() {
				result, err = execExp.Exec(ctx, exp)
			})
			It("Returns the result and no error", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(result).ToNot(BeNil())
			})

			Context("GetSQL returns an error", func() {
				BeforeEach(func() {
					pegomock.When(exp.GetSQL(matchers.AnyCoreDialect())).
						ThenReturn(nil, errors.New("random error"))
				})
				It("Returns an error", func() {
					Expect(err).To(MatchError("random error"))
					Expect(result).To(BeNil())
				})
			})
		})
	})
})

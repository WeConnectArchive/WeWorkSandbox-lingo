package execute_test

import (
	"context"
	gosql "database/sql"
	"errors"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"

	"github.com/weworksandbox/lingo"
	"github.com/weworksandbox/lingo/execute"
	"github.com/weworksandbox/lingo/execute/matchers"
	"github.com/weworksandbox/lingo/sql"
)

var _ = Describe("sqlexpexec.go", func() {

	Context("#NewSQLExp", func() {
		var (
			// Input
			s execute.SQL
			d lingo.Dialect

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

		Context("#BeginTx", func() {
			var (
				ctx  context.Context
				opts gosql.TxOptions

				txSQLExp execute.TxSQLExp
				err      error

				mockTxSQL execute.TxSQL
			)
			BeforeEach(func() {
				ctx = context.Background()
				opts = gosql.TxOptions{
					Isolation: gosql.LevelLinearizable,
					ReadOnly:  true,
				}

				mockTxSQL = NewMockTxSQL()
				pegomock.When(s.BeginTx(matchers.AnyContextContext(), matchers.AnyPtrToSqlTxOptions())).
					ThenReturn(mockTxSQL, nil)
			})
			JustBeforeEach(func() {
				txSQLExp, err = execExp.BeginTx(ctx, &opts)
			})
			It("Returns a TxSQLExp and no error", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(txSQLExp).ToNot(BeNil())
			})

			Context("BeginTx returns an error", func() {
				BeforeEach(func() {
					pegomock.When(s.BeginTx(matchers.AnyContextContext(), matchers.AnyPtrToSqlTxOptions())).
						ThenReturn(nil, errors.New("random error here"))
				})
				It("Returns a nil TxSQLExp and an error", func() {
					Expect(err).To(MatchError("random error here"))
					Expect(txSQLExp).To(BeNil())
				})
			})
		})

		Context("#InTx", func() {
			var (
				ctx      context.Context
				opts     gosql.TxOptions
				execThis execute.ExecSQLExpInTx

				err error

				execThisCalled bool
				mockTxSQL      execute.TxSQL

				didPanic      bool
				panickedValue interface{}
			)
			BeforeEach(func() {
				ctx = context.Background()
				opts = gosql.TxOptions{
					Isolation: gosql.LevelLinearizable,
					ReadOnly:  true,
				}
				execThis = func(ctx context.Context, s execute.ExpQuery) error {
					execThisCalled = true
					return nil
				}

				mockTxSQL = NewMockTxSQL()
				pegomock.When(s.BeginTx(matchers.AnyContextContext(), matchers.AnyPtrToSqlTxOptions())).
					ThenReturn(mockTxSQL, nil)

				pegomock.When(mockTxSQL.CommitOrRollback(
					matchers.AnyContextContext(), AnyError(),
				)).ThenReturn(nil)

				// Reset the any values we dont always set above
				execThisCalled = false
				err = nil
			})
			JustBeforeEach(func() {
				didPanic = true
				defer func() {
					panickedValue = recover()
				}()
				err = execExp.InTx(ctx, &opts, execThis)
				didPanic = false
			})
			It("Returns no error and calls commit", func() {
				Expect(didPanic).To(BeFalse())
				Expect(panickedValue).To(BeNil())

				Expect(execThisCalled).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())

				inOrder := pegomock.InOrderContext{}
				s.(*MockSQL).VerifyWasCalledInOrder(pegomock.Once(), &inOrder).BeginTx(ctx, &opts)
				mockTxSQL.(*MockTxSQL).VerifyWasCalledInOrder(pegomock.Once(), &inOrder).CommitOrRollback(ctx, err)
			})

			Context("BeginTx returns an error", func() {
				BeforeEach(func() {
					pegomock.When(s.BeginTx(matchers.AnyContextContext(), matchers.AnyPtrToSqlTxOptions())).
						ThenReturn(nil, errors.New("begin tx error"))
				})
				It("Returns the Tx error", func() {
					Expect(didPanic).To(BeFalse())
					Expect(panickedValue).To(BeNil())

					Expect(execThisCalled).To(BeFalse())
					Expect(err).To(MatchError("begin tx error"))
				})
			})

			Context("execThis returns an error", func() {
				BeforeEach(func() {
					execThis = func(ctx context.Context, s execute.ExpQuery) error {
						execThisCalled = true
						return errors.New("my error here")
					}

					pegomock.When(mockTxSQL.CommitOrRollback(
						matchers.AnyContextContext(), AnyError(),
					)).Then(func(params []pegomock.Param) pegomock.ReturnValues {
						return pegomock.ReturnValues{params[1]}
					})
				})
				It("Returns the error and rolls back", func() {
					Expect(execThisCalled).To(BeTrue())
					Expect(err).To(MatchError("my error here"))

					Expect(didPanic).To(BeFalse())
					Expect(panickedValue).To(BeNil())

					inOrder := pegomock.InOrderContext{}
					s.(*MockSQL).VerifyWasCalledInOrder(pegomock.Once(), &inOrder).BeginTx(ctx, &opts)
					mockTxSQL.(*MockTxSQL).VerifyWasCalledInOrder(pegomock.Once(), &inOrder).CommitOrRollback(ctx, err)
				})
			})

			Context("execThis panics", func() {
				type myType struct {
					mux   sync.Mutex
					value interface{}
				}
				var (
					panickedValue *myType
				)
				BeforeEach(func() {
					execThis = func(ctx context.Context, s execute.ExpQuery) error {
						execThisCalled = true
						panickedValue = &myType{
							value: 99,
						}
						panic(panickedValue)
					}
				})
				It("Catches the panic, rolls back, and panics the same value", func() {
					Expect(execThisCalled).To(BeTrue())
					Expect(err).To(BeNil())

					Expect(didPanic).To(BeTrue())
					Expect(panickedValue).To(Equal(panickedValue))

					inOrder := pegomock.InOrderContext{}
					s.(*MockSQL).VerifyWasCalledInOrder(pegomock.Once(), &inOrder).BeginTx(ctx, &opts)
					_, rollbackErr := mockTxSQL.(*MockTxSQL).
						VerifyWasCalledInOrder(pegomock.Once(), &inOrder).
						CommitOrRollback(matchers.EqContextContext(ctx), AnyError()).
						GetCapturedArguments()

					Expect(rollbackErr).To(MatchError(fmt.Sprintf("panicked with %v", panickedValue)))
				})
			})
		})

		Context("#Query", func() {
			var (
				ctx context.Context
				exp lingo.Expression

				rowScanner execute.RowScanner
				err        error

				expSQL sql.Data
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

				expSQL = NewMockData()
				pegomock.When(expSQL.String()).ThenReturn(tSQL)
				pegomock.When(expSQL.Values()).ThenReturn(sVals)

				pegomock.When(exp.ToSQL(matchers.AnyLingoDialect())).
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

			Context("ToSQL returns an error", func() {
				BeforeEach(func() {
					pegomock.When(exp.ToSQL(matchers.AnyLingoDialect())).
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
				exp lingo.Expression

				iVal int32
				sVal string
				err  error

				expSQL sql.Data
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

				expSQL = NewMockData()
				pegomock.When(expSQL.String()).ThenReturn(tSQL)
				pegomock.When(expSQL.Values()).ThenReturn(sVals)

				pegomock.When(exp.ToSQL(matchers.AnyLingoDialect())).
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

			Context("ToSQL returns an error", func() {
				BeforeEach(func() {
					pegomock.When(exp.ToSQL(matchers.AnyLingoDialect())).
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
				exp lingo.Expression

				result gosql.Result
				err    error

				expSQL sql.Data
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

				expSQL = NewMockData()
				pegomock.When(expSQL.String()).ThenReturn(tSQL)
				pegomock.When(expSQL.Values()).ThenReturn(sVals)

				pegomock.When(exp.ToSQL(matchers.AnyLingoDialect())).
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

			Context("ToSQL returns an error", func() {
				BeforeEach(func() {
					pegomock.When(exp.ToSQL(matchers.AnyLingoDialect())).
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

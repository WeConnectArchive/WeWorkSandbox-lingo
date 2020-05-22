package execute_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"sync"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/weworksandbox/lingo/pkg/core/execute"
)

var _ = Context("sql.go", func() {

	Context("#NewSQL", func() {
		var (
			// Input
			txDB execute.TxDBQuery

			// Output
			s execute.SQL

			// Behavior
			mock sqlmock.Sqlmock
		)
		BeforeEach(func() {
			var err error
			txDB, mock, err = sqlmock.New()
			Expect(err).ToNot(HaveOccurred(), "unable to setup sqlmock")
			mock.MatchExpectationsInOrder(true)
		})
		JustBeforeEach(func() {
			s = execute.NewSQL(txDB)
		})
		It("Has a valid execute.SQL", func() {
			Expect(s).ToNot(BeNil())
		})

		Context("#Exec", func() {
			var (
				ctx   context.Context
				tSQL  string
				sVals []driver.Value

				result sql.Result
				err    error

				mockQuery *sqlmock.ExpectedExec
			)
			BeforeEach(func() {
				// Input
				ctx = context.Background()
				tSQL = "insert ? ? ?"
				sVals = []driver.Value{
					1, "str", nil,
				}

				// Behavior
				mockQuery = mock.
					ExpectExec(tSQL).
					WithArgs(sVals...)
			})
			JustBeforeEach(func() {
				result, err = s.Exec(ctx, tSQL, driverValueToInterface(sVals)...)
			})

			Context("✅ Successful", func() {
				BeforeEach(func() {
					mockQuery.WillReturnResult(sqlmock.NewResult(1, 1))
				})
				JustAfterEach(func() {
					Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				})
				It("Returns a sql.Result with LastInsertID and RowsAffected", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(result).ToNot(BeNil())

					lastID, err := result.LastInsertId()
					Expect(err).ToNot(HaveOccurred())
					Expect(lastID).To(BeEquivalentTo(1))

					rowsAffected, err := result.RowsAffected()
					Expect(err).ToNot(HaveOccurred())
					Expect(rowsAffected).To(BeEquivalentTo(1))
				})
			})

			Context("🛑 ExecContext has an error", func() {
				BeforeEach(func() {
					mockQuery.WillReturnError(errors.New("random error"))
				})
				It("Returns a nil sql.Result and an error", func() {
					Expect(err).To(MatchError("random error"))
					Expect(result).To(BeNil())
				})
			})
		})

		Context("#Query", func() {
			var (
				ctx   context.Context
				tSQL  string
				sVals []driver.Value

				scanner execute.RowScanner
				err     error

				mockRows  *sqlmock.Rows
				mockQuery *sqlmock.ExpectedQuery
				resultRow [][]driver.Value
			)
			BeforeEach(func() {
				// Input
				ctx = context.Background()
				tSQL = `select ? ? ?`
				sVals = []driver.Value{
					int32(1), "str", time.Now(),
				}

				// Behavior
				resultRow = [][]driver.Value{
					{
						int32(12345),
						"second column data row 1",
						time.Now(),
					},
					{
						int32(98765),
						"second column data row 2",
						time.Now().Add(time.Hour),
					},
				}
			})
			BeforeEach(func() {
				mockRows = sqlmock.NewRows([]string{
					"Int32",
					"String",
					"Time",
				}).AddRow(resultRow[0]...).
					AddRow(resultRow[1]...)

				mockQuery = mock.
					ExpectQuery(tSQL).
					WithArgs(sVals...).
					WillReturnRows(mockRows).
					RowsWillBeClosed()
			})
			JustBeforeEach(func() {
				scanner, err = s.Query(ctx, tSQL, driverValueToInterface(sVals)...)
			})

			It("✅ Returns a RowScanner and no error", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(scanner).ToNot(BeNil())
			})

			Context("QueryContext returns an error", func() {
				BeforeEach(func() {
					mockQuery.WillReturnError(errors.New("random query error"))
				})
				It("Returns an error", func() {
					Expect(err).To(MatchError("random query error"))
					Expect(scanner).To(BeNil())
				})
			})
		})

		Context("#QueryRow", func() {
			var (
				ctx   context.Context
				tSQL  string
				sVals []driver.Value
				iVal  int32
				sVal  string
				tVal  time.Time

				err error

				rowVals []driver.Value
			)
			BeforeEach(func() {
				ctx = context.Background()
				tSQL = `select ? ? ?`
				sVals = []driver.Value{
					int32(1),
					"second column",
					time.Now(),
				}

				rowVals = []driver.Value{
					int32(12345),
					"second result",
					time.Now().Add(time.Hour),
				}
			})
			JustBeforeEach(func() {
				err = s.QueryRow(ctx, tSQL, driverValueToInterface(sVals), &iVal, &sVal, &tVal)
			})

			Context("Success", func() {
				BeforeEach(func() {
					mock.ExpectQuery(tSQL).
						WithArgs(sVals...).
						WillReturnRows(
							sqlmock.NewRows([]string{
								"First",
								"Second",
								"Third",
							}).AddRow(rowVals...),
						).RowsWillBeClosed()
				})

				It("Returns no error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			Context("QueryRow errors", func() {
				BeforeEach(func() {
					mock.ExpectQuery(tSQL).
						WithArgs(sVals...).
						WillReturnError(errors.New("random error"))
				})

				It("Returns the error", func() {
					Expect(err).To(MatchError("random error"))
				})
			})
		})

		Context("#InTx", func() {
			var (
				ctx      context.Context
				opts     *sql.TxOptions
				execThis execute.ExecSQLInTx

				err error

				didPanic      bool
				panickedValue interface{}
				mockBegin     *sqlmock.ExpectedBegin
			)
			BeforeEach(func() {
				ctx = context.Background()
				opts = &sql.TxOptions{
					Isolation: sql.LevelReadUncommitted,
					ReadOnly:  false,
				}

				mockBegin = mock.ExpectBegin()
			})

			JustBeforeEach(func() {
				didPanic = true
				defer func() {
					panickedValue = recover()
				}()
				err = s.InTx(ctx, opts, execThis)
				didPanic = false
			})

			Context("execThis returns nil", func() {
				var (
					mockCommit *sqlmock.ExpectedCommit
					execCalled bool
				)
				BeforeEach(func() {
					execThis = func(ctx context.Context, s execute.SQLQuery) error {
						execCalled = true
						return nil
					}

					mockCommit = mock.ExpectCommit()
				})
				It("Returns no error and calls commit", func() {
					Expect(didPanic).To(BeFalse())
					Expect(panickedValue).To(BeNil())
					Expect(execCalled).To(BeTrue())

					Expect(err).ToNot(HaveOccurred())
					Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				})

				Context("Commit returns an error", func() {
					BeforeEach(func() {
						mockCommit.WillReturnError(errors.New("random commit error"))
					})
					It("Returns no error and calls commit", func() {
						Expect(didPanic).To(BeFalse())
						Expect(panickedValue).To(BeNil())

						Expect(err).To(MatchError("random commit error"))
						Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
					})
				})
			})

			Context("Begin returns an error", func() {
				BeforeEach(func() {
					mockBegin.WillReturnError(errors.New("random begin error"))
				})
				It("Returns no error and calls commit", func() {
					Expect(didPanic).To(BeFalse())
					Expect(panickedValue).To(BeNil())

					Expect(err).To(MatchError("random begin error"))
					Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				})
			})

			Context("execThis returns an error", func() {
				var (
					mockRollback *sqlmock.ExpectedRollback
				)
				BeforeEach(func() {
					execThis = func(ctx context.Context, s execute.SQLQuery) error {
						return errors.New("random error")
					}

					mockRollback = mock.ExpectRollback()
				})
				It("Returns the execThis error and calls Rollback", func() {
					Expect(didPanic).To(BeFalse())
					Expect(panickedValue).To(BeNil())

					Expect(err).To(MatchError("random error"))
					Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				})

				Context("Rollback returns an error", func() {
					BeforeEach(func() {
						mockRollback.WillReturnError(errors.New("random rollback error"))
					})
					It("Returns the rollback error wrapped around the execThis error", func() {
						Expect(didPanic).To(BeFalse())
						Expect(panickedValue).To(BeNil())

						Expect(err).To(MatchError(SatisfyAll(
							ContainSubstring("random error"),
							ContainSubstring("random rollback error"),
						)))
						Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
					})
				})
			})

			Context("execThis panics", func() {
				type myType struct {
					mux   sync.Mutex
					value interface{}
				}
				var (
					mockRollback  *sqlmock.ExpectedRollback
					panickedValue *myType
				)
				BeforeEach(func() {
					execThis = func(ctx context.Context, s execute.SQLQuery) error {
						panickedValue = &myType{
							value: 99,
						}
						panic(panickedValue)
					}

					mockRollback = mock.ExpectRollback()
				})

				It("Rolls back the transaction and rethrows the value", func() {
					Expect(didPanic).To(BeTrue())
					Expect(panickedValue).To(Equal(panickedValue))
					Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				})

				Context("Rollback returns an error", func() {
					BeforeEach(func() {
						mockRollback.WillReturnError(errors.New("my random error"))
					})

					It("Rolls back the transaction and rethrows the value", func() {
						Expect(didPanic).To(BeTrue())
						Expect(panickedValue).To(Equal(panickedValue))
						Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
					})
				})
			})
		})
	})
})

func driverValueToInterface(vals []driver.Value) []interface{} {
	var result = make([]interface{}, len(vals))
	for idx := range vals {
		result[idx] = vals[idx]
	}
	return result
}

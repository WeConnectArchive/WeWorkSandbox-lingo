package execute

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("rowscanner.go", func() {

	var (
		ctx  context.Context
		rows *sql.Rows

		scanner rowScanner

		db   *sql.DB
		mock sqlmock.Sqlmock
	)
	BeforeEach(func() {
		ctx = context.Background()

		var err error
		db, mock, err = sqlmock.New()
		Expect(err).ToNot(HaveOccurred())
	})

	JustBeforeEach(func() {
		scanner = rowScanner{
			rows: rows,
		}
	})

	Context("with rows setup", func() {
		var (
			resultRow [][]driver.Value

			mockRows  *sqlmock.Rows
			mockQuery *sqlmock.ExpectedQuery
		)

		BeforeEach(func() {
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

			mockRows = sqlmock.NewRows([]string{
				"Int32",
				"String",
				"Time",
			}).AddRow(resultRow[0]...).
				AddRow(resultRow[1]...)

			mockQuery = mock.
				ExpectQuery("SELECT 1").
				WillReturnRows(mockRows).
				RowsWillBeClosed()
		})

		Context("#ScanRow", func() {
			var (
				iVal    int32
				strVal  string
				timeVal time.Time
			)
			var (
				nextRow bool
			)
			BeforeEach(func() {
				var err error
				rows, err = db.Query("SELECT 1")
				Expect(err).ToNot(HaveOccurred())
			})
			JustBeforeEach(func() {
				nextRow = scanner.ScanRow(&iVal, &strVal, &timeVal)
			})
			It("✅ Retrieves first row with 3 columns", func() {
				Expect(nextRow).To(BeTrue())
				Expect(iVal).To(Equal(resultRow[0][0]))
				Expect(strVal).To(Equal(resultRow[0][1]))
				Expect(timeVal).To(Equal(resultRow[0][2]))
			})

			Context("🛑 #ScanRow with not enough columns for query", func() {
				JustBeforeEach(func() {
					nextRow = scanner.ScanRow(iVal, strVal)
				})
				It("Returns false", func() {
					Expect(nextRow).To(BeFalse())
				})
				Context("✅ #Err", func() {
					var (
						scanErr error
					)
					JustBeforeEach(func() {
						scanErr = scanner.Err(ctx)
					})
					It("Returns the Scan error", func() {
						Expect(scanErr).To(MatchError("sql: expected 3 destination arguments in Scan, not 2"))
					})
				})
			})

			Context("#ScanRow", func() {
				JustBeforeEach(func() {
					nextRow = scanner.ScanRow(&iVal, &strVal, &timeVal)
				})
				It("✅ Retrieves second row with 3 columns with no more columns", func() {
					Expect(nextRow).To(BeTrue())
					Expect(iVal).To(Equal(resultRow[1][0]))
					Expect(strVal).To(Equal(resultRow[1][1]))
					Expect(timeVal).To(Equal(resultRow[1][2]))
				})

				Context("#Err", func() {
					var (
						scanErr error
					)
					JustBeforeEach(func() {
						scanErr = scanner.Err(ctx)
					})
					It("Returns no error and closes rows", func() {
						Expect(scanErr).ToNot(HaveOccurred())
					})
				})

				Context("🛑 Unable to retrieve first row", func() {
					BeforeEach(func() {
						mockRows.RowError(1, errors.New("random error"))
						mockRows.CloseError(errors.New("close error"))
						mockQuery.RowsWillBeClosed()
					})
					It("Returns false due to error", func() {
						Expect(nextRow).To(BeFalse())
						Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
					})
				})

				Context("✅ #ScanRow successful but done with rows", func() {
					JustBeforeEach(func() {
						nextRow = scanner.ScanRow(&iVal, &strVal, &timeVal)
					})
					It("Returns false", func() {
						Expect(nextRow).To(BeFalse())
						Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
					})

					Context("#Err", func() {
						var (
							scanErr error
						)
						JustBeforeEach(func() {
							scanErr = scanner.Err(ctx)
						})
						It("Returns no error", func() {
							Expect(scanErr).ToNot(HaveOccurred())
							Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
						})
					})
				})
			})
		})

		Context("#Close", func() {
			BeforeEach(func() {
				var err error
				rows, err = db.Query("SELECT 1")
				Expect(err).ToNot(HaveOccurred())
			})
			JustBeforeEach(func() {
				scanner.Close(ctx)
			})
			It("Calls close", func() {
				Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
			})
		})

		Context("#Err", func() {
			var (
				errResult error
			)
			BeforeEach(func() {
				var err error
				rows, err = db.Query("SELECT 1")
				Expect(err).ToNot(HaveOccurred())
			})
			JustBeforeEach(func() {
				errResult = scanner.Err(ctx)
			})
			It("Returns no error", func() {
				Expect(errResult).ToNot(HaveOccurred())
			})

			Context("Close returns an error", func() {
				BeforeEach(func() {
					mockRows.CloseError(errors.New("close error"))
				})
				It("Returns the close error", func() {
					Expect(errResult).To(MatchError("close error"))
					Expect(mock.ExpectationsWereMet()).ToNot(HaveOccurred())
				})
			})
		})
	})
})

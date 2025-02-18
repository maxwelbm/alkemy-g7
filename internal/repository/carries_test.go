package repository_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

var logMockCarrier = mocks.MockLog{}

func TestCarriers_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.NewCarriersRepository(db, logMockCarrier)

	t.Run("Success GetByID", func(t *testing.T) {
		expectedCarrier := model.Carries{
			ID:          1,
			CID:         "CID001",
			CompanyName: "ABC Company",
			Address:     "123 Main St",
			Telephone:   "1234567890",
			LocalityID:  1,
		}

		rows := mock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id FROM carriers WHERE `id` = ?"}).
			AddRow(expectedCarrier.ID, expectedCarrier.CID, expectedCarrier.CompanyName, expectedCarrier.Address, expectedCarrier.Telephone, expectedCarrier.LocalityID)
		expectedQuery := "SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `carriers` WHERE `id` = ?"
		mock.ExpectQuery(expectedQuery).
			WithArgs(expectedCarrier.ID).
			WillReturnRows(rows)

		carrier, err := rp.GetByID(expectedCarrier.ID)

		assert.NoError(t, err)
		assert.Equal(t, expectedCarrier, carrier)
	})

	t.Run("Not Found GetByID", func(t *testing.T) {

		expectedQuery := "SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `carriers` WHERE `id` = ?"
		mock.ExpectQuery(expectedQuery).
			WithArgs(1).
			WillReturnError(sql.ErrNoRows)

		carrier, err := rp.GetByID(1)

		assert.Error(t, err)
		assert.IsType(t, &customerror.CarrierError{}, err)
		assert.Equal(t, model.Carries{}, carrier)

	})

}

func TestCarriers_PostCarrier(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.NewCarriersRepository(db, logMockCarrier)

	t.Run("Success PostCarrier", func(t *testing.T) {
		newCarrier := model.Carries{
			CID:         "CID001",
			CompanyName: "ABC Company",
			Address:     "123 Main St",
			Telephone:   "1234567890",
			LocalityID:  1,
		}

		mock.ExpectExec("INSERT INTO `carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").
			WithArgs(newCarrier.CID, newCarrier.CompanyName, newCarrier.Address, newCarrier.Telephone, newCarrier.LocalityID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := rp.PostCarrier(newCarrier)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), id)
	})

	t.Run("Error Duplicate CID", func(t *testing.T) {
		newCarrier := model.Carries{
			CID:         "CID001",
			CompanyName: "ABC Company",
			Address:     "123 Main St",
			Telephone:   "1234567890",
			LocalityID:  1,
		}

		mock.ExpectExec("INSERT INTO `carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").
			WithArgs(newCarrier.CID, newCarrier.CompanyName, newCarrier.Address, newCarrier.Telephone, newCarrier.LocalityID).
			WillReturnError(&mysql.MySQLError{Number: 1062})

		id, err := rp.PostCarrier(newCarrier)

		assert.Error(t, err)
		assert.Equal(t, int64(0), id)
	})

	t.Run("Error Other DB Error", func(t *testing.T) {
		newCarrier := model.Carries{
			CID:         "CID001",
			CompanyName: "ABC Company",
			Address:     "123 Main St",
			Telephone:   "1234567890",
			LocalityID:  1,
		}

		mock.ExpectExec("INSERT INTO `carriers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").
			WithArgs(newCarrier.CID, newCarrier.CompanyName, newCarrier.Address, newCarrier.Telephone, newCarrier.LocalityID).
			WillReturnError(errors.New("some other error"))

		id, err := rp.PostCarrier(newCarrier)

		assert.Error(t, err)
		assert.Equal(t, int64(0), id)
	})
}

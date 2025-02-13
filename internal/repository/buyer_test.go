package repository_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/stretchr/testify/assert"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestBuyerRepository_Post(t *testing.T) {

	//inicializando o mock que retorna um *sql.DB, um mockSQl e um Error
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.NewBuyerRepository(db)

	t.Run("Verifies successful addition of a buyer", func(t *testing.T) {
		buyer := model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mock.ExpectPrepare("INSERT INTO buyers").ExpectExec().
			WithArgs(buyer.CardNumberID, buyer.FirstName, buyer.LastName).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := rp.Post(buyer)

		expectedId := 1

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.NoError(t, err)
		assert.Equal(t, int64(expectedId), id)

	})

	t.Run("Verifies error addition of a buyer", func(t *testing.T) {
		buyer := model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mock.ExpectPrepare("INSERT INTO buyers").WillReturnError(errors.New("prepare statement error"))

		_, err := rp.Post(buyer)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.Error(t, err)

	})

	t.Run("return MySQLError type 1062", func(t *testing.T) {
		buyer := model.Buyer{FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mock.ExpectPrepare("INSERT INTO buyers")

		mock.ExpectExec("INSERT INTO buyers").
			WithArgs(buyer.CardNumberID, buyer.FirstName, buyer.LastName).
			WillReturnError(&mysql.MySQLError{
				Number:   1062,
				SQLState: [5]byte{'2', '3', '0', '0', '0'},
				Message:  "Duplicate entry",
			})

		_, err := rp.Post(buyer)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.Error(t, err)

	})
}

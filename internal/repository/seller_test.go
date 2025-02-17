package repository_test

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSellersRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositorySellers(db)

	t.Run("test repository method for get all sellers with success", func(t *testing.T) {
		expectedSellers := []model.Seller{
			{ID: 1, CID: 1, CompanyName: "Enterprise Liberty", Address: "456 Elm St", Telephone: "4443335454", Locality: 1},
			{ID: 2, CID: 2, CompanyName: "Libre Mercado", Address: "123 Montain St Avenue", Telephone: "5554545999", Locality: 2},
		}

		rows := mock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id"})
		for _, seller := range expectedSellers {
			rows.AddRow(seller.ID, seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality)
		}

		mock.ExpectQuery("SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `sellers`").WillReturnRows(rows)

		sellers, err := rp.Get()
		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, err)
		assert.Equal(t, expectedSellers, sellers)
		assert.NoError(t, errMock)
	})

	//t.Run("test repository method for get all sellers with errors", func(t *testing.T) {
	//	mock.ExpectQuery("SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `sellers`").WillReturnError(customerror.ErrDefaultSellerSQL)
	//
	//	sellers, err := rp.Get()
	//	mockErr := mock.ExpectationsWereMet()
	//	expectedError := customerror.ErrDefaultSellerSQL
	//	assert.ErrorIs(t, expectedError, err)
	//	assert.Equal(t, []model.Seller(nil), sellers)
	//	assert.NoError(t, mockErr)
	//})

	t.Run("test repository method for get all sellers with query error", func(t *testing.T) {
		mock.ExpectQuery("SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `sellers`").WillReturnError(sql.ErrNoRows)

		sellers, err := rp.Get()
		mockErr := mock.ExpectationsWereMet()
		assert.Error(t, err)
		assert.Empty(t, sellers)
		assert.NoError(t, mockErr)
	})
}

func TestSellersRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositorySellers(db)

	t.Run("test repository method for get seller by ID with success", func(t *testing.T) {
		ID := 1
		seller := model.Seller{ID: 1, CID: 1, CompanyName: "Enterprise Liberty", Address: "456 Elm St", Telephone: "4443335454", Locality: 1}

		rows := sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id"}).
			AddRow(seller.ID, seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality)

		mock.ExpectQuery("SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `sellers` WHERE `id` = ?").
			WithArgs(ID).
			WillReturnRows(rows)

		result, err := rp.GetByID(ID)
		assert.NoError(t, err)

		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)

		assert.Equal(t, seller, result)
	})

	t.Run("test repository method for get seller by id with error not found", func(t *testing.T) {
		ID := 99

		mock.ExpectQuery("SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `sellers` WHERE `id` = ?").
			WithArgs(ID).
			WillReturnError(sql.ErrNoRows)

		seller, err := rp.GetByID(ID)
		errorExpected := customerror.NewSellerErr(customerror.ErrNotFound.Error(), http.StatusNotFound)
		errorMock := mock.ExpectationsWereMet()

		assert.NoError(t, errorMock)
		assert.Error(t, errorExpected, err)
		assert.Equal(t, model.Seller{}, seller)

	})
}

func TestSellersRepository_Post(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositorySellers(db)

	t.Run("test repository method for create seller with success", func(t *testing.T) {
		seller := model.Seller{ID: 1, CID: 1, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 1}

		mock.ExpectExec("INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").
			WithArgs(seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery("SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `sellers` WHERE `id` = ?").
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id"}).
				AddRow(seller.ID, seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality))

		sl, err := rp.Post(&seller)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.NoError(t, err)
		assert.Equal(t, seller.ID, sl.ID)
		assert.Equal(t, seller, sl)
	})

	t.Run("test repository method for create seller with insert error", func(t *testing.T) {
		seller := model.Seller{ID: 5, CID: 5, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 5}

		mock.ExpectExec("INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").
			WithArgs(seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality).
			WillReturnResult(sqlmock.NewErrorResult(errors.New("error")))

		sl, err := rp.Post(&seller)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.Error(t, err)
		assert.Empty(t, sl)
		assert.NotEqual(t, seller, sl)
	})

	t.Run("test repository method for create seller with sql duplicated error", func(t *testing.T) {
		seller := model.Seller{ID: 3, CID: 3, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 3}

		mock.ExpectExec("INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").
			WithArgs(seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality).
			WillReturnError(&mysql.MySQLError{Number: 1062})

		sl, err := rp.Post(&seller)
		errMock := mock.ExpectationsWereMet()
		expectedErr := customerror.ErrCIDSellerAlreadyExist

		assert.NoError(t, errMock)
		assert.ErrorIs(t, expectedErr, err)
		assert.Error(t, err)
		assert.Empty(t, sl)
	})

	t.Run("test repository method for create seller with sql null attribute error", func(t *testing.T) {
		seller := model.Seller{ID: 3, CID: 3, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 3}

		mock.ExpectExec("INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").
			WithArgs(seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality).
			WillReturnError(&mysql.MySQLError{Number: 1048})

		sl, err := rp.Post(&seller)
		errMock := mock.ExpectationsWereMet()
		expectedErr := customerror.ErrNullSellerAttribute

		assert.NoError(t, errMock)
		assert.ErrorIs(t, expectedErr, err)
		assert.Error(t, err)
		assert.Empty(t, sl)
	})

	t.Run("test repository method for create seller with sql invalid json error", func(t *testing.T) {
		seller := model.Seller{ID: 3, CID: 3, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 3}

		mock.ExpectExec("INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").
			WithArgs(seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality).
			WillReturnError(&mysql.MySQLError{Number: 1064})

		sl, err := rp.Post(&seller)
		errMock := mock.ExpectationsWereMet()
		expectedErr := customerror.ErrInvalidSellerJSONFormat

		assert.NoError(t, errMock)
		assert.ErrorIs(t, expectedErr, err)
		assert.Error(t, err)
		assert.Empty(t, sl)
	})

	t.Run("test repository method for create seller with sql default error", func(t *testing.T) {
		seller := model.Seller{ID: 3, CID: 3, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 3}

		mock.ExpectExec("INSERT INTO `sellers` (`cid`, `company_name`, `address`, `telephone`, `locality_id`) VALUES (?, ?, ?, ?, ?)").
			WithArgs(seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality).
			WillReturnError(&mysql.MySQLError{Number: 1205})

		sl, err := rp.Post(&seller)
		errMock := mock.ExpectationsWereMet()
		expectedErr := customerror.ErrDefaultSellerSQL

		assert.NoError(t, errMock)
		assert.ErrorIs(t, expectedErr, err)
		assert.Error(t, err)
		assert.Empty(t, sl)
	})
}

func TestSellersRepository_Patch(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositorySellers(db)

	t.Run("test repository method for update seller with success", func(t *testing.T) {
		ID := 4
		seller := model.Seller{ID: 4, CID: 4, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 4}

		mock.ExpectExec("UPDATE `sellers` SET `cid` = ?, `company_name` = ?, `address` = ?, `telephone` = ?, `locality_id` = ? WHERE `id` = ?").
			WithArgs(seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality, seller.ID).
			WillReturnResult(sqlmock.NewResult(int64(ID), 1))

		mock.ExpectQuery("SELECT `id`, `cid`, `company_name`, `address`, `telephone`, `locality_id` FROM `sellers` WHERE `id` = ?").
			WithArgs(ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "cid", "company_name", "address", "telephone", "locality_id"}).
				AddRow(seller.ID, seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality))

		sl, err := rp.Patch(ID, &seller)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.NoError(t, err)
		assert.Equal(t, seller.ID, sl.ID)
		assert.Equal(t, seller, sl)
	})

	t.Run("test repository method for update seller with sql duplicated error", func(t *testing.T) {
		ID := 7
		seller := model.Seller{ID: 7, CID: 7, CompanyName: "Enterprise Science", Address: "1200 Central Park Avenue", Telephone: "999444555", Locality: 7}

		mock.ExpectExec("UPDATE `sellers` SET `cid` = ?, `company_name` = ?, `address` = ?, `telephone` = ?, `locality_id` = ? WHERE `id` = ?").
			WithArgs(seller.CID, seller.CompanyName, seller.Address, seller.Telephone, seller.Locality, seller.ID).
			WillReturnError(&mysql.MySQLError{Number: 1062})

		sl, err := rp.Patch(ID, &seller)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.Error(t, err)
		assert.Empty(t, sl)
	})
}

func TestSellersRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositorySellers(db)

	t.Run("test repository method for delete seller with success", func(t *testing.T) {
		ID := 1

		mock.ExpectExec("DELETE FROM `sellers` WHERE `id` = ?").
			WithArgs(ID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := rp.Delete(ID)

		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, mockErr)
		assert.NoError(t, err)
	})

	t.Run("test repository method for delete seller with sql error", func(t *testing.T) {
		ID := 1

		mock.ExpectExec("DELETE FROM `sellers` WHERE `id` = ?").
			WithArgs(1).
			WillReturnError(&mysql.MySQLError{Number: 1451})

		err := rp.Delete(ID)

		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, mockErr)
		assert.Error(t, err)
	})
}

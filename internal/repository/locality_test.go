package repository_test

import (
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func TestLocalitiesRepository_CreateLocality(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositoryLocalities(db, logMock)

	t.Run("test repository method for create locality with success", func(t *testing.T) {
		l := model.Locality{ID: 1, Locality: "Denver", Province: "Colorado", Country: "EUA"}

		mock.ExpectExec("INSERT INTO `locality` (`locality_name`, `province_name`, `country_name`) VALUES (?, ?, ?)").
			WithArgs(l.Locality, l.Province, l.Country).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery("SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality` WHERE `id` = ?").
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "locality_name", "province_name", "country_name"}).
				AddRow(l.ID, l.Locality, l.Province, l.Country))

		locality, err := rp.CreateLocality(&l)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.NoError(t, err)
		assert.Equal(t, l.ID, locality.ID)
		assert.Equal(t, l, locality)
	})

	t.Run("test repository method for create locality with insert error", func(t *testing.T) {
		l := model.Locality{ID: 7, Locality: "Manhattan", Province: "New York", Country: "EUA"}

		mock.ExpectExec("INSERT INTO `locality` (`locality_name`, `province_name`, `country_name`) VALUES (?, ?, ?)").
			WithArgs(l.Locality, l.Province, l.Country).
			WillReturnResult(sqlmock.NewErrorResult(errors.New("error")))

		locality, err := rp.CreateLocality(&l)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.Error(t, err)
		assert.Empty(t, locality)
		assert.NotEqual(t, l, locality)
	})

	t.Run("test repository method for create locality with sql null attribute error", func(t *testing.T) {
		l := model.Locality{ID: 10, Locality: "Los Angeles", Province: "California", Country: "EUA"}

		mock.ExpectExec("INSERT INTO `locality` (`locality_name`, `province_name`, `country_name`) VALUES (?, ?, ?)").
			WithArgs(l.Locality, l.Province, l.Country).
			WillReturnError(&mysql.MySQLError{Number: 1048})

		locality, err := rp.CreateLocality(&l)
		errMock := mock.ExpectationsWereMet()
		expectedErr := customerror.ErrNullLocalityAttribute

		assert.NoError(t, errMock)
		assert.ErrorIs(t, expectedErr, err)
		assert.Error(t, err)
		assert.Empty(t, locality)
	})

	t.Run("test repository method for create locality with sql invalid json error", func(t *testing.T) {
		l := model.Locality{ID: 9, Locality: "Little Rock", Province: "Arkansas", Country: "EUA"}

		mock.ExpectExec("INSERT INTO `locality` (`locality_name`, `province_name`, `country_name`) VALUES (?, ?, ?)").
			WithArgs(l.Locality, l.Province, l.Country).
			WillReturnError(&mysql.MySQLError{Number: 1064})

		locality, err := rp.CreateLocality(&l)
		errMock := mock.ExpectationsWereMet()
		expectedErr := customerror.ErrInvalidLocalityJSONFormat

		assert.NoError(t, errMock)
		assert.ErrorIs(t, expectedErr, err)
		assert.Error(t, err)
		assert.Empty(t, locality)
	})

	t.Run("test repository method for create seller with sql default error", func(t *testing.T) {
		l := model.Locality{ID: 17, Locality: "Phoenix", Province: "Arizona", Country: "EUA"}

		mock.ExpectExec("INSERT INTO `locality` (`locality_name`, `province_name`, `country_name`) VALUES (?, ?, ?)").
			WithArgs(l.Locality, l.Province, l.Country).
			WillReturnError(&mysql.MySQLError{Number: 1205})

		locality, err := rp.CreateLocality(&l)
		errMock := mock.ExpectationsWereMet()
		expectedErr := customerror.ErrDefaultLocalitySQL

		assert.NoError(t, errMock)
		assert.ErrorIs(t, expectedErr, err)
		assert.Error(t, err)
		assert.Empty(t, locality)
	})
}

func TestLocalitiesRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositoryLocalities(db, logMock)

	t.Run("test repository method for get locality by id successfully", func(t *testing.T) {
		ID := 1
		l := model.Locality{ID: 17, Locality: "Phoenix", Province: "Arizona", Country: "EUA"}

		rows := sqlmock.NewRows([]string{"id", "locality_name", "province_name", "country_name"}).
			AddRow(l.ID, l.Locality, l.Province, l.Country)

		mock.ExpectQuery("SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality` WHERE `id` = ?").
			WithArgs(ID).
			WillReturnRows(rows)

		result, err := rp.GetByID(ID)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, err)
		assert.NoError(t, errMock)
		assert.Equal(t, l, result)
	})

	t.Run("test repository method for get locality by id with error not found", func(t *testing.T) {
		ID := 99

		mock.ExpectQuery("SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality` WHERE `id` = ?").
			WithArgs(ID).
			WillReturnError(sql.ErrNoRows)

		seller, err := rp.GetByID(ID)
		errorExpected := customerror.NewSellerErr(customerror.ErrNotFound.Error(), http.StatusNotFound)
		errorMock := mock.ExpectationsWereMet()

		assert.NoError(t, errorMock)
		assert.Error(t, errorExpected, err)
		assert.Equal(t, model.Locality{}, seller)
	})
}

func TestLocalitiesRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositoryLocalities(db, logMock)

	t.Run("test repository method for get all localities successfully", func(t *testing.T) {
		expectedLocalities := []model.Locality{
			{ID: 1, Locality: "Denver", Province: "Colorado", Country: "EUA"},
			{ID: 2, Locality: "Phoenix", Province: "Arizona", Country: "EUA"},
		}

		rows := sqlmock.NewRows([]string{"id", "locality_name", "province_name", "country_name"})
		for _, locality := range expectedLocalities {
			rows.AddRow(locality.ID, locality.Locality, locality.Province, locality.Country)
		}

		mock.ExpectQuery("SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality`").
			WillReturnRows(rows)

		localities, err := rp.Get()
		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, err)
		assert.Equal(t, expectedLocalities, localities)
		assert.NoError(t, errMock)
	})

	t.Run("test repository method for get all localities with query error", func(t *testing.T) {
		mock.ExpectQuery("SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality`").
			WillReturnError(sql.ErrNoRows)

		localities, err := rp.Get()
		mockErr := mock.ExpectationsWereMet()

		assert.Error(t, err)
		assert.Empty(t, localities)
		assert.NoError(t, mockErr)
	})
}

func TestLocalitiesRepository_GetSellers(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositoryLocalities(db, logMock)

	t.Run("test repository method for get report all sellers successfully", func(t *testing.T) {
		expectedReport := []model.LocalitiesJSONSellers{
			{ID: "4", Locality: "Phoenix", Sellers: 5},
			{ID: "5", Locality: "Kanto", Sellers: 8},
			{ID: "6", Locality: "New York", Sellers: 9},
			{ID: "7", Locality: "Kansas", Sellers: 3},
		}

		rows := sqlmock.NewRows([]string{"id", "locality_name", "sellers_count"})
		for _, locality := range expectedReport {
			rows.AddRow(locality.ID, locality.Locality, locality.Sellers)
		}

		mock.ExpectQuery("SELECT l.id, l.locality_name, COUNT(s.locality_id) AS `sellers_count` FROM `sellers` s RIGHT JOIN `locality` l ON s.locality_id = l.id GROUP BY l.id, l.locality_name ORDER BY l.locality_name").
			WillReturnRows(rows)

		report, err := rp.GetSellers(0)
		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, err)
		assert.Equal(t, expectedReport, report)
		assert.NoError(t, errMock)
	})

	t.Run("test repository method for get report all sellers with sql no rows", func(t *testing.T) {
		mock.ExpectQuery("SELECT l.id, l.locality_name, COUNT(s.locality_id) AS `sellers_count` FROM `sellers` s RIGHT JOIN `locality` l ON s.locality_id = l.id GROUP BY l.id, l.locality_name ORDER BY l.locality_name").
			WillReturnError(sql.ErrNoRows)

		report, err := rp.GetSellers(0)
		mockErr := mock.ExpectationsWereMet()

		assert.Error(t, err)
		assert.Empty(t, report)
		assert.NoError(t, mockErr)
	})
}

func TestLocalitiesRepository_GetCarriers(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositoryLocalities(db, logMock)

	t.Run("test repository method for get report all carriers successfully", func(t *testing.T) {
		expectedReport := []model.LocalitiesJSONCarriers{
			{ID: "4", Locality: "Phoenix", Carriers: 5},
			{ID: "5", Locality: "Kanto", Carriers: 8},
			{ID: "6", Locality: "New York", Carriers: 9},
			{ID: "7", Locality: "Kansas", Carriers: 3},
		}

		rows := sqlmock.NewRows([]string{"id", "locality_name", "carriers_count"})
		for _, locality := range expectedReport {
			rows.AddRow(locality.ID, locality.Locality, locality.Carriers)
		}

		mock.ExpectQuery("SELECT l.id, l.locality_name, COUNT(c.locality_id) AS `carriers_count` FROM `carriers` c RIGHT JOIN `locality` l ON c.locality_id = l.id GROUP BY l.id, l.locality_name ORDER BY l.locality_name").
			WillReturnRows(rows)

		report, err := rp.GetCarriers(0)
		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, err)
		assert.Equal(t, expectedReport, report)
		assert.NoError(t, errMock)
	})

	t.Run("test repository method for get report all carriers with sql no rows", func(t *testing.T) {
		mock.ExpectQuery("SELECT l.id, l.locality_name, COUNT(c.locality_id) AS `carriers_count` FROM `carriers` c RIGHT JOIN `locality` l ON c.locality_id = l.id GROUP BY l.id, l.locality_name ORDER BY l.locality_name").
			WillReturnError(sql.ErrNoRows)

		report, err := rp.GetCarriers(0)
		mockErr := mock.ExpectationsWereMet()

		assert.Error(t, err)
		assert.Empty(t, report)
		assert.NoError(t, mockErr)
	})
}

func TestLocalitiesRepository_GetReportSellersWithID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositoryLocalities(db, logMock)

	t.Run("test repository method for get report sellers by ID successfully", func(t *testing.T) {
		expectedReport := []model.LocalitiesJSONSellers{
			{ID: "4", Locality: "Phoenix", Sellers: 5},
		}
		ID := 4
		l := model.Locality{ID: 4, Locality: "Phoenix", Province: "Arizona", Country: "EUA"}

		rows := sqlmock.NewRows([]string{"id", "locality_name", "province_name", "country_name"}).
			AddRow(l.ID, l.Locality, l.Province, l.Country)

		mock.ExpectQuery("SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality` WHERE `id` = ?").
			WithArgs(ID).
			WillReturnRows(rows)

		row := sqlmock.NewRows([]string{"id", "locality_name", "sellers_count"}).
			AddRow(expectedReport[0].ID, expectedReport[0].Locality, expectedReport[0].Sellers)

		mock.ExpectQuery("SELECT l.id, l.locality_name, COUNT(s.locality_id) AS `sellers_count` FROM `sellers` s RIGHT JOIN `locality` l ON s.locality_id = l.id WHERE l.id = ? GROUP BY l.id, l.locality_name").
			WithArgs(ID).
			WillReturnRows(row)

		report, err := rp.GetReportSellersWithID(ID)
		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, err)
		assert.Equal(t, expectedReport, report)
		assert.NoError(t, errMock)
	})

	t.Run("test repository method for get report sellers with ID not found", func(t *testing.T) {
		ID := 999
		expectedErr := customerror.ErrLocalityNotFound

		mock.ExpectQuery("SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality` WHERE `id` = ?").
			WithArgs(ID).
			WillReturnError(sql.ErrNoRows)

		report, err := rp.GetReportSellersWithID(ID)
		errMock := mock.ExpectationsWereMet()

		assert.ErrorIs(t, expectedErr, err)
		assert.Empty(t, report)
		assert.NoError(t, errMock)
	})
}

func TestLocalitiesRepository_GetReportCarriersWithID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := repository.CreateRepositoryLocalities(db, logMock)

	t.Run("test repository method for get report carriers by ID successfully", func(t *testing.T) {
		expectedReport := []model.LocalitiesJSONCarriers{
			{ID: "4", Locality: "Phoenix", Carriers: 5},
		}
		ID := 4
		l := model.Locality{ID: 4, Locality: "Phoenix", Province: "Arizona", Country: "EUA"}

		rows := sqlmock.NewRows([]string{"id", "locality_name", "province_name", "country_name"}).
			AddRow(l.ID, l.Locality, l.Province, l.Country)

		mock.ExpectQuery("SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality` WHERE `id` = ?").
			WithArgs(ID).
			WillReturnRows(rows)

		row := sqlmock.NewRows([]string{"id", "locality_name", "carriers_count"}).
			AddRow(expectedReport[0].ID, expectedReport[0].Locality, expectedReport[0].Carriers)

		mock.ExpectQuery("SELECT l.id, l.locality_name, COUNT(c.locality_id) AS `carriers_count` FROM `carriers` c RIGHT JOIN `locality` l ON c.locality_id = l.id WHERE l.id = ? GROUP BY l.id, l.locality_name").
			WithArgs(ID).
			WillReturnRows(row)

		report, err := rp.GetReportCarriersWithID(ID)
		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, err)
		assert.Equal(t, expectedReport, report)
		assert.NoError(t, errMock)
	})

	t.Run("test repository method for get report carriers with ID not found", func(t *testing.T) {
		ID := 999
		expectedErr := customerror.ErrLocalityNotFound

		mock.ExpectQuery("SELECT `id`, `locality_name`, `province_name`, `country_name` FROM `locality` WHERE `id` = ?").
			WithArgs(ID).
			WillReturnError(sql.ErrNoRows)

		report, err := rp.GetReportCarriersWithID(ID)
		errMock := mock.ExpectationsWereMet()

		assert.ErrorIs(t, expectedErr, err)
		assert.Empty(t, report)
		assert.NoError(t, errMock)
	})
}

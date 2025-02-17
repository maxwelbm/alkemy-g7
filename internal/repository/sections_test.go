package repository_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func TestSectionRepository_Post(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.CreateRepositorySections(db)

	t.Run("given a valid section the create it with no error", func(t *testing.T) {
		expectedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mock.ExpectExec("INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)").WithArgs(expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID).WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?").WithArgs(expectedSection.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}).AddRow(expectedSection.ID, expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID))

		section, err := rp.Post(&expectedSection)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, err)
		assert.NoError(t, errMock)
		assert.Equal(t, expectedSection, section)
	})

	t.Run("given a duplicate section then retun error", func(t *testing.T) {
		expectedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mock.ExpectExec("INSERT INTO `sections` (`section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)").WithArgs(expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID).WillReturnError(&mysql.MySQLError{
			Number:   1062,
			SQLState: [5]byte{'2', '3', '0', '0', '0'},
			Message:  "Duplicate entry",
		})

		_, err := rp.Post(&expectedSection)

		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.Error(t, err)
	})
}

func TestSectionRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.CreateRepositorySections(db)

	t.Run("given a valid id the return the section with no error", func(t *testing.T) {
		sectionID := 1
		expectedSection := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		rows := sqlmock.NewRows([]string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"}).AddRow(expectedSection.ID, expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID)

		mock.ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?").WithArgs(sectionID).WillReturnRows(rows)

		result, err := rp.GetByID(sectionID)
		assert.NoError(t, err)

		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)

		assert.Equal(t, expectedSection, result)
	})

	t.Run("given an invalid id then return error", func(t *testing.T) {
		sectionID := 99

		mock.ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id = ?").WithArgs(sectionID).WillReturnError(sql.ErrNoRows)

		section, err := rp.GetByID(sectionID)
		errorExpected := customerror.HandleError("section", customerror.ErrorNotFound, "")
		errMock := mock.ExpectationsWereMet()

		assert.NoError(t, errMock)
		assert.Error(t, errorExpected, err)
		assert.Equal(t, model.Section{}, section)
	})
}

func TestSectionRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.CreateRepositorySections(db)

	t.Run("return all the sections", func(t *testing.T) {
		expectedSections := []model.Section{
			{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1},
			{ID: 2, SectionNumber: "S02", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1},
		}

		rows := mock.NewRows([]string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity", "warehouse_id", "product_type_id"})
		for _, section := range expectedSections {
			rows.AddRow(section.ID, section.SectionNumber, section.CurrentTemperature, section.MinimumTemperature, section.CurrentCapacity, section.MinimumCapacity, section.MaximumCapacity, section.WarehouseID, section.ProductTypeID)
		}

		mock.ExpectQuery("SELECT `id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id` FROM `sections`").WillReturnRows(rows)

		sections, err := rp.Get()
		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, err)
		assert.NoError(t, errMock)
		assert.Equal(t, expectedSections, sections)
	})

	t.Run("return all the sections", func(t *testing.T) {
		mock.ExpectQuery("SELECT `id`, `section_number`, `current_temperature`, `minimum_temperature`, `current_capacity`, `minimum_capacity`, `maximum_capacity`, `warehouse_id`, `product_type_id` FROM `sections`").WillReturnError(errors.New("unmapped error"))

		sections, err := rp.Get()
		errMock := mock.ExpectationsWereMet()
		assert.Error(t, err)
		assert.NoError(t, errMock)
		assert.Equal(t, []model.Section(nil), sections)
	})
}

func TestSectionRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.CreateRepositorySections(db)

	t.Run("given a valid section then update it and return no error", func(t *testing.T) {
		sectionID := 1
		section := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}

		mock.ExpectExec("UPDATE `sections` SET `section_number` = ?, `current_temperature` = ?, `minimum_temperature` = ?, `current_capacity` = ?, `minimum_capacity` = ?, `maximum_capacity` = ?, `warehouse_id` = ?, `product_type_id` = ? WHERE `id` = ?").WithArgs(section.SectionNumber, section.CurrentTemperature, section.MinimumTemperature, section.CurrentCapacity, section.MinimumCapacity, section.MaximumCapacity, section.WarehouseID, section.ProductTypeID, section.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		_, err := rp.Update(sectionID, &section)
		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, err)
		assert.NoError(t, mockErr)
	})

	t.Run("given a duplicate section then return error", func(t *testing.T) {
		sectionID := 1
		section := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}
		expectedError := customerror.HandleError("section", customerror.ErrorConflict, "")

		mock.ExpectExec("UPDATE `sections` SET `section_number` = ?, `current_temperature` = ?, `minimum_temperature` = ?, `current_capacity` = ?, `minimum_capacity` = ?, `maximum_capacity` = ?, `warehouse_id` = ?, `product_type_id` = ? WHERE `id` = ?").WithArgs(section.SectionNumber, section.CurrentTemperature, section.MinimumTemperature, section.CurrentCapacity, section.MinimumCapacity, section.MaximumCapacity, section.WarehouseID, section.ProductTypeID, section.ID).WillReturnError(&mysql.MySQLError{Number: 1062, SQLState: [5]byte{'2', '3', '0', '0', '0'}, Message: "Duplicate entry"})

		_, err := rp.Update(sectionID, &section)
		mockErr := mock.ExpectationsWereMet()

		assert.Error(t, expectedError, err)
		assert.NoError(t, mockErr)
	})

	t.Run("return no rows error", func(t *testing.T) {
		sectionID := 1
		section := model.Section{ID: 1, SectionNumber: "S01", CurrentTemperature: 10.0, MinimumTemperature: 5.0, CurrentCapacity: 10, MinimumCapacity: 5, MaximumCapacity: 20, WarehouseID: 1, ProductTypeID: 1}
		expectedError := customerror.HandleError("section", customerror.ErrorNotFound, "")

		mock.ExpectExec("UPDATE `sections` SET `section_number` = ?, `current_temperature` = ?, `minimum_temperature` = ?, `current_capacity` = ?, `minimum_capacity` = ?, `maximum_capacity` = ?, `warehouse_id` = ?, `product_type_id` = ? WHERE `id` = ?").WillReturnError(sql.ErrNoRows)

		_, err := rp.Update(sectionID, &section)
		mockErr := mock.ExpectationsWereMet()

		assert.Error(t, err)
		assert.NoError(t, mockErr)
		assert.Equal(t, expectedError, err)

	})
}

func TestSectionRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.CreateRepositorySections(db)

	t.Run("given a valid section then delete the section", func(t *testing.T) {
		sectionID := 1

		mock.ExpectExec("DELETE FROM `sections` WHERE `id` = ?").WithArgs(sectionID).WillReturnResult(sqlmock.NewResult(0, 1))

		err := rp.Delete(sectionID)

		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, mockErr)
		assert.NoError(t, err)
	})

	t.Run("return a error if section not found", func(t *testing.T) {
		sectionID := 99

		expectedErr := customerror.HandleError("section", customerror.ErrorNotFound, "")

		mock.ExpectExec("DELETE FROM `sections` WHERE `id` = ?").WithArgs(sectionID).WillReturnError(sql.ErrNoRows)

		err := rp.Delete(sectionID)

		mockErr := mock.ExpectationsWereMet()

		assert.NoError(t, mockErr)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestCountProductBatchesBySectionID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.CreateRepositorySections(db)

	t.Run("given a valid product batch then return the count with no error", func(t *testing.T) {
		sectionID := 1
		expectedCount := model.SectionProductBatches{ID: 1, SectionNumber: "S01", ProductsCount: 3}

		rows := sqlmock.NewRows([]string{"id", "section_number", "products_count"}).AddRow(1, "S01", 3)

		mock.ExpectQuery("SELECT s.id, s.section_number, COUNT(pb.section_id) as products_count FROM sections s INNER JOIN product_batches pb ON pb.section_id = s.id WHERE s.id = ? GROUP BY s.id").WithArgs(sectionID).WillReturnRows(rows)

		count, err := rp.CountProductBatchesBySectionID(sectionID)

		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)
		assert.NoError(t, err)
		assert.Equal(t, expectedCount, count)

	})

	t.Run("given an invalid product batch then return error", func(t *testing.T) {
		sectionID := 99
		expectedCount := model.SectionProductBatches{}

		mock.ExpectQuery("SELECT s.id, s.section_number, COUNT(pb.section_id) as products_count FROM sections s INNER JOIN product_batches pb ON pb.section_id = s.id WHERE s.id = ? GROUP BY s.id").WillReturnError(sql.ErrNoRows)

		count, err := rp.CountProductBatchesBySectionID(sectionID)

		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)
		assert.Error(t, err)
		assert.Equal(t, expectedCount, count)

	})
}

func TestCountProductBatchesSections(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rp := repository.CreateRepositorySections(db)

	t.Run("given a slice of product batches then count it and return with no error", func(t *testing.T) {
		expectedCount := []model.SectionProductBatches{
			{ID: 1, SectionNumber: "S01", ProductsCount: 3},
			{ID: 2, SectionNumber: "S02", ProductsCount: 5},
		}

		rows := sqlmock.NewRows([]string{"id", "section_number", "products_count"}).
			AddRow(1, "S01", 3).
			AddRow(2, "S02", 5)

		mock.ExpectQuery("SELECT s.id, s.section_number, COUNT(pb.section_id) as products_count FROM sections s INNER JOIN product_batches pb ON pb.section_id = s.id GROUP BY s.id").WillReturnRows(rows)

		countPB, err := rp.CountProductBatchesSections()

		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)
		assert.NoError(t, err)
		assert.Equal(t, expectedCount, countPB)
	})

	t.Run("given an invalid product batches then return error", func(t *testing.T) {
		expectedCount := []model.SectionProductBatches(nil)

		mock.ExpectQuery("SELECT s.id, s.section_number, COUNT(pb.section_id) as products_count FROM sections s INNER JOIN product_batches pb ON pb.section_id = s.id GROUP BY s.id").WillReturnError(errors.New("unmapped error"))

		countPB, err := rp.CountProductBatchesSections()

		errMock := mock.ExpectationsWereMet()
		assert.NoError(t, errMock)
		assert.Error(t, err)
		assert.Equal(t, expectedCount, countPB)

	})
}

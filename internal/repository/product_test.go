package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	appErr "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func TestProductRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rp := NewProductRepository(db, logMock)

	t.Run("retrieving all products", func(t *testing.T) {
		products := map[int]model.Product{
			1: {ID: 1, ProductCode: "CODE1", Description: "Product 1", Width: 10.5, Height: 20.5, Length: 30.5, NetWeight: 100, ExpirationRate: 0.5, RecommendedFreezingTemperature: -18, FreezingRate: 0.3, ProductTypeID: 1, SellerID: 1},
			2: {ID: 2, ProductCode: "CODE2", Description: "Product 2", Width: 20.5, Height: 30.5, Length: 40.5, NetWeight: 200, ExpirationRate: 0.6, RecommendedFreezingTemperature: -16, FreezingRate: 0.4, ProductTypeID: 2, SellerID: 2},
		}

		rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"})
		for _, p := range products {
			rows.AddRow(p.ID, p.ProductCode, p.Description, p.Width, p.Height, p.Length, p.NetWeight, p.ExpirationRate, p.RecommendedFreezingTemperature, p.FreezingRate, p.ProductTypeID, p.SellerID)
		}

		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products").
			WillReturnRows(rows)

		result, err := rp.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, products, result)
	})

	t.Run("error executing query", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products").
			WillReturnRows(sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"}).
				AddRow(1, "CODE1", "Product 1", 10.5, 20.5, 30.5, 100, 0.5, -18, 0.3, 1, 1).RowError(0, fmt.Errorf("Row error")))

		_, err := rp.GetAll()

		fmt.Println(err)

		assert.EqualError(t, errors.New("Row error"), err.Error())

	})

	t.Run("error on scanning product", func(t *testing.T) {
		expected := errors.New("Error executing query")
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products").
			WillReturnError(expected)

		_, err := rp.GetAll()

		assert.EqualError(t, expected, err.Error())
	})

	t.Run("error on convert type of product", func(t *testing.T) {
		expected := errors.New("sql: Scan error on column index 3, name \"width\": converting driver.Value type string (\"Invalid Field\") to a float64: invalid syntax")
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products").
			WillReturnRows(sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"}).
				AddRow(1, "CODE1", "Product 1", "Invalid Field", 20.5, 30.5, 100, 0.5, -18, 0.3, 1, 1))

		_, err := rp.GetAll()

		assert.EqualError(t, expected, err.Error())
	})

	t.Run("error on rows.Err()", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products").
			WillReturnRows(sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"}).
				AddRow(1, "CODE1", "Product 1", 10.5, 20.5, 30.5, 100, 0.5, -18, 0.3, 1, 1).RowError(0, fmt.Errorf("Row error")))

		_, err := rp.GetAll()

		fmt.Println(err)

		assert.EqualError(t, errors.New("Row error"), err.Error())

	})
}

func TestProductRepository_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db, logMock)

	t.Run("retrieving an existing product by ID", func(t *testing.T) {
		productID := 1
		expectedProduct := model.Product{
			ID:                             productID,
			ProductCode:                    "CODE1",
			Description:                    "Product 1",
			Width:                          10.5,
			Height:                         20.5,
			Length:                         30.5,
			NetWeight:                      100,
			ExpirationRate:                 0.5,
			RecommendedFreezingTemperature: -18,
			FreezingRate:                   0.3,
			ProductTypeID:                  1,
			SellerID:                       1,
		}

		rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"}).
			AddRow(expectedProduct.ID, expectedProduct.ProductCode, expectedProduct.Description, expectedProduct.Width, expectedProduct.Height, expectedProduct.Length, expectedProduct.NetWeight, expectedProduct.ExpirationRate, expectedProduct.RecommendedFreezingTemperature, expectedProduct.FreezingRate, expectedProduct.ProductTypeID, expectedProduct.SellerID)

		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products WHERE id = ?").
			WithArgs(productID).
			WillReturnRows(rows)

		result, err := repo.GetByID(productID)
		assert.NoError(t, err)
		assert.Equal(t, expectedProduct, result)
	})

	t.Run("product not found", func(t *testing.T) {
		productID := 100

		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products WHERE id = ?").
			WithArgs(productID).
			WillReturnError(sql.ErrNoRows)

		_, err := repo.GetByID(productID)
		fmt.Println(err)

		assert.Error(t, err)
		assert.EqualError(t, appErr.HandleError("product", appErr.ErrorNotFound, ""), err.Error())

	})

	t.Run("error scanning product", func(t *testing.T) {
		productID := 1

		mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products WHERE id = ?").
			WithArgs(productID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"}).
				AddRow(productID, "CODE1", "Product 1", 10.5, 20.5, 30.5, 100, 0.5, -18, 0.3, 1, "INVALID_TYPE"))

		_, err := repo.GetByID(productID)
		assert.Error(t, err)
	})
}

func TestProductRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db, logMock)

	product := model.Product{
		ProductCode:                    "CODE1",
		Description:                    "Product 1",
		Width:                          10.5,
		Height:                         20.5,
		Length:                         30.5,
		NetWeight:                      100,
		ExpirationRate:                 0.5,
		RecommendedFreezingTemperature: -18,
		FreezingRate:                   0.3,
		ProductTypeID:                  1,
		SellerID:                       1,
	}

	t.Run("successful creation of a product", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO products (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)").
			WithArgs(product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature, product.FreezingRate, product.ProductTypeID, product.SellerID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		newProduct, err := repo.Create(product)
		assert.NoError(t, err)
		assert.Equal(t, 1, newProduct.ID)
	})

	t.Run("error on insert execution", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO products (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)").
			WithArgs(product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature, product.FreezingRate, product.ProductTypeID, product.SellerID).
			WillReturnError(errors.New("simulated database error"))

		_, err := repo.Create(product)
		fmt.Println(err)
		assert.Error(t, err)
	})

	t.Run("error on getting last insert id", func(t *testing.T) {
		mock.ExpectExec("INSERT INTO products (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)").
			WithArgs(product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature, product.FreezingRate, product.ProductTypeID, product.SellerID).
			WillReturnResult(sqlmock.NewErrorResult(errors.New("Error to get id")))

		_, err := repo.Create(product)
		assert.EqualError(t, errors.New("Error to get id"), err.Error())
	})

}

func TestProductRepository_Update(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db, logMock)

	product := model.Product{
		ProductCode:                    "UPDATED_CODE",
		Description:                    "Updated Product",
		Width:                          15.0,
		Height:                         25.0,
		Length:                         35.0,
		NetWeight:                      150,
		ExpirationRate:                 0.7,
		RecommendedFreezingTemperature: -15,
		FreezingRate:                   0.4,
		ProductTypeID:                  2,
		SellerID:                       3,
	}

	t.Run("successful update of a product", func(t *testing.T) {
		productID := 1

		mock.ExpectExec("UPDATE products SET product_code = ?, description = ?, width = ?, height = ?, length = ?, net_weight = ?, expiration_rate = ?, recommended_freezing_temperature = ?, freezing_rate = ?, product_type_id = ?, seller_id = ? WHERE id = ?").
			WithArgs(product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature, product.FreezingRate, product.ProductTypeID, product.SellerID, productID).
			WillReturnResult(sqlmock.NewResult(1, 1)) // Simula a atualização bem-sucedida

		updatedProduct, err := repo.Update(productID, product)
		assert.NoError(t, err)
		assert.Equal(t, productID, updatedProduct.ID) // Verifica se o ID retornado é o esperado
	})

	t.Run("error while updating product", func(t *testing.T) {
		productID := 1

		mock.ExpectExec("UPDATE products SET product_code = ?, description = ?, width = ?, height = ?, length = ?, net_weight = ?, expiration_rate = ?, recommended_freezing_temperature = ?, freezing_rate = ?, product_type_id = ?, seller_id = ? WHERE id = ?").
			WithArgs(product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.RecommendedFreezingTemperature, product.FreezingRate, product.ProductTypeID, product.SellerID, productID).
			WillReturnError(errors.New("simulated update error"))

		_, err := repo.Update(productID, product)

		assert.Error(t, err)
	})
}

func TestProductRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewProductRepository(db, logMock)

	t.Run("successful deletion of a product", func(t *testing.T) {
		productID := 1

		mock.ExpectExec("DELETE FROM products WHERE id = ?").
			WithArgs(productID).
			WillReturnResult(sqlmock.NewResult(1, 1)) // Simula a deleção bem-sucedida

		err := repo.Delete(productID)
		assert.NoError(t, err)
	})

	t.Run("error while deleting a product", func(t *testing.T) {
		productID := 100

		mock.ExpectExec("DELETE FROM products WHERE id = ?").
			WithArgs(productID).
			WillReturnError(errors.New("simulated delete error"))

		err := repo.Delete(productID)
		assert.Error(t, err)
		assert.Equal(t, appErr.HandleError("product", appErr.ErrorDep, "product record"), err)
	})
}

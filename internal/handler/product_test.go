package handler_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProductsHandler(t *testing.T) {

	t.Run("GetAllProducts - should return list of products", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("GetAllProducts").Return([]model.Product{
			{ID: 1, ProductCode: "P001", Description: "Product 1", Width: 10, Height: 20, Length: 0, NetWeight: 0, ExpirationRate: 0, RecommendedFreezingTemperature: 0, FreezingRate: 0, ProductTypeID: 0, SellerID: 0},
			{ID: 2, ProductCode: "P002", Description: "Product 2", Width: 15, Height: 25, Length: 0, NetWeight: 0, ExpirationRate: 0, RecommendedFreezingTemperature: 0, FreezingRate: 0, ProductTypeID: 0, SellerID: 0},
		}, nil)

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		req := httptest.NewRequest("GET", "/api/v1/products", nil)
		res := httptest.NewRecorder()

		productHd.GetAllProducts(res, req)

		expected := `{"data":[{"id":1,"product_code":"P001","description":"Product 1","width":10,"height":20,"length":0,"net_weight":0,"expiration_rate":0,"recommended_freezing_temperature":0,"freezing_rate":0,"product_type_id":0,"seller_id":0},{"id":2,"product_code":"P002","description":"Product 2","width":15,"height":25,"length":0,"net_weight":0,"expiration_rate":0,"recommended_freezing_temperature":0,"freezing_rate":0,"product_type_id":0,"seller_id":0}]}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.JSONEq(t, expected, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

	t.Run("GetAllProducts - Return error", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("GetAllProducts").Return([]model.Product{}, errors.New("error to get all products"))

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		req := httptest.NewRequest("GET", "/api/v1/products", nil)
		res := httptest.NewRecorder()

		productHd.GetAllProducts(res, req)

		expected := "{\"message\":\"error to get all products\"}"

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Contains(t, expected, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

}

func TestGetProductById(t *testing.T) {
	t.Run("GetProductByID - Success", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("GetProductByID", 1).Return(model.Product{
			ID:                             1,
			ProductCode:                    "P001",
			Description:                    "Product 1",
			Width:                          10,
			Height:                         20,
			Length:                         0,
			NetWeight:                      0,
			ExpirationRate:                 0,
			RecommendedFreezingTemperature: 0,
			FreezingRate:                   0,
			ProductTypeID:                  0,
			SellerID:                       0,
		}, nil)

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Get("/api/v1/products/{id}", productHd.GetProductByID)

		req := httptest.NewRequest("GET", "/api/v1/products/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"data":{"id":1,"product_code":"P001","description":"Product 1","width":10,"height":20,"length":0,"net_weight":0,"expiration_rate":0,"recommended_freezing_temperature":0,"freezing_rate":0,"product_type_id":0,"seller_id":0}}`

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Contains(t, expected, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

	t.Run("GetProductByID - Error in Id", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Get("/api/v1/products/{id}", productHd.GetProductByID)

		req := httptest.NewRequest("GET", "/api/v1/products/A", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := "{\"message\":\"invalid id\"}"

		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Contains(t, expected, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

	t.Run("GetProductByID - Error when getting a product", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("GetProductByID", 1).Return(model.Product{}, errors.New("product not found"))

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Get("/api/v1/products/{id}", productHd.GetProductByID)

		req := httptest.NewRequest("GET", "/api/v1/products/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := "{\"message\":\"product not found\"}"

		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Contains(t, expected, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

}

func TestInsertProduct(t *testing.T) {
	t.Run("Create Product - Success", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("CreateProduct", mock.Anything).Return(model.Product{ID: 1, ProductCode: "P003", Description: "New Product", Width: 1, Height: 1, Length: 1, NetWeight: 1, ExpirationRate: 1, RecommendedFreezingTemperature: 1, FreezingRate: 1, ProductTypeID: 1, SellerID: 1}, nil)

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		req := httptest.NewRequest("POST", "/api/v1/products", strings.NewReader(`{
			"product_code": "P003",
			"description": "New Product",
			"width": 1,
			"height": 1,
			"length": 1,
			"net_weight": 1,
			"expiration_rate": 1,
			"recommended_freezing_temperature":1,
			"freezing_rate": 1,
			"product_type_id": 1,
			"seller_id": 1
		}`))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		productHd.CreateProduct(res, req)

		expected := `{"data":{"id":1,"product_code":"P003","description":"New Product","width":1,"height":1,"length":1,"net_weight":1,"expiration_rate":1,"recommended_freezing_temperature":1,"freezing_rate":1,"product_type_id":1,"seller_id":1}}`

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.JSONEq(t, expected, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})
	t.Run("Create Product - Error in Body", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		req := httptest.NewRequest("POST", "/api/v1/products", strings.NewReader(`{
			"product_code": 05`))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		productHd.CreateProduct(res, req)

		expected := `{"message":"invalid json syntax"}`

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.JSONEq(t, expected, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

	t.Run("Create Product - Success", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("CreateProduct", mock.Anything).Return(model.Product{}, errors.New("product error"))

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		req := httptest.NewRequest("POST", "/api/v1/products", strings.NewReader(`{
			"product_code": "P003",
			"description": "New Product",
			"width": 1,
			"height": 1,
			"length": 1,
			"net_weight": 1,
			"expiration_rate": 1,
			"recommended_freezing_temperature":1,
			"freezing_rate": 1,
			"product_type_id": 1,
			"seller_id": 1
		}`))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		productHd.CreateProduct(res, req)

		expected := `{"message":"product error"}`

		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.JSONEq(t, expected, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

}

func TestUpdateProduct(t *testing.T) {
	t.Run("Update - Success", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("UpdateProduct", 1, mock.Anything).Return(model.Product{ID: 1, ProductCode: "P003", Description: "Updated Product", Width: 0, Height: 0, Length: 0, NetWeight: 0, ExpirationRate: 0, RecommendedFreezingTemperature: 0, FreezingRate: 0, ProductTypeID: 0, SellerID: 0}, nil)

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Patch("/api/v1/products/{id}", productHd.UpdateProduct)

		req := httptest.NewRequest("PATCH", "/api/v1/products/1", strings.NewReader(`{
			"product_code": "P003",
			"description": "Updated Product"
		}`))
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := `{"data":{"id":1,"product_code":"P003","description":"Updated Product","width":0,"height":0,"length":0,"net_weight":0,"expiration_rate":0,"recommended_freezing_temperature":0,"freezing_rate":0,"product_type_id":0,"seller_id":0}}`
		assert.Equal(t, http.StatusOK, res.Code)
		assert.JSONEq(t, expected, res.Body.String())
	})

	t.Run("Update - Not found", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("UpdateProduct", 2, mock.Anything).Return(model.Product{}, customerror.HandleError("product", customerror.ErrorNotFound, ""))

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Patch("/api/v1/products/{id}", productHd.UpdateProduct)

		req := httptest.NewRequest("PATCH", "/api/v1/products/2", strings.NewReader(`{
			"product_code": "P002"
		}`))

		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expectedError := `{"message":"product not found"}`
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.JSONEq(t, expectedError, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})
	t.Run("Update - Error body", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Patch("/api/v1/products/{id}", productHd.UpdateProduct)

		req := httptest.NewRequest("PATCH", "/api/v1/products/2", strings.NewReader(`{
			"product_code": 123
		}`))

		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expectedError := `{"message":"invalid json syntax"}`
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.JSONEq(t, expectedError, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

	t.Run("Update - Invalid id", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Patch("/api/v1/products/{id}", productHd.UpdateProduct)

		req := httptest.NewRequest("PATCH", "/api/v1/products/a", strings.NewReader(`{
			"product_code": 123
		}`))

		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := "{\"message\":\"invalid id\"}"
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.JSONEq(t, expected, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

}

func TestDeleteProduct(t *testing.T) {
	t.Run("DeleteProduct - Success", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("DeleteProduct", 1).Return(nil)

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Delete("/api/v1/products/{id}", productHd.DeleteProductByID)

		req := httptest.NewRequest("DELETE", "/api/v1/products/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("DeleteProduct - Error id invalid", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Delete("/api/v1/products/{id}", productHd.DeleteProductByID)

		req := httptest.NewRequest("DELETE", "/api/v1/products/q", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expected := "{\"message\":\"invalid id\"}"
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.JSONEq(t, expected, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

	t.Run("DeleteProduct - Error not found", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("DeleteProduct", 2).Return(customerror.HandleError("product", customerror.ErrorNotFound, ""))

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Delete("/api/v1/products/{id}", productHd.DeleteProductByID)

		req := httptest.NewRequest("DELETE", "/api/v1/products/2", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expectedError := `{"message":"product not found"}`
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.JSONEq(t, expectedError, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})
	t.Run("DeleteProduct - Error generic", func(t *testing.T) {
		productServiceMock := new(mocks.MockIProductService)

		productServiceMock.On("DeleteProduct", 2).Return(errors.New("generic error"))

		productHd := handler.NewProductHandler(productServiceMock, logMock)

		r := chi.NewRouter()
		r.Delete("/api/v1/products/{id}", productHd.DeleteProductByID)

		req := httptest.NewRequest("DELETE", "/api/v1/products/2", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expectedError := `{"message":"Internal Server Error"}`
		assert.Equal(t, http.StatusInternalServerError, res.Code)
		assert.JSONEq(t, expectedError, res.Body.String())

		productServiceMock.AssertExpectations(t)
	})

}

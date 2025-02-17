package service

import (
	"errors"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductRecService_CreateProductRecords(t *testing.T) {
	product := model.ProductRecords{
		ProductID:     1,
		ID:            1,
		PurchasePrice: 11.9,
		SalePrice:     32.4,
	}
	t.Run("Success create a product rec", func(t *testing.T) {
		productRecRepo := new(mocks.MockIProductRecRepository)
		productSv := new(mocks.MockIProductService)
		sv := NewProductRecService(productRecRepo, productSv)

		pId := 1

		productSv.On("GetProductByID", pId).Return(model.Product{}, nil)

		productRecRepo.On("Create", mock.Anything).Return(product, nil)

		res, err := sv.CreateProductRecords(product)

		assert.NoError(t, err)
		assert.Equal(t, product.PurchasePrice, res.PurchasePrice)
		assert.Equal(t, product.ID, res.ID)
		assert.Equal(t, product.SalePrice, res.SalePrice)
	})

	t.Run("Error validation a product rec", func(t *testing.T) {
		productRecRepo := new(mocks.MockIProductRecRepository)
		productSv := new(mocks.MockIProductService)
		sv := NewProductRecService(productRecRepo, productSv)

		product.PurchasePrice = 0

		_, err := sv.CreateProductRecords(product)

		assert.EqualError(t, err, "product record had errors: Purchase price is invalid")
	})

	t.Run("Error not found", func(t *testing.T) {
		productRecRepo := new(mocks.MockIProductRecRepository)
		productSv := new(mocks.MockIProductService)
		sv := NewProductRecService(productRecRepo, productSv)

		product.PurchasePrice = 11.0

		productSv.On("GetProductByID", mock.Anything).Return(model.Product{}, errors.New("product not found"))

		_, err := sv.CreateProductRecords(product)

		assert.EqualError(t, err, "product not found")
	})

	t.Run("Error in creation product", func(t *testing.T) {
		productRecRepo := new(mocks.MockIProductRecRepository)
		productSv := new(mocks.MockIProductService)
		sv := NewProductRecService(productRecRepo, productSv)

		product.PurchasePrice = 11.0

		productSv.On("GetProductByID", mock.Anything).Return(model.Product{}, nil)
		productRecRepo.On("Create", mock.Anything).Return(model.ProductRecords{}, errors.New("error creating product"))

		_, err := sv.CreateProductRecords(product)

		assert.EqualError(t, err, "error creating product")
	})

}

func TestProductRecService_GetProductRecordByID(t *testing.T) {
	product := model.ProductRecords{
		ProductID:     1,
		ID:            1,
		PurchasePrice: 11.9,
		SalePrice:     32.4,
	}

	t.Run("Sucess getting product rec", func(t *testing.T) {
		productRecRepo := new(mocks.MockIProductRecRepository)
		productSv := new(mocks.MockIProductService)
		sv := NewProductRecService(productRecRepo, productSv)

		productRecRepo.On("GetByID", mock.Anything).Return(product, nil)

		res, err := sv.GetProductRecordByID(1)

		assert.NoError(t, err)
		assert.Equal(t, product, res)
	})

	t.Run("Error getting product rec", func(t *testing.T) {
		productRecRepo := new(mocks.MockIProductRecRepository)
		productSv := new(mocks.MockIProductService)
		sv := NewProductRecService(productRecRepo, productSv)

		productRecRepo.On("GetByID", mock.Anything).Return(model.ProductRecords{}, errors.New("Not found"))

		_, err := sv.GetProductRecordByID(1)

		assert.EqualError(t, err, "Not found")
	})
}

func TestProductRecService_GetProductRecordReport(t *testing.T) {
	t.Run("Success getting filtered reports by product ID", func(t *testing.T) {
		productRecRepo := new(mocks.MockIProductRecRepository)
		productSv := new(mocks.MockIProductService)
		sv := NewProductRecService(productRecRepo, productSv)

		idProduct := 1

		mockReports := []model.ProductRecordsReport{
			{ProductID: 1, Description: "Product A", RecordsCount: 2},
			{ProductID: 2, Description: "Product B", RecordsCount: 3},
		}

		productRecRepo.On("GetAllReport").Return(mockReports, nil)
		productSv.On("GetProductByID", idProduct).Return(model.Product{}, nil)

		res, err := sv.GetProductRecordReport(idProduct)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(res))                    
		assert.Equal(t, 1, res[0].ProductID)            
		assert.Equal(t, "Product A", res[0].Description)
	})

	t.Run("Success getting all reports when product ID is 0", func(t *testing.T) {
		productRecRepo := new(mocks.MockIProductRecRepository)
		productSv := new(mocks.MockIProductService)
		sv := NewProductRecService(productRecRepo, productSv)

		idProduct := 0
		mockReports := []model.ProductRecordsReport{
			{ProductID: 1, Description: "Product A", RecordsCount: 2},
			{ProductID: 2, Description: "Product B", RecordsCount: 3},
		}
		productRecRepo.On("GetAllReport").Return(mockReports, nil)

		res, err := sv.GetProductRecordReport(idProduct)

		assert.NoError(t, err)
		assert.Equal(t, mockReports, res)
	})

	t.Run("Error when calling GetAllReport", func(t *testing.T) {
		productRecRepo := new(mocks.MockIProductRecRepository)
		productSv := new(mocks.MockIProductService)
		sv := NewProductRecService(productRecRepo, productSv)

		idProduct := 1
		productRecRepo.On("GetAllReport").Return(nil, assert.AnError)

		res, err := sv.GetProductRecordReport(idProduct)

		assert.Error(t, err)
		assert.Nil(t, res)
	})

	t.Run("Error when getting product by ID", func(t *testing.T) {
		productRecRepo := new(mocks.MockIProductRecRepository)
		productSv := new(mocks.MockIProductService)
		sv := NewProductRecService(productRecRepo, productSv)

		idProduct := 1
		mockReports := []model.ProductRecordsReport{
			{ProductID: 1, Description: "Product A", RecordsCount: 2},
			{ProductID: 2, Description: "Product B", RecordsCount: 3},
		}
		productRecRepo.On("GetAllReport").Return(mockReports, nil)
		productSv.On("GetProductByID", idProduct).Return(model.Product{}, assert.AnError)

		res, err := sv.GetProductRecordReport(idProduct)

		assert.Error(t, err)
		assert.Equal(t, 0, len(res))
	})

}

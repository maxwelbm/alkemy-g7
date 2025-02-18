package service_test

import (
	"errors"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func loadDependencies() *service.ProductService {
	productRepoMock := new(mocks.MockIProductsRepo)
	sellerRepositoryMock := new(mocks.MockISellerRepo)
	productServiceMock := service.NewProductService(productRepoMock, sellerRepositoryMock, logMock)
	return productServiceMock
}

func TestGetAllProducts(t *testing.T) {
	t.Run("should return the list of products", func(t *testing.T) {
		productService := loadDependencies()
		data := make(map[int]model.Product)
		data[1] = model.Product{
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
		}

		expectedValue := []model.Product{{
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
		}}

		mockRepo := productService.ProductRepository.(*mocks.MockIProductsRepo)

		mockRepo.On("GetAll").Return(data, nil)

		productList, err := productService.GetAllProducts()

		assert.NoError(t, err)
		assert.Equal(t, expectedValue, productList)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error to get a product", func(t *testing.T) {
		productService := loadDependencies()
		expectedError := errors.New("error to get a product")
		mockRepo := productService.ProductRepository.(*mocks.MockIProductsRepo)

		mockRepo.On("GetAll").Return(map[int]model.Product{}, errors.New("error to get a product"))

		_, err := productService.GetAllProducts()

		assert.EqualError(t, expectedError, err.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestGetProductByID(t *testing.T) {
	expectedProduct := model.Product{
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
	}

	t.Run("should return the product", func(t *testing.T) {
		productService := loadDependencies()
		mockRepo := productService.ProductRepository.(*mocks.MockIProductsRepo)
		mockRepo.On("GetByID", 1).Return(expectedProduct, nil)

		product, err := productService.GetProductByID(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedProduct, product)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return not found error", func(t *testing.T) {
		productService := loadDependencies()

		id := 2
		mockRepo := productService.ProductRepository.(*mocks.MockIProductsRepo)
		expectedError := customerror.HandleError("product", customerror.ErrorNotFound, "")

		mockRepo.On("GetByID", id).Return(model.Product{}, expectedError)

		product, err := productService.GetProductByID(id)

		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, model.Product{}, product)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteById(t *testing.T) {
	data := model.Product{
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
	}

	t.Run("should delete the product", func(t *testing.T) {
		productService := loadDependencies()
		mockRepo := productService.ProductRepository.(*mocks.MockIProductsRepo)
		mockRepo.On("GetByID", 1).Return(data, nil)
		mockRepo.On("Delete", 1).Return(nil)

		err := productService.DeleteProduct(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return not found error", func(t *testing.T) {
		productService := loadDependencies()
		id := 2
		mockRepo := productService.ProductRepository.(*mocks.MockIProductsRepo)
		expectedError := customerror.HandleError("product", customerror.ErrorNotFound, "")
		mockRepo.On("GetByID", id).Return(model.Product{}, expectedError)

		err := productService.DeleteProduct(id)

		assert.Equal(t, err, expectedError)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should error in delete product", func(t *testing.T) {
		expectedError := errors.New("error in delete product")
		productService := loadDependencies()
		mockRepo := productService.ProductRepository.(*mocks.MockIProductsRepo)
		mockRepo.On("GetByID", 1).Return(data, nil)
		mockRepo.On("Delete", 1).Return(errors.New("error in delete product"))

		err := productService.DeleteProduct(1)

		assert.EqualError(t, err, expectedError.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateProduct(t *testing.T) {
	listOfProducts := make(map[int]model.Product)
	listOfProducts[1] = model.Product{
		ID:                             1,
		ProductCode:                    "P001",
		Description:                    "Product 1",
		Width:                          10,
		Height:                         20,
		Length:                         1,
		NetWeight:                      1,
		ExpirationRate:                 1,
		RecommendedFreezingTemperature: 1,
		FreezingRate:                   1,
		ProductTypeID:                  1,
		SellerID:                       1,
	}

	dataProduct := model.Product{
		ID:                             1,
		ProductCode:                    "P002",
		Description:                    "Product 1",
		Width:                          10,
		Height:                         20,
		Length:                         1,
		NetWeight:                      1,
		ExpirationRate:                 1,
		RecommendedFreezingTemperature: 1,
		FreezingRate:                   1,
		ProductTypeID:                  1,
		SellerID:                       1,
	}

	dataSeller := model.Seller{
		ID: 1,
	}

	t.Run("Should return success to create product", func(t *testing.T) {
		productService := loadDependencies()

		productRepoMock := productService.ProductRepository.(*mocks.MockIProductsRepo)
		sellerRepoMock := productService.SellerRepository.(*mocks.MockISellerRepo)

		sellerRepoMock.On("GetByID", dataProduct.SellerID).
			Return(dataSeller, nil)
		productRepoMock.On("GetAll").Return(listOfProducts, nil)
		productRepoMock.On("Create", dataProduct).Return(dataProduct, nil)

		product, err := productService.CreateProduct(dataProduct)

		assert.NoError(t, err)
		assert.Equal(t, dataProduct, product)
		productRepoMock.AssertExpectations(t)
		sellerRepoMock.AssertExpectations(t)
	})

	t.Run("Should return validation error", func(t *testing.T) {
		productService := loadDependencies()

		productRepoMock := productService.ProductRepository.(*mocks.MockIProductsRepo)
		sellerRepoMock := productService.SellerRepository.(*mocks.MockISellerRepo)

		invalidProduct := model.Product{
			ID:                             1,
			ProductCode:                    "",
			Description:                    "Product 1",
			Width:                          10,
			Height:                         20,
			Length:                         1,
			NetWeight:                      1,
			ExpirationRate:                 1,
			RecommendedFreezingTemperature: 1,
			FreezingRate:                   1,
			ProductTypeID:                  1,
			SellerID:                       1,
		}

		sellerRepoMock.On("GetByID", invalidProduct.SellerID).
			Return(dataSeller, nil)
		productRepoMock.On("GetAll").Return(listOfProducts, nil)

		product, err := productService.CreateProduct(invalidProduct)

		assert.EqualError(t, err, "validation errors: ProductCode is required")
		assert.Equal(t, model.Product{}, product)
	})

	t.Run("Should return validation error in get id", func(t *testing.T) {
		productService := loadDependencies()

		productRepoMock := productService.ProductRepository.(*mocks.MockIProductsRepo)
		sellerRepoMock := productService.SellerRepository.(*mocks.MockISellerRepo)

		invalidProduct := model.Product{
			ID:                             1,
			ProductCode:                    "",
			Description:                    "Product 1",
			Width:                          10,
			Height:                         20,
			Length:                         1,
			NetWeight:                      1,
			ExpirationRate:                 1,
			RecommendedFreezingTemperature: 1,
			FreezingRate:                   1,
			ProductTypeID:                  1,
			SellerID:                       1,
		}

		sellerRepoMock.On("GetByID", invalidProduct.SellerID).
			Return(dataSeller, nil)
		productRepoMock.On("GetAll").Return(listOfProducts, nil)

		product, err := productService.CreateProduct(invalidProduct)

		assert.EqualError(t, err, "validation errors: ProductCode is required")
		assert.Equal(t, model.Product{}, product)
	})

	t.Run("Should return exist by product code", func(t *testing.T) {
		productService := loadDependencies()

		productRepoMock := productService.ProductRepository.(*mocks.MockIProductsRepo)
		sellerRepoMock := productService.SellerRepository.(*mocks.MockISellerRepo)

		sellerRepoMock.On("GetByID", dataProduct.SellerID).
			Return(dataSeller, nil)
		productRepoMock.On("GetAll").Return(listOfProducts, nil)

		product, err := productService.CreateProduct(listOfProducts[1])

		assert.Equal(t, customerror.CustomError{Object: listOfProducts[1].ProductCode, Err: customerror.ErrConflict}, err)
		assert.Equal(t, model.Product{}, product)
		productRepoMock.AssertExpectations(t)
		sellerRepoMock.AssertExpectations(t)
	})

	t.Run("Should return error seller not found", func(t *testing.T) {
		expectedErr := errors.New("seller not found")
		productService := loadDependencies()

		sellerRepoMock := productService.SellerRepository.(*mocks.MockISellerRepo)

		sellerRepoMock.On("GetByID", dataProduct.SellerID).
			Return(model.Seller{}, errors.New("seller not found"))

		_, err := productService.CreateProduct(dataProduct)

		assert.EqualError(t, expectedErr, err.Error())
		sellerRepoMock.AssertExpectations(t)
	})

	t.Run("Should return error in create product", func(t *testing.T) {
		expectedError := errors.New("error to create product")
		productService := loadDependencies()

		productRepoMock := productService.ProductRepository.(*mocks.MockIProductsRepo)
		sellerRepoMock := productService.SellerRepository.(*mocks.MockISellerRepo)

		sellerRepoMock.On("GetByID", dataProduct.SellerID).
			Return(dataSeller, nil)
		productRepoMock.On("GetAll").Return(listOfProducts, nil)
		productRepoMock.On("Create", dataProduct).Return(model.Product{}, errors.New("error to create product"))

		_, err := productService.CreateProduct(dataProduct)

		assert.EqualError(t, expectedError, err.Error())
		productRepoMock.AssertExpectations(t)
		sellerRepoMock.AssertExpectations(t)
	})
}

func TestUpdateProducts(t *testing.T) {
	productService := loadDependencies()

	t.Run("Should return success to update product", func(t *testing.T) {
		prm := productService.ProductRepository.(*mocks.MockIProductsRepo)
		srm := productService.SellerRepository.(*mocks.MockISellerRepo)

		inputProduct := model.Product{
			ID:                             1,
			ProductCode:                    "P002",
			Description:                    "Product updated 1",
			Width:                          10,
			Height:                         20,
			Length:                         1,
			NetWeight:                      1,
			ExpirationRate:                 1,
			RecommendedFreezingTemperature: 1,
			FreezingRate:                   1,
			ProductTypeID:                  1,
			SellerID:                       1,
		}

		listOfProducts := map[int]model.Product{
			1: {
				ID:                             1,
				ProductCode:                    "P001",
				Description:                    "Product 1",
				Width:                          10,
				Height:                         20,
				Length:                         1,
				NetWeight:                      1,
				ExpirationRate:                 1,
				RecommendedFreezingTemperature: 1,
				FreezingRate:                   1,
				ProductTypeID:                  1,
				SellerID:                       1,
			},
		}

		srm.On("GetByID", 1).Return(model.Seller{ID: 1}, nil)
		prm.On("GetAll").Return(listOfProducts, nil)
		prm.On("GetByID", 1).Return(listOfProducts[1], nil)
		prm.On("Update", 1, mock.Anything).Return(inputProduct, nil)

		productUpdated, err := productService.UpdateProduct(1, inputProduct)

		assert.NoError(t, err)
		assert.Equal(t, inputProduct, productUpdated)
		prm.AssertExpectations(t)
		srm.AssertExpectations(t)
	})

	t.Run("Should return error for seller not found", func(t *testing.T) {
		productService := loadDependencies()
		prm := productService.ProductRepository.(*mocks.MockIProductsRepo)
		srm := productService.SellerRepository.(*mocks.MockISellerRepo)

		srm.On("GetByID", 1).Return(model.Seller{}, errors.New("seller not found"))

		productUpdated, err := productService.UpdateProduct(1, model.Product{SellerID: 1})

		assert.EqualError(t, err, "seller not found")
		assert.Equal(t, model.Product{}, productUpdated)
		prm.AssertExpectations(t)
		srm.AssertExpectations(t)
	})

	t.Run("Should return not found error for product", func(t *testing.T) {
		productService := loadDependencies()
		prm := productService.ProductRepository.(*mocks.MockIProductsRepo)
		srm := productService.SellerRepository.(*mocks.MockISellerRepo)

		srm.On("GetByID", 1).Return(model.Seller{ID: 1}, nil)
		prm.On("GetAll").Return(make(map[int]model.Product), nil)
		prm.On("GetByID", 2).Return(model.Product{}, customerror.HandleError("product", customerror.ErrorNotFound, ""))

		productUpdated, err := productService.UpdateProduct(2, model.Product{SellerID: 1})

		assert.Equal(t, customerror.HandleError("product", customerror.ErrorNotFound, ""), err)
		assert.Equal(t, model.Product{}, productUpdated)
		prm.AssertExpectations(t)
		srm.AssertExpectations(t)
	})

	t.Run("Should return conflict error, because cannot update product code if this code already exists", func(t *testing.T) {
		productService := loadDependencies()
		prm := productService.ProductRepository.(*mocks.MockIProductsRepo)
		srm := productService.SellerRepository.(*mocks.MockISellerRepo)

		srm.On("GetByID", 1).Return(model.Seller{ID: 1}, nil)

		listOfProducts := map[int]model.Product{
			1: {
				ID:                             1,
				ProductCode:                    "P001",
				Description:                    "Product 1",
				Width:                          10,
				Height:                         20,
				Length:                         1,
				NetWeight:                      1,
				ExpirationRate:                 1,
				RecommendedFreezingTemperature: 1,
				FreezingRate:                   1,
				ProductTypeID:                  1,
				SellerID:                       1,
			},
		}

		prm.On("GetAll").Return(listOfProducts, nil)

		productUpdated, err := productService.UpdateProduct(2, model.Product{
			ID:                             1,
			ProductCode:                    "P001",
			Description:                    "Product updated 1",
			Width:                          10,
			Height:                         20,
			Length:                         1,
			NetWeight:                      1,
			ExpirationRate:                 1,
			RecommendedFreezingTemperature: 1,
			FreezingRate:                   1,
			ProductTypeID:                  1,
			SellerID:                       1,
		})

		assert.Equal(t, customerror.CustomError{Object: "P001", Err: customerror.ErrConflict}, err)
		assert.Equal(t, model.Product{}, productUpdated)
		prm.AssertExpectations(t)
		srm.AssertExpectations(t)
	})
}

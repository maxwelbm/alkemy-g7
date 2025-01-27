package service

import (
	"errors"
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type productsRepositoryMock struct {
	mock.Mock
}

func (p *productsRepositoryMock) Create(product model.Product) (model.Product, error) {
	args := p.Called(product)
	return args.Get(0).(model.Product), args.Error(1)
}

func (p *productsRepositoryMock) Delete(id int) error {
	args := p.Called(id)
	return args.Error(0)
}

func (p *productsRepositoryMock) GetAll() (map[int]model.Product, error) {
	args := p.Called()
	return args.Get(0).(map[int]model.Product), args.Error(1)
}

func (p *productsRepositoryMock) GetByID(id int) (model.Product, error) {
	args := p.Called(id)
	return args.Get(0).(model.Product), args.Error(1)
}

func (p *productsRepositoryMock) Update(id int, product model.Product) (model.Product, error) {
	args := p.Called(id, product)
	return args.Get(0).(model.Product), args.Error(1)
}

type sellerRepositoryMock struct {
	mock.Mock
}

func (s *sellerRepositoryMock) Get() ([]model.Seller, error) {
	panic("1")
}
func (s *sellerRepositoryMock) GetById(id int) (model.Seller, error) {
	args := s.Called(id)
	return args.Get(0).(model.Seller), args.Error(1)
}
func (s *sellerRepositoryMock) Post(seller *model.Seller) (model.Seller, error) {
	panic("3")
}
func (s *sellerRepositoryMock) Patch(id int, seller *model.Seller) (model.Seller, error) {
	panic("4")
}
func (s *sellerRepositoryMock) Delete(id int) error {
	panic("5")
}

func loadDependencies() *ProductService {
	productRepoMock := new(productsRepositoryMock)
	sellerRepositoryMock := new(sellerRepositoryMock)
	productServiceMock := NewProductService(productRepoMock, sellerRepositoryMock)
	return productServiceMock
}

func TestGetAllProducts(t *testing.T) {
	productService := loadDependencies()
	t.Run("should return the list of products", func(t *testing.T) {
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

		mockRepo := productService.ProductRepository.(*productsRepositoryMock)

		mockRepo.On("GetAll").Return(data, nil)

		productList, err := productService.GetAllProducts()

		assert.NoError(t, err)
		assert.Equal(t, expectedValue, productList)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	productService := loadDependencies()

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

	testCases := []struct {
		name            string
		id              int
		expectedProduct model.Product
		expectedError   error
	}{
		{
			name:            "should return the product",
			id:              1,
			expectedProduct: expectedProduct,
			expectedError:   nil,
		},
		{
			name:            "should return not found error",
			id:              2,
			expectedProduct: model.Product{},
			expectedError:   custom_error.HandleError("product", custom_error.ErrorNotFound, ""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := productService.ProductRepository.(*productsRepositoryMock)

			if tc.expectedError != nil {
				mockRepo.On("GetById", tc.id).Return(model.Product{}, tc.expectedError)
			} else {
				mockRepo.On("GetById", tc.id).Return(expectedProduct, nil)
			}

			product, err := productService.ProductRepository.GetByID(tc.id)

			if tc.expectedError != nil {
				assert.ErrorIs(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedProduct, product)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteById(t *testing.T) {
	productService := loadDependencies()

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

	testCases := []struct {
		name          string
		id            int
		expectedError error
	}{
		{
			name:          "should return the product",
			id:            1,
			expectedError: nil,
		},
		{
			name:          "should return not found error",
			id:            2,
			expectedError: custom_error.HandleError("product", custom_error.ErrorNotFound, ""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := productService.ProductRepository.(*productsRepositoryMock)

			if tc.expectedError != nil {
				mockRepo.On("GetById", tc.id).Return(model.Product{}, tc.expectedError)
			} else {
				mockRepo.On("GetById", tc.id).Return(data, nil)
				mockRepo.On("Delete", tc.id).Return(nil)
			}

			err := productService.DeleteProduct(tc.id)

			if err != nil {
				assert.Equal(t, err, tc.expectedError)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestCreateProduct(t *testing.T) {
	productService := loadDependencies()

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

	testCases := []struct {
		name                            string
		idTestCase                      int
		param                           model.Product
		expectedReturnProductMockSucess model.Product
		expectedReturnSellerMockSucess  model.Seller
		expectedReturnMockError         error
	}{
		{
			name:                            "Should return sucess to create product",
			idTestCase:                      1,
			param:                           dataProduct,
			expectedReturnProductMockSucess: dataProduct,
			expectedReturnSellerMockSucess:  dataSeller,
			expectedReturnMockError:         nil,
		},
		{
			name:       "Should return validation error",
			idTestCase: 2,
			param: model.Product{
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
			},
			expectedReturnProductMockSucess: dataProduct,
			expectedReturnSellerMockSucess:  dataSeller,
			expectedReturnMockError:         errors.New("validation errors: ProductCode is required"),
		},
		{
			name:                            "Should return seller not found",
			idTestCase:                      3,
			param:                           dataProduct,
			expectedReturnProductMockSucess: dataProduct,
			expectedReturnSellerMockSucess:  dataSeller,
			expectedReturnMockError:         errors.New("seller not found"),
		},
		{
			name:                            "Should return exist by product code",
			idTestCase:                      4,
			param:                           listOfProducts[1],
			expectedReturnProductMockSucess: dataProduct,
			expectedReturnSellerMockSucess:  dataSeller,
			expectedReturnMockError:         custom_error.CustomError{Object: listOfProducts[1].ProductCode, Err: custom_error.ErrConflict},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			productRepoMock := productService.ProductRepository.(*productsRepositoryMock)
			sellerRepoMock := productService.SellerRepository.(*sellerRepositoryMock)

			if tc.expectedReturnMockError != nil {
				if tc.idTestCase == 3 {
					sellerRepoMock.On("GetById", tc.expectedReturnProductMockSucess.SellerID).
						Return(model.Seller{}, errors.New("seller not found"))
				}
				if tc.idTestCase == 4 {
					sellerRepoMock.On("GetById", tc.expectedReturnProductMockSucess.SellerID).
						Return(tc.expectedReturnSellerMockSucess, nil)

					productRepoMock.On("GetAll").Return(listOfProducts, nil)

				}

			} else {
				sellerRepoMock.On("GetById", tc.expectedReturnProductMockSucess.SellerID).
					Return(tc.expectedReturnSellerMockSucess, nil)
				productRepoMock.On("GetAll").Return(listOfProducts, nil)
				productRepoMock.On("Create", tc.param).Return(tc.expectedReturnProductMockSucess, nil)
			}

			product, err := productService.CreateProduct(tc.param)

			if err != nil {
				assert.Equal(t, err, tc.expectedReturnMockError)
			} else {
				assert.Equal(t, tc.expectedReturnProductMockSucess, product)
			}
			productRepoMock.AssertExpectations(t)
			sellerRepoMock.AssertExpectations(t)
		})
	}

}

func TestUpdateProducts(t *testing.T) {
	type testCase struct {
		name            string
		productID       int
		inputProduct    model.Product
		mockSetup       func(prm *productsRepositoryMock, srm *sellerRepositoryMock)
		expectedProduct model.Product
		expectedError   error
	}

	testCases := []testCase{
		{
			name:      "Should return success to update product",
			productID: 1,
			inputProduct: model.Product{
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
			},
			mockSetup: func(prm *productsRepositoryMock, srm *sellerRepositoryMock) {
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

				srm.On("GetById", 1).Return(model.Seller{ID: 1}, nil)
				prm.On("GetAll").Return(listOfProducts, nil)
				prm.On("GetById", 1).Return(listOfProducts[1], nil)
				prm.On("Update", 1, mock.Anything).Return(model.Product{
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
				}, nil)
			},
			expectedProduct: model.Product{
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
			},
			expectedError: nil,
		},
		{
			name:      "Should return error for seller not found",
			productID: 1,
			inputProduct: model.Product{
				SellerID: 1,
			},
			mockSetup: func(prm *productsRepositoryMock, srm *sellerRepositoryMock) {
				srm.On("GetById", 1).Return(model.Seller{}, errors.New("seller not found"))
			},
			expectedProduct: model.Product{},
			expectedError:   errors.New("seller not found"),
		},
		{
			name:      "Should return not found error for product",
			productID: 2,
			inputProduct: model.Product{
				SellerID: 1,
			},
			mockSetup: func(prm *productsRepositoryMock, srm *sellerRepositoryMock) {
				srm.On("GetById", 1).Return(model.Seller{ID: 1}, nil)
				prm.On("GetAll").Return(make(map[int]model.Product), nil)
				prm.On("GetById", 2).Return(model.Product{}, custom_error.HandleError("product", custom_error.ErrorNotFound, ""))
			},
			expectedProduct: model.Product{},
			expectedError:   custom_error.HandleError("product", custom_error.ErrorNotFound, ""),
		},
		{
			name:      "Should return conflict error, because cannot update product code if this code already exists",
			productID: 2,
			inputProduct: model.Product{
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
			},
			mockSetup: func(prm *productsRepositoryMock, srm *sellerRepositoryMock) {
				srm.On("GetById", 1).Return(model.Seller{ID: 1}, nil)
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
			},
			expectedProduct: model.Product{},
			expectedError:   custom_error.CustomError{Object: "P001", Err: custom_error.ErrConflict},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			productService := loadDependencies()
			prm := productService.ProductRepository.(*productsRepositoryMock)
			srm := productService.SellerRepository.(*sellerRepositoryMock)

			tc.mockSetup(prm, srm)

			productUpdated, err := productService.UpdateProduct(tc.productID, tc.inputProduct)

			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedProduct, productUpdated)
			}

			prm.AssertExpectations(t)
			srm.AssertExpectations(t)
		})
	}
}

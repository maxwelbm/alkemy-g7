package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type productsRepositoryMock struct {
	mock.Mock
}

func (p *productsRepositoryMock) Create(product model.Product) (model.Product, error) {
	args := p.Called()
	return args.Get(0).(model.Product), args.Error(1)
}

func (p *productsRepositoryMock) Delete(id int) error {
	args := p.Called()
	return args.Error(1)
}

func (p *productsRepositoryMock) GetAll() (map[int]model.Product, error) {
	args := p.Called()
	return args.Get(0).(map[int]model.Product), args.Error(1)
}

func (p *productsRepositoryMock) GetById(id int) (model.Product, error) {
	args := p.Called(id)
	return args.Get(0).(model.Product), args.Error(1)
}

// Update implements interfaces.IProductsRepo.
func (p *productsRepositoryMock) Update(id int, product model.Product) (model.Product, error) {
	panic("unimplemented")
}

type sellerRepositoryMock struct {
	mock.Mock
}

func (s *sellerRepositoryMock) Get() ([]model.Seller, error) {
	panic("1")
}
func (s *sellerRepositoryMock) GetById(id int) (model.Seller, error) {
	panic("2")
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

		expectedValue := []model.Product{model.Product{
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
		name             string
		id               int
		expectedProduct  model.Product
		expectedError    error
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

			product, err := productService.ProductRepository.GetById(tc.id)

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

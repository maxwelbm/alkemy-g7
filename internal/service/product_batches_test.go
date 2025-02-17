package service_test

import (
	"testing"
	"time"

	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
)

func setupProductBatches(t *testing.T) *service.ProductBatchesService {
	mockRepo := mocks.NewMockIProductBatchesRepo(t)
	pbService := service.CreateProductBatchesService(mockRepo, mocks.NewMockIProductService(t), mocks.NewMockISectionService(t))
	return pbService
}

func TestServiceProductBatches_Post(t *testing.T) {
	t.Run("given a valid product batch id then return it with no error", func(t *testing.T) {
		svc := setupProductBatches(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		createdPB := model.ProductBatches{
			ID:                 1,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            parsedTime,
			InitialQuantity:    5,
			ManufacturingDate:  parsedTime,
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}

		mockProductService := svc.SvcProd.(*mocks.MockIProductService)
		mockProductService.On("GetProductByID", createdPB.ProductID).Return(model.Product{
			ID:                             1,
			ProductCode:                    "P01",
			Description:                    "Product 1",
			Width:                          100.00,
			Height:                         50.00,
			Length:                         20.00,
			NetWeight:                      5.00,
			ExpirationRate:                 10.00,
			RecommendedFreezingTemperature: 20.00,
			FreezingRate:                   10.00,
			ProductTypeID:                  1,
			SellerID:                       1,
		}, nil)

		mockSectionService := svc.SvcSec.(*mocks.MockISectionService)
		mockSectionService.On("GetByID", createdPB.SectionID).Return(model.Section{
			ID:                 1,
			SectionNumber:      "S01",
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			CurrentCapacity:    5,
			MinimumCapacity:    1,
			MaximumCapacity:    10,
			WarehouseID:        1,
			ProductTypeID:      1,
		}, nil)

		mockRepo := svc.Rp.(*mocks.MockIProductBatchesRepo)
		mockRepo.On("Post", &createdPB).Return(createdPB, nil)

		pb, err := svc.Post(&createdPB)

		assert.NoError(t, err)
		assert.Equal(t, createdPB, pb)
	})

	t.Run("given an invalid product then return error", func(t *testing.T) {
		svc := setupProductBatches(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		createdPB := model.ProductBatches{
			ID:                 1,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            parsedTime,
			InitialQuantity:    5,
			ManufacturingDate:  parsedTime,
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}

		expectedError := customerror.HandleError("product", customerror.ErrorNotFound, "")

		mockProductService := svc.SvcProd.(*mocks.MockIProductService)
		mockProductService.On("GetProductByID", createdPB.ProductID).Return(model.Product{}, expectedError)

		pb, err := svc.Post(&createdPB)

		assert.ErrorIs(t, expectedError, err)
		assert.Equal(t, model.ProductBatches{}, pb)
		mockProductService.AssertExpectations(t)
	})

	t.Run("given an invalid section then return error", func(t *testing.T) {
		svc := setupProductBatches(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		createdPB := model.ProductBatches{
			ID:                 1,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            parsedTime,
			InitialQuantity:    5,
			ManufacturingDate:  parsedTime,
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}

		expectedError := customerror.HandleError("section", customerror.ErrorNotFound, "")

		mockProductService := svc.SvcProd.(*mocks.MockIProductService)
		mockProductService.On("GetProductByID", createdPB.ProductID).Return(model.Product{
			ID:                             1,
			ProductCode:                    "P01",
			Description:                    "Product 1",
			Width:                          100.00,
			Height:                         50.00,
			Length:                         20.00,
			NetWeight:                      5.00,
			ExpirationRate:                 10.00,
			RecommendedFreezingTemperature: 20.00,
			FreezingRate:                   10.00,
			ProductTypeID:                  1,
			SellerID:                       1,
		}, nil)

		mockSectionService := svc.SvcSec.(*mocks.MockISectionService)
		mockSectionService.On("GetByID", createdPB.SectionID).Return(model.Section{}, expectedError)

		pb, err := svc.Post(&createdPB)

		assert.ErrorIs(t, expectedError, err)
		assert.Equal(t, model.ProductBatches{}, pb)
	})

	t.Run("given an invalid product batch then return error", func(t *testing.T) {
		svc := setupProductBatches(t)

		createdPB := model.ProductBatches{}

		expectedError := customerror.HandleError("product batches", customerror.ErrorInvalid, "")

		pb, err := svc.Post(&createdPB)

		assert.Error(t, expectedError, err)
		assert.Equal(t, model.ProductBatches{}, pb)

	})
}

func TestServiceProductBatches_GetByID(t *testing.T) {
	t.Run("Get existing purchase order successfully", func(t *testing.T) {
		svc := setupProductBatches(t)

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)

		createdPB := model.ProductBatches{
			ID:                 1,
			BatchNumber:        "B01",
			CurrentQuantity:    10,
			CurrentTemperature: 10.00,
			MinimumTemperature: 5.00,
			DueDate:            parsedTime,
			InitialQuantity:    5,
			ManufacturingDate:  parsedTime,
			ManufacturingHour:  10,
			ProductID:          1,
			SectionID:          1,
		}

		mockRepo := svc.Rp.(*mocks.MockIProductBatchesRepo)
		mockRepo.On("GetByID", createdPB.ID).Return(createdPB, nil)

		pb, err := svc.GetByID(createdPB.ID)

		assert.NoError(t, err)
		assert.Equal(t, createdPB, pb)
		mockRepo.AssertExpectations(t)
	})
}

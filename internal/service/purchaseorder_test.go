package service_test

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customError"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func setupPurchaseOrderService() *service.PurchaseOrderService {
	mockRepo := new(repository.PurchaseOrderRepositoryMock)
	purchaseService := service.NewPurchaseOrderService(mockRepo, service.NewMockBuyerService(), service.NewProductrecMock())
	return purchaseService
}

func TestCreatePurchaseOrder(t *testing.T) {
	t.Run("Created Purchase Order successfully", func(t *testing.T) {
		Svc := setupPurchaseOrderService()
		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)
		createdOrder := model.PurchaseOrder{
			ID:              1,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         1,
			ProductRecordID: 1,
		}

		mockBuyerService := Svc.SvcBuyer.(*service.MockBuyerService)
		mockBuyerService.On("GetBuyerByID", createdOrder.BuyerID).Return(model.Buyer{
			ID:           1,
			CardNumberID: "CN001",
			FirstName:    "Jhon",
			LastName:     "Doe",
		}, nil)

		mockProductRec := Svc.SvcProductRec.(*service.ProductrecMock)
		mockProductRec.On("GetProductRecordByID", createdOrder.ProductRecordID).Return(model.ProductRecords{
			ID:             1,
			LastUpdateDate: parsedTime,
			PurchasePrice:  1,
			SalePrice:      1,
			ProductID:      1,
		}, nil)

		mockRepo := Svc.Rp.(*repository.PurchaseOrderRepositoryMock)
		mockRepo.On("Post", createdOrder).Return(int64(1), nil)
		mockRepo.On("GetByID", createdOrder.ID).Return(createdOrder, nil)

		purchaser, err := Svc.CreatePurchaseOrder(createdOrder)
		assert.NoError(t, err)
		assert.Equal(t, createdOrder, purchaser)
		mockRepo.AssertExpectations(t)
		mockBuyerService.AssertExpectations(t)

	})

	t.Run("Buyer not exists", func(t *testing.T) {
		Svc := setupPurchaseOrderService()

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)
		createdOrder := model.PurchaseOrder{
			ID:              1,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         99,
			ProductRecordID: 1,
		}
		expectedError := customError.NewBuyerError(http.StatusNotFound, customError.ErrNotFound.Error(), "Buyer")

		mockBuyerService := Svc.SvcBuyer.(*service.MockBuyerService)
		mockBuyerService.On("GetBuyerByID", createdOrder.BuyerID).Return(model.Buyer{}, expectedError)

		purchaser, err := Svc.CreatePurchaseOrder(createdOrder)

		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, purchaser, model.PurchaseOrder{})
		mockBuyerService.AssertExpectations(t)
	})

	t.Run("Product Record not exists", func(t *testing.T) {
		Svc := setupPurchaseOrderService()

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)
		createdOrder := model.PurchaseOrder{
			ID:              1,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         1,
			ProductRecordID: 99,
		}

		mockBuyerService := Svc.SvcBuyer.(*service.MockBuyerService)
		mockBuyerService.On("GetBuyerByID", createdOrder.BuyerID).Return(model.Buyer{
			ID:           1,
			CardNumberID: "CN001",
			FirstName:    "Jhon",
			LastName:     "Doe",
		}, nil)

		expectedError := customError.HandleError("product record", customError.ErrorNotFound, "")

		mockProductRec := Svc.SvcProductRec.(*service.ProductrecMock)
		mockProductRec.On("GetProductRecordByID", createdOrder.ProductRecordID).Return(model.ProductRecords{}, expectedError)

		purchaser, err := Svc.CreatePurchaseOrder(createdOrder)

		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, purchaser, model.PurchaseOrder{})
		mockProductRec.AssertExpectations(t)
		mockBuyerService.AssertExpectations(t)
	})

	t.Run("Order number already exists", func(t *testing.T) {
		Svc := setupPurchaseOrderService()
		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)
		createdOrder := model.PurchaseOrder{
			ID:              1,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         1,
			ProductRecordID: 1,
		}

		mockBuyerService := Svc.SvcBuyer.(*service.MockBuyerService)
		mockBuyerService.On("GetBuyerByID", createdOrder.BuyerID).Return(model.Buyer{
			ID:           1,
			CardNumberID: "CN001",
			FirstName:    "Jhon",
			LastName:     "Doe",
		}, nil)

		mockProductRec := Svc.SvcProductRec.(*service.ProductrecMock)
		mockProductRec.On("GetProductRecordByID", createdOrder.ProductRecordID).Return(model.ProductRecords{
			ID:             1,
			LastUpdateDate: parsedTime,
			PurchasePrice:  1,
			SalePrice:      1,
			ProductID:      1,
		}, nil)

		expectedError := customError.NewPurcahseOrderError(http.StatusConflict, customError.ErrConflict.Error(), "order_number")

		mockRepo := Svc.Rp.(*repository.PurchaseOrderRepositoryMock)
		mockRepo.On("Post", createdOrder).Return(int64(0), expectedError)

		purchaser, err := Svc.CreatePurchaseOrder(createdOrder)
		assert.ErrorIs(t, err, expectedError)
		assert.Equal(t, model.PurchaseOrder{}, purchaser)
		mockRepo.AssertExpectations(t)
		mockProductRec.AssertExpectations(t)
		mockBuyerService.AssertExpectations(t)
	})
}

func TestGetPurchaseOrderByID(t *testing.T) {
	t.Run("Get existing purchase order successfully", func(t *testing.T) {
		Svc := setupPurchaseOrderService()

		dateString := "2025-01-01T00:00:00Z"
		layout := time.RFC3339

		parsedTime, err := time.Parse(layout, dateString)
		assert.NoError(t, err)
		PurchaseSearch := model.PurchaseOrder{
			ID:              1,
			OrderNumber:     "ON001",
			OrderDate:       parsedTime,
			TrackingCode:    "TC001",
			BuyerID:         99,
			ProductRecordID: 1,
		}

		mockRepo := Svc.Rp.(*repository.PurchaseOrderRepositoryMock)
		mockRepo.On("GetByID", PurchaseSearch.ID).Return(PurchaseSearch, nil)

		purchaser, err := Svc.GetPurchaseOrderByID(PurchaseSearch.ID)

		assert.NoError(t, err)
		assert.Equal(t, PurchaseSearch, purchaser)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Purchase order not found", func(t *testing.T) {
		Svc := setupPurchaseOrderService()

		exepctedError := customError.NewPurcahseOrderError(http.StatusNotFound, customError.ErrNotFound.Error(), "Purchase Order")

		mockRepo := Svc.Rp.(*repository.PurchaseOrderRepositoryMock)
		mockRepo.On("GetByID", 99).Return(model.PurchaseOrder{}, exepctedError)

		purchaser, err := Svc.GetPurchaseOrderByID(99)

		assert.ErrorIs(t, err, exepctedError)
		assert.Equal(t, model.PurchaseOrder{}, purchaser)
		mockRepo.AssertExpectations(t)
	})
}

package service_test

import (
	"errors"
	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func setup(t *testing.T) *service.BuyerService {
	mockRepo := mocks.NewMockIBuyerRepo(t)

	return service.NewBuyerService(mockRepo)
}

func TestGetAllBuyer(t *testing.T) {
	t.Run("return a list of all existing buyers successfully", func(t *testing.T) {
		svc := setup(t)

		expectedBuyers := []model.Buyer{{ID: 1, FirstName: "John", LastName: "Doe", CardNumberID: "1234"},
			{ID: 2, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}}

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("Get").Return(expectedBuyers, nil)

		buyers, err := svc.GetAllBuyer()

		assert.Equal(t, expectedBuyers, buyers)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("return an error when fetching buyers", func(t *testing.T) {
		svc := setup(t)

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("Get").Return([]model.Buyer{}, errors.New("Unmapped error"))

		buyers, err := svc.GetAllBuyer()

		assert.Equal(t, []model.Buyer{}, buyers)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)

	})

}

func TestGetBuyerByID(t *testing.T) {
	t.Run("return buyer by id existing successfully", func(t *testing.T) {
		svc := setup(t)

		expectedBuyer := model.Buyer{ID: 2, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("GetByID", 2).Return(expectedBuyer, nil)

		buyer, err := svc.GetBuyerByID(2)

		assert.Equal(t, expectedBuyer, buyer)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Return buyer not Found", func(t *testing.T) {
		svc := setup(t)

		expectedError := customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer")

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("GetByID", 99).Return(model.Buyer{}, expectedError)

		buyer, err := svc.GetBuyerByID(99)

		assert.Equal(t, model.Buyer{}, buyer)
		assert.ErrorIs(t, err, expectedError)
		mockRepo.AssertExpectations(t)

	})

	t.Run("return an error when fetching buyer", func(t *testing.T) {
		svc := setup(t)

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("GetByID", 2).Return(model.Buyer{}, errors.New("unmapped error"))

		buyer, err := svc.GetBuyerByID(2)

		assert.Equal(t, model.Buyer{}, buyer)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreateBuyer(t *testing.T) {
	t.Run("Buyer created successfully", func(t *testing.T) {
		svc := setup(t)

		createdBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}
		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("Post", createdBuyer).Return(int64(1), nil)

		mockRepo.On("GetByID", 1).Return(createdBuyer, nil)

		buyer, err := svc.CreateBuyer(createdBuyer)

		assert.NoError(t, err)
		assert.Equal(t, createdBuyer, buyer)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Return error card_number already exists", func(t *testing.T) {
		svc := setup(t)

		expectedError := customerror.NewBuyerError(http.StatusConflict, customerror.ErrConflict.Error(), "card_number_id")
		createdBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("Post", createdBuyer).Return(int64(0), expectedError)

		buyer, err := svc.CreateBuyer(createdBuyer)

		assert.Equal(t, model.Buyer{}, buyer)
		assert.ErrorIs(t, err, expectedError)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)

	})

	t.Run("return an error when creating buyer", func(t *testing.T) {
		svc := setup(t)

		createdBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("Post", createdBuyer).Return(int64(0), errors.New("unmapped Error"))

		buyer, err := svc.CreateBuyer(createdBuyer)

		assert.Equal(t, model.Buyer{}, buyer)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Buyer Not Found", func(t *testing.T) {
		svc := setup(t)

		createdBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}
		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("Post", createdBuyer).Return(int64(1), nil)

		expectedError := customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer")
		mockRepo.On("GetByID", 1).Return(model.Buyer{}, expectedError)

		buyer, err := svc.CreateBuyer(createdBuyer)

		assert.ErrorIs(t, err, expectedError)
		assert.Error(t, err)
		assert.Equal(t, model.Buyer{}, buyer)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateBuyer(t *testing.T) {
	t.Run("buyer updated successfuly", func(t *testing.T) {
		svc := setup(t)

		updatedBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("GetByID", 1).Return(updatedBuyer, nil)

		mockRepo.On("Update", 1, updatedBuyer).Return(nil)

		buyer, err := svc.UpdateBuyer(1, updatedBuyer)

		assert.NoError(t, err)
		assert.Equal(t, updatedBuyer, buyer)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Buyer Not Found", func(t *testing.T) {
		svc := setup(t)

		UpdateBuyer := model.Buyer{ID: 99, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}
		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)

		expectedError := customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer")
		mockRepo.On("GetByID", 99).Return(model.Buyer{}, expectedError)

		buyer, err := svc.UpdateBuyer(99, UpdateBuyer)

		assert.ErrorIs(t, err, expectedError)
		assert.Error(t, err)
		assert.Equal(t, model.Buyer{}, buyer)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Return error card_number already exists", func(t *testing.T) {
		svc := setup(t)

		expectedError := customerror.NewBuyerError(http.StatusConflict, customerror.ErrConflict.Error(), "card_number_id")
		updatedBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("GetByID", 1).Return(updatedBuyer, nil)

		mockRepo.On("Update", 1, updatedBuyer).Return(expectedError)

		buyer, err := svc.UpdateBuyer(1, updatedBuyer)

		assert.Equal(t, model.Buyer{}, buyer)
		assert.ErrorIs(t, err, expectedError)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)

	})

	t.Run("return an error when updating buyer", func(t *testing.T) {
		svc := setup(t)

		updatedBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("GetByID", 1).Return(updatedBuyer, nil)

		mockRepo.On("Update", 1, updatedBuyer).Return(errors.New("unmapped error"))

		buyer, err := svc.UpdateBuyer(1, updatedBuyer)

		assert.Equal(t, model.Buyer{}, buyer)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)

	})

}

func TestDeleteBuyerByID(t *testing.T) {
	t.Run("Delete buyer successfuly", func(t *testing.T) {
		svc := setup(t)

		deletedBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}
		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("GetByID", 1).Return(deletedBuyer, nil)

		mockRepo.On("Delete", 1).Return(nil)

		err := svc.DeleteBuyerByID(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Buyer Not Found", func(t *testing.T) {
		svc := setup(t)

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)

		expectedError := customerror.NewBuyerError(http.StatusNotFound, customerror.ErrNotFound.Error(), "Buyer")
		mockRepo.On("GetByID", 99).Return(model.Buyer{}, expectedError)

		err := svc.DeleteBuyerByID(99)

		assert.ErrorIs(t, err, expectedError)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("There are dependencies with the buyer", func(t *testing.T) {
		svc := setup(t)

		deletedBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}
		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("GetByID", 1).Return(deletedBuyer, nil)

		expectedError := customerror.NewBuyerError(http.StatusConflict, customerror.ErrDependencies.Error(), "Buyer")
		mockRepo.On("Delete", 1).Return(expectedError)

		err := svc.DeleteBuyerByID(1)

		assert.ErrorIs(t, err, expectedError)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("return an error when deleting buyer", func(t *testing.T) {
		svc := setup(t)

		deletedBuyer := model.Buyer{ID: 1, FirstName: "Ac", LastName: "Milan", CardNumberID: "4321"}
		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("GetByID", 1).Return(deletedBuyer, nil)

		mockRepo.On("Delete", 1).Return(errors.New("unmapped Error"))

		err := svc.DeleteBuyerByID(1)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestCountPurchaseOrderBuyer(t *testing.T) {
	t.Run("Count purchase by buyers successfuly", func(t *testing.T) {
		svc := setup(t)

		countPurchases := []model.BuyerPurchaseOrder{model.BuyerPurchaseOrder{
			ID:                  1,
			CardNumberID:        "cn001",
			FirstName:           "Jhon",
			LastName:            "DOe",
			PurchaseOrdersCount: 12,
		}}

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("CountPurchaseOrderBuyers").Return(countPurchases, nil)

		count, err := svc.CountPurchaseOrderBuyer()

		assert.NoError(t, err)
		assert.Equal(t, count, countPurchases)
		mockRepo.AssertExpectations(t)
	})
}

func TestCountPurchaseOrderByBuyerID(t *testing.T) {
	t.Run("Count purchase by buyer successfuly", func(t *testing.T) {
		svc := setup(t)

		countPurchase := model.BuyerPurchaseOrder{
			ID:                  1,
			CardNumberID:        "cn001",
			FirstName:           "Jhon",
			LastName:            "DOe",
			PurchaseOrdersCount: 12,
		}

		mockRepo := svc.Rp.(*mocks.MockIBuyerRepo)
		mockRepo.On("CountPurchaseOrderByBuyerID", countPurchase.ID).Return(countPurchase, nil)

		count, err := svc.CountPurchaseOrderByBuyerID(countPurchase.ID)

		assert.NoError(t, err)
		assert.Equal(t, count, countPurchase)
		mockRepo.AssertExpectations(t)
	})
}

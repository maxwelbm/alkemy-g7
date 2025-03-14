package service_test

import (
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/mocks"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/stretchr/testify/assert"
)

func setupWarehouse(t *testing.T) *service.WareHouseDefault {
	mockRepo := mocks.NewMockIWarehouseRepo(t)

	return service.NewWareHouseService(mockRepo, logMock)
}

func TestGetAllWarehouse(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {

		svc := setupWarehouse(t)

		expectWarehouse := []model.WareHouse{
			{
				ID:                 1,
				WareHouseCode:      "test",
				Telephone:          "test",
				MinimunCapacity:    1,
				MinimunTemperature: 1,
				Address:            "test",
			},
			{
				ID:                 2,
				WareHouseCode:      "test",
				Telephone:          "test",
				MinimunCapacity:    1,
				MinimunTemperature: 1,
				Address:            "test",
			},
		}

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)
		mockRepo.On("GetAllWareHouse").Return(expectWarehouse, nil)

		w, err := svc.GetAllWareHouse()

		assert.Equal(t, expectWarehouse, w)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllError", func(t *testing.T) {
		svc := setupWarehouse(t)

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)
		mockRepo.On("GetAllWareHouse").Return([]model.WareHouse{}, assert.AnError)

		w, err := svc.GetAllWareHouse()

		assert.Empty(t, w)
		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetByIdWareHouse(t *testing.T) {
	t.Run("GetById", func(t *testing.T) {
		svc := setupWarehouse(t)

		expectWarehouse := model.WareHouse{
			ID:                 1,
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)
		mockRepo.On("GetByIDWareHouse", 1).Return(expectWarehouse, nil)

		w, err := svc.GetByIDWareHouse(1)

		assert.Equal(t, expectWarehouse, w)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetByIdError", func(t *testing.T) {
		svc := setupWarehouse(t)

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)
		mockRepo.On("GetByIDWareHouse", 1).Return(model.WareHouse{}, assert.AnError)

		w, err := svc.GetByIDWareHouse(1)

		assert.Empty(t, w)
		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteByIdWareHouse(t *testing.T) {
	t.Run("DeleteById", func(t *testing.T) {
		svc := setupWarehouse(t)

		expectedWarehouse := model.WareHouse{
			ID:                 1,
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)
		mockRepo.On("GetByIDWareHouse", 1).Return(expectedWarehouse, nil)
		mockRepo.On("DeleteByIDWareHouse", 1).Return(nil)

		err := svc.DeleteByIDWareHouse(1)

		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteByIdError", func(t *testing.T) {
		svc := setupWarehouse(t)

		expectedWarehouse := model.WareHouse{
			ID:                 2,
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)
		mockRepo.On("GetByIDWareHouse", 1).Return(expectedWarehouse, nil)
		mockRepo.On("DeleteByIDWareHouse", 1).Return(assert.AnError)

		err := svc.DeleteByIDWareHouse(1)

		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteByIdNotFound", func(t *testing.T) {
		svc := setupWarehouse(t)

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)
		mockRepo.On("GetByIDWareHouse", 1).Return(model.WareHouse{}, assert.AnError)

		err := svc.DeleteByIDWareHouse(1)

		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestPostWareHouse(t *testing.T) {
	t.Run("Post", func(t *testing.T) {
		svc := setupWarehouse(t)

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)

		warehouse := model.WareHouse{
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}
		mockRepo.On("PostWareHouse", warehouse).Return(int64(1), nil)

		mockRepo.On("GetByIDWareHouse", 1).Return(warehouse, nil)

		w, err := svc.PostWareHouse(warehouse)

		assert.Nil(t, err)
		assert.Equal(t, warehouse, w)
		mockRepo.AssertExpectations(t)
	})

	t.Run("PostError", func(t *testing.T) {
		svc := setupWarehouse(t)

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)

		warehouse := model.WareHouse{
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}

		mockRepo.On("PostWareHouse", warehouse).Return(int64(0), assert.AnError)

		w, err := svc.PostWareHouse(warehouse)

		assert.NotNil(t, err)
		assert.Equal(t, model.WareHouse{}, w)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateWarehouse(t *testing.T) {
	t.Run("Update", func(t *testing.T) {
		svc := setupWarehouse(t)

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)

		warehouse := model.WareHouse{
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}
		mockRepo.On("GetByIDWareHouse", 1).Return(warehouse, nil)
		mockRepo.On("UpdateWareHouse", 1, warehouse).Return(nil)

		w, err := svc.UpdateWareHouse(1, warehouse)

		assert.Nil(t, err)
		assert.Equal(t, warehouse, w)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateError", func(t *testing.T) {
		svc := setupWarehouse(t)

		mockRepo := svc.Rp.(*mocks.MockIWarehouseRepo)

		warehouse := model.WareHouse{
			WareHouseCode:      "test",
			Telephone:          "test",
			MinimunCapacity:    1,
			MinimunTemperature: 1,
			Address:            "test",
		}
		mockRepo.On("GetByIDWareHouse", 3).Return(model.WareHouse{}, nil)
		mockRepo.On("UpdateWareHouse", 3, warehouse).Return(assert.AnError)

		w, err := svc.UpdateWareHouse(3, warehouse)

		assert.NotNil(t, err)
		assert.Equal(t, model.WareHouse{}, w)
		mockRepo.AssertExpectations(t)
	})
}

package service_test

import (
	"testing"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/stretchr/testify/assert"
)

func setup() *service.WareHouseDefault {
	mockRepo := new(repository.WareHouseMockRepo)

	return service.NewWareHoureService(mockRepo)
}

func TestGetAllWarehouse(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {

		svc := setup()

		expectWarehouse := []model.WareHouse{
			{
				Id:                 1,
				WareHouseCode:      "test",
				Telephone:          "test",
				MinimunCapacity:    1,
				MinimunTemperature: 1,
				Address:            "test",
			},
			{
				Id:                 2,
				WareHouseCode:      "test",
				Telephone:          "test",
				MinimunCapacity:    1,
				MinimunTemperature: 1,
				Address:            "test",
			},
		}
		mockRepo := svc.Rp.(*repository.WareHouseMockRepo)
		mockRepo.On("GetAllWareHouse").Return(expectWarehouse, nil)

		w, err := svc.GetAllWareHouse()

		assert.Equal(t, expectWarehouse, w)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetAllError", func(t *testing.T) {
		svc := setup()

		mockRepo := svc.Rp.(*repository.WareHouseMockRepo)
		mockRepo.On("GetAllWareHouse").Return([]model.WareHouse{}, assert.AnError)

		w, err := svc.GetAllWareHouse()

		assert.Empty(t, w)
		assert.NotNil(t, err)
		mockRepo.AssertExpectations(t)
	})
}

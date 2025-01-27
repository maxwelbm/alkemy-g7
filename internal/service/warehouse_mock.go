package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type WarehouseServiceMock struct {
	mock.Mock
}


func (mr *WarehouseServiceMock) DeleteByIDWareHouse(id int) error {
	args := mr.Called(id)

	return args.Error(0)
}

func (mr *WarehouseServiceMock) GetByIDWareHouse(id int) (w model.WareHouse, err error) {
	args := mr.Called(id)

	w = args.Get(0).(model.WareHouse)
	err = args.Error(1)

	return
}

func (mr *WarehouseServiceMock) PostWareHouse(warehouse model.WareHouse) (w model.WareHouse, err error) {
	args := mr.Called(warehouse)


	w = args.Get(0).(model.WareHouse)
	err = args.Error(1)

	return
}

func (mr *WarehouseServiceMock) UpdateWareHouse(id int, warehouse model.WareHouse) (w model.WareHouse, err error) {
	args := mr.Called(id, warehouse)

	w = args.Get(0).(model.WareHouse)
	err = args.Error(1)

	return
}

func (mr *WarehouseServiceMock) GetAllWareHouse() (w []model.WareHouse, err error) {
	args := mr.Called()

	w = args.Get(0).([]model.WareHouse)
	err = args.Error(1)

	return
}

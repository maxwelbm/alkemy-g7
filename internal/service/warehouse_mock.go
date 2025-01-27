package service

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/stretchr/testify/mock"
)

type WarehouseServiceMock struct {
	mock.Mock
}

func (mock *WarehouseServiceMock) DeleteByIdWareHouse(id int) error {
	args := mock.Called(id)

	return args.Error(0)
}

func (mock *WarehouseServiceMock) GetByIdWareHouse(id int) (w model.WareHouse, err error) {
	args := mock.Called(id)

	w = args.Get(0).(model.WareHouse)
	err = args.Error(1)

	return
}

func (mock *WarehouseServiceMock) PostWareHouse(warehouse model.WareHouse) (w model.WareHouse, err error) {
	args := mock.Called(warehouse)

	w = args.Get(0).(model.WareHouse)
	err = args.Error(1)

	return
}

func (mock *WarehouseServiceMock) UpdateWareHouse(id int, warehouse model.WareHouse) (w model.WareHouse, err error) {
	panic("unimplemented")
}

func (mock *WarehouseServiceMock) GetAllWareHouse() (w []model.WareHouse, err error) {
	args := mock.Called()

	w = args.Get(0).([]model.WareHouse)
	err = args.Error(1)

	return
}

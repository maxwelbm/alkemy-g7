package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IWarehouseService interface {
	GetAllWareHouse() (w map[int]model.WareHouse, err error)
	GetByIdWareHouse(id int) (w model.WareHouse, err error)
	PostWareHouse(warehouse model.WareHouse) (w model.WareHouse, err error)
	UpdateWareHouse(id int, warehouse model.WareHouse) (w model.WareHouse, err error)
	DeleteByIdWareHouse(id int) error
}

package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IWarehouseRepo interface {
	GetAllWareHouse() (w []model.WareHouse, err error)
	GetByIDWareHouse(id int) (w model.WareHouse, err error)
	PostWareHouse(warehouse model.WareHouse) (id int64, err error)
	UpdateWareHouse(id int, warehouse model.WareHouse) (err error)
	DeleteByIDWareHouse(id int) error
}

package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IWarehouseRepo interface {
	GetAllWareHouse() (w map[int]model.WareHouse, err error)
	GetByIdWareHouse(id int) (w model.WareHouse, err error)
	Post(warehouse model.WareHouse) (model.WareHouse, error)
	Update(id int, warehouse model.WareHouse) (model.WareHouse, error)
	Delete(id int) error
}

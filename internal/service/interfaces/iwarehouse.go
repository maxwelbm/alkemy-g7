package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IWarehouseService interface {
	GetAllWareHouse() (w map[int]model.WareHouse, err error)
	GetByIdWareHouse(id int) (w model.WareHouse, err error)
	Post(warehouse model.WareHouse) (w model.WareHouse, err error)
	Update(id int, warehouse model.WareHouse) (w model.WareHouse, err error)
	Delete(id int) error
}

package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IWarehouseRepo interface {
	Get() (map[int]model.WareHouse, error)
	GetById(id int) (model.WareHouse, error)
	Post(warehouse model.WareHouse) (model.WareHouse, error)
	Update(id int, warehouse model.WareHouse) (model.WareHouse, error)
	Delete(id int) error
}

package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IWarehouseService interface {
	Get() (map[int]model.WareHouseJson, error)
	GetById(id int) (model.WareHouseJson, error)
	Post(warehouse model.WareHouse) (model.WareHouse, error)
	Update(id int, warehouse model.WareHouse) (model.WareHouse, error)
	Delete(id int) error
}

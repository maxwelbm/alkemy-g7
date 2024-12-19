package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IBuyerRepo interface {
	Get() (map[int]model.Buyer, error)
	GetById(id int) (model.Buyer, error)
	Post(newBuyer model.Buyer) (model.Buyer, error)
	Update(id int, newBuyer model.Buyer) (model.Buyer, error)
	Delete(id int) error
}

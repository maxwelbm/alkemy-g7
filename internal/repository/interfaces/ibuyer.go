package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IBuyerRepo interface {
	Get() (buyers []model.Buyer, err error)
	GetById(id int) (buyer model.Buyer, err error)
	Post(newBuyer model.Buyer) (model.Buyer, error)
	Update(id int, newBuyer model.Buyer) error
	Delete(id int) (err error)
}

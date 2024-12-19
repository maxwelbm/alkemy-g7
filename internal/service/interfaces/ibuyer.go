package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IBuyerservice interface {
	GetAllBuyer() (map[int]model.Buyer, error)
	GetBuyerByID(id int) (model.Buyer, error)
	DeleteBuyerByID(id int) error
	CreateBuyer(newBuyer model.Buyer) (model.Buyer, error)
	UpdateBuyer(id int, newBuyer model.Buyer) (model.Buyer, error)
}

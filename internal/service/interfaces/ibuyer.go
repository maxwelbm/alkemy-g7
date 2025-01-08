package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type IBuyerservice interface {
	GetAllBuyer() (buyers []model.Buyer, err error)
	GetBuyerByID(id int) (buyer model.Buyer, err error)
	DeleteBuyerByID(id int) (err error)
	CreateBuyer(newBuyer model.Buyer) (model.Buyer, error)
	UpdateBuyer(id int, newBuyer model.Buyer) (model.Buyer, error)
}

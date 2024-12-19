package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type ISellerService interface {
	GetAll() (sellers []model.Seller, err error)
	GetById(id int) (sl model.Seller, err error)
	CreateSeller(seller model.Seller) (sl model.Seller, err error)
	UpdateSeller(id int, seller model.SellerUpdate) (sl model.Seller, err error)
	DeleteSeller(id int) error
}

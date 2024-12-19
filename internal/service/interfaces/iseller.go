package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type ISellerService interface {
	GetAll() (sellers []model.Seller, err error)
	GetById(id int) (sl model.Seller, err error)
	Post(seller model.Seller) (sl model.Seller, err error)
	Update(id int, seller model.Seller) (sl model.Seller, err error)
	Delete(id int) error
}

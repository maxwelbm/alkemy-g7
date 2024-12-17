package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type ISellerService interface {
	GetAll() (map[int]model.Seller, error)
	GetById(id int) (model.Seller, error)
	Post(seller model.Seller) (model.Seller, error)
	Update(id int, seller model.Seller) (model.Seller, error)
	Delete(id int) error
}

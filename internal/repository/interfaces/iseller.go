package interfaces

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type ISellerRepo interface {
	Get() ([]model.Seller, error)
	GetById(id int) (model.Seller, error)
	Post(seller *model.Seller) (model.Seller, error)
	Patch(id int, seller *model.Seller) (model.Seller, error)
	Delete(id int) error
}

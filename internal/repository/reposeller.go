package repository

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

func CreateRepositorySellers(db []model.Seller) *SellersRepository {
	defaultDb := make([]model.Seller, 0)
	if db != nil {
		defaultDb = db
	}
	return &SellersRepository{db: defaultDb}
}

type SellersRepository struct {
	db []model.Seller
}

func (rp *SellersRepository) Get() (map[int]model.Seller, error) {
	var sellers = make(map[int]model.Seller)

	for _, seller := range rp.db {
		sellers[seller.ID] = seller
	}

	return sellers, nil
}

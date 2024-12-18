package repository

import (
	"errors"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

func CreateRepositorySellers(db map[int]model.Seller) *SellersRepository {
	defaultDb := make(map[int]model.Seller, 0)
	if db != nil {
		defaultDb = db
	}
	return &SellersRepository{db: defaultDb}
}

type SellersRepository struct {
	db map[int]model.Seller
}

func (rp *SellersRepository) Get() (map[int]model.Seller, error) {
	var sellers = make(map[int]model.Seller)

	for _, seller := range rp.db {
		sellers[seller.ID] = seller
	}

	return sellers, nil
}

func (rp *SellersRepository) GetByID(id int) (seller model.Seller, err error) {
	for _, value := range rp.db {
		if value.ID == id {
			return value, nil
		}
	}

	err = errors.New("Any seller with this ID not found")

	return seller, err
}

func (rp *SellersRepository) Post(seller model.Seller) (sl model.Seller, err error) {
	id := 0
	for _, value := range rp.db {
		if value.ID > id {
			id = value.ID
		}
	}

	for _, value := range rp.db {
		if value.CID == seller.CID {
			err := errors.New("Seller's CID already exist")
			return sl, err
		}
	}

	seller.ID = id + 1
	id = seller.ID
	rp.db[id] = seller

	return seller, nil
}

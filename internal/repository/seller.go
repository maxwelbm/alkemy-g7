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

func (rp *SellersRepository) validateCID(sellers map[int]model.Seller, cid int) error {
	for _, s := range sellers {
		if s.CID == cid {
			err := errors.New("Seller's CID already exist")
			return err
			}
		}
	return nil
}

func (rp *SellersRepository) Get() (sellers []model.Seller, err error) {
	sellers = make([]model.Seller, 0)

	for _, s := range rp.db {
		sellers = append(sellers, s)
	}

	return sellers, nil
}

func (rp *SellersRepository) GetById(id int) (sl model.Seller, err error) {
	sl, exist := rp.db[id]
	if !exist {
		err = errors.New("Seller not found")
		return sl, err
	}
	return sl, err
}

func (rp *SellersRepository) Post(seller model.Seller) (sl model.Seller, err error) {
	id := 0
	for _, value := range rp.db {
		if value.ID > id {
			id = value.ID
		}
	}

	if err := rp.validateCID(rp.db, seller.CID); err != nil {
		return sl, err
	}

	seller.ID = id + 1
	id = seller.ID
	rp.db[id] = seller

	return seller, nil
}

func (rp *SellersRepository) Patch(id int, seller model.SellerUpdate) (sl model.Seller, err error) {
	if seller.CID != nil {
		if rp.db[id].CID != *seller.CID {
			if err := rp.validateCID(rp.db, *seller.CID); err != nil {
				return sl, err
			}
		}
	}

	if seller.CID != nil {
		sel := rp.db[id]
		sel.CID = *seller.CID
		rp.db[id] = sel
	}

	if seller.CompanyName != nil {
		sel := rp.db[id]
		sel.CompanyName = *seller.CompanyName
		rp.db[id] = sel
	}

	if seller.Address != nil {
		sel := rp.db[id]
		sel.Address = *seller.Address
		rp.db[id] = sel
	}

	if seller.Telephone != nil {
		sel := rp.db[id]
		sel.Telephone = *seller.Telephone
		rp.db[id] = sel
	}

	sl = rp.db[id]

	return sl, nil
}

func (rp *SellersRepository) Delete(id int) error {
	delete(rp.db, id)
	return nil
}

package repository

import (
	"errors"
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

type BuyerRepository struct {
	dbBuyer database.Database
}

func (br BuyerRepository) Delete(id int) error {
	buyer, err := br.GetById(id)

	if err != nil && errors.Is(err.(*custom_error.CustomError).Err, custom_error.NotFound) {
		return &custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	}

	delete(br.dbBuyer.TbBuyer, buyer.Id)

	return nil
}

func (br *BuyerRepository) Get() (map[int]model.Buyer, error) {

	if len(br.dbBuyer.TbBuyer) == 0 {
		return nil, fmt.Errorf("no buyers found")
	}
	return br.dbBuyer.TbBuyer, nil

}

func (br *BuyerRepository) GetById(id int) (model.Buyer, error) {
	buyer, ok := br.dbBuyer.TbBuyer[id]

	if !ok {
		return model.Buyer{}, &custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	}

	return buyer, nil
}

func (br BuyerRepository) Post(buyer model.Buyer) (model.Buyer, error) {
	panic("unimplemented")
}

func (br BuyerRepository) Update(id int, buyer model.Buyer) (model.Buyer, error) {
	panic("unimplemented")
}

func NewBuyerRepository(db database.Database) *BuyerRepository {
	return &BuyerRepository{dbBuyer: db}
}

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

	if err != nil {
		if errors.Is(err.(*custom_error.CustomError).Err, custom_error.NotFound) {
			return &custom_error.CustomError{Object: id, Err: custom_error.NotFound}
		}
		return err
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

	fmt.Println(ok)

	if !ok {
		return model.Buyer{}, &custom_error.CustomError{Object: id, Err: custom_error.NotFound}
	}

	return buyer, nil
}

func (br *BuyerRepository) Post(newBuyer model.Buyer) (model.Buyer, error) {
	BuyerExists := isCardNumberIdExists(newBuyer.CardNumberId, br)
	lastId := getLastIdBuyer(br.dbBuyer.TbBuyer)

	if BuyerExists {
		return model.Buyer{}, &custom_error.CustomError{Object: newBuyer, Err: custom_error.Conflict}
	}
	buyer := model.Buyer{
		Id:           lastId,
		CardNumberId: newBuyer.CardNumberId,
		FirstName:    newBuyer.FirstName,
		LastName:     newBuyer.LastName,
	}

	br.dbBuyer.TbBuyer[buyer.Id] = buyer

	return br.GetById(buyer.Id)

}

func (br BuyerRepository) Update(id int, newBuyer model.Buyer) (model.Buyer, error) {

	br.dbBuyer.TbBuyer[id] = newBuyer

	return br.GetById(id)
}

func isCardNumberIdExists(CardNumberId string, br *BuyerRepository) bool {

	for _, b := range br.dbBuyer.TbBuyer {
		if b.CardNumberId == CardNumberId {
			return true
		}
	}

	return false
}

func NewBuyerRepository(db database.Database) *BuyerRepository {
	return &BuyerRepository{dbBuyer: db}
}

func getLastIdBuyer(db map[int]model.Buyer) int {
	lastId := 0
	for _, buyer := range db {
		if buyer.Id > lastId {
			lastId = buyer.Id
		}
	}
	return lastId + 1
}

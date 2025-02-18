package service

import (
	"fmt"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository/interfaces"
	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"
)

type BuyerService struct {
	Rp  interfaces.IBuyerRepo
	log logger.Logger
}

func NewBuyerService(rp interfaces.IBuyerRepo, log logger.Logger) *BuyerService {
	return &BuyerService{Rp: rp, log: log}
}

func (bs *BuyerService) GetAllBuyer() (buyers []model.Buyer, err error) {
	bs.log.Log("BuyerService", "INFO", "initializing Get function")
	return bs.Rp.Get()
}

func (bs *BuyerService) GetBuyerByID(id int) (buyer model.Buyer, err error) {
	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("initializing GetByID function with parameter: %d", id))
	return bs.Rp.GetByID(id)
}

func (bs *BuyerService) DeleteBuyerByID(id int) (err error) {
	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("initializing DeleteBuyerID function with parameter: %d", id))
	_, err = bs.GetBuyerByID(id)

	if err != nil {
		bs.log.Log("BuyerService", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}
	bs.log.Log("BuyerService", "INFO", "Return in repository successful")
	return bs.Rp.Delete(id)
}

func (bs *BuyerService) CreateBuyer(newBuyer model.Buyer) (buyer model.Buyer, err error) {
	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("initializing CreateBuyer function with parameter: %v", newBuyer))
	id, err := bs.Rp.Post(newBuyer)

	if err != nil {
		bs.log.Log("BuyerService", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}
	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("Searching for Buyer created with ID: %v", id))
	buyer, err = bs.GetBuyerByID(int(id))

	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("Create Buyer Successfull: %v", buyer))
	return
}

func (bs *BuyerService) UpdateBuyer(id int, newBuyer model.Buyer) (buyer model.Buyer, err error) {
	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("initializing UpdateBuyer function with parameter: %v And ID: %d", newBuyer, id))
	existingBuyer, err := bs.GetBuyerByID(id)

	if err != nil {
		bs.log.Log("BuyerService", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	if newBuyer.CardNumberID != "" {
		existingBuyer.CardNumberID = newBuyer.CardNumberID
	}

	if newBuyer.FirstName != "" {
		existingBuyer.FirstName = newBuyer.FirstName
	}

	if newBuyer.LastName != "" {
		existingBuyer.LastName = newBuyer.LastName
	}

	err = bs.Rp.Update(existingBuyer.ID, existingBuyer)

	if err != nil {
		bs.log.Log("BuyerService", "ERROR", fmt.Sprintf("Error: %v", err))
		return
	}

	buyer, err = bs.GetBuyerByID(id)
	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("return buyer updated: %v", buyer))
	return
}

func (bs *BuyerService) CountPurchaseOrderByBuyerID(id int) (countBuyerPurchaseOrder model.BuyerPurchaseOrder, err error) {
	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("initializing CountPurchaseOrderByBuyerID function with parameter ID: %d", id))
	countBuyerPurchaseOrder, err = bs.Rp.CountPurchaseOrderByBuyerID(id)
	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("return CountPurchaseOrderByBuyerIDeBuyer: %d sucessfull", id))
	return
}

func (bs *BuyerService) CountPurchaseOrderBuyer() (countBuyerPurchaseOrder []model.BuyerPurchaseOrder, err error) {
	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("initializing CountPurchaseOrderBuyer function"))
	countBuyerPurchaseOrder, err = bs.Rp.CountPurchaseOrderBuyers()
	bs.log.Log("BuyerService", "INFO", fmt.Sprintf("return CountPurchaseOrderBuyer: %v sucessfull", countBuyerPurchaseOrder))
	return
}

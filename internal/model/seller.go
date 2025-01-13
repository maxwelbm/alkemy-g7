package model

import (
	"errors"
)

type Seller struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

var (
	ErrorSellerNotFound    error = errors.New("Seller not found")
	ErrorCIDAlreadyExist   error = errors.New("Seller's CID already exist")
	ErrorMissingID         error = errors.New("Missing int ID")
	ErrorInvalidJSONFormat error = errors.New("Invalid JSON request format attribute")
	ErrorNullAttribute     error = errors.New("Invalid body, empty value received.")
)

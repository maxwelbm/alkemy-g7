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
	Locality    int    `json:"locality_id"`
}

var (
	ErrorSellerNotFound    error = errors.New("Seller not found in the database")
	ErrorCIDSellerAlreadyExist   error = errors.New("Seller's CID already exists")
	ErrorMissingSellerID         error = errors.New("Missing 'id' parameter in the request")
	ErrorInvalidSellerJSONFormat error = errors.New("Invalid JSON format in the request body")
	ErrorNullSellerAttribute     error = errors.New("Invalid request body: received empty or null value")
)

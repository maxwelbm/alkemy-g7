package model

import "errors"

type Seller struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type SellerUpdate struct {
	CID         *int    `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *string `json:"telephone"`
}

type ResponseBodySeller struct {
	Data any `json:"data"`
}

type ResponseBodyErrorSeller struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var (
	ErrorSellerNotFound error = errors.New("Seller not found")
	ErrorCIDAlreadyExist error = errors.New("Seller's CID already exist")
	ErrorMissingID error = errors.New("Missing int ID")
	ErrorInvalidJSON error = errors.New("Invalid JSON format")
	ErrorAttribute error = errors.New("Invalid attribute, empty value received.")
	StatusNotFound string = "Not Found"
	StatusBadRequest string = "Bad Request"
	StatusConflict string = "Conflict"
	StatusUnprocessableEntity string = "Unprocessable Entity"
)
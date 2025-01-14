package model

import (
	"errors"
)

type Locality struct {
	ID       string `json:"id"`
	Locality string `json:"locality_name"`
	Province string `json:"province_name"`
	Country  string `json:"country_name"`
}

type LocalitiesJSONSellers struct {
	ID       string `json:"locality_id"`
	Locality string `json:"locality_name"`
	Sellers  string `json:"sellers_count"`
}

type LocalitiesJSONCarriers struct {
	ID       string `json:"locality_id"`
	Locality string `json:"locality_name"`
	Carriers  string `json:"carriers_count"`
}

var (
	ErrorLocalityNotFound          error = errors.New("Locality not found")
	ErrorIDAlreadyExist            error = errors.New("Locality ID already exist")
	ErrorMissingLocalityID         error = errors.New("Missing int ID")
	ErrorInvalidLocalityJSONFormat error = errors.New("Invalid JSON request format attribute")
	ErrorInvalidPathParam          error = errors.New("Invalid request path param")
	ErrorNullLocalityAttribute     error = errors.New("Invalid body, empty value received.")
)

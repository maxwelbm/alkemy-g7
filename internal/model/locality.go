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
	Carriers string `json:"carriers_count"`
}

var (
	ErrorLocalityNotFound          error = errors.New("Locality not found in the database")
	ErrorMissingLocalityID         error = errors.New("Missing 'id' parameter in the request")
	ErrorInvalidLocalityJSONFormat error = errors.New("Invalid JSON format in the request body")
	ErrorInvalidLocalityPathParam  error = errors.New("Invalid value for request path parameter")
	ErrorNullLocalityAttribute     error = errors.New("Invalid request body: received empty or null value")
)

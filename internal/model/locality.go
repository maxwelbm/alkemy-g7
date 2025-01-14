package model

import (
	er 	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
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
	Sellers  int `json:"sellers_count"`

type LocalitiesJSONCarriers struct {
	ID       string `json:"locality_id"`
	Locality string `json:"locality_name"`
	Carriers int `json:"carriers_count"`
}

func (s *Locality) ValidateEmptyFields(l *Locality) error {
	if l.Locality == "" || l.Province == "" || l.Country == "" {
		return er.ErrorNullLocalityAttribute
	}
	return nil
}

var (
	ErrorLocalityNotFound          error = errors.New("Locality not found in the database")
	ErrorIDAlreadyExist            error = errors.New("Locality ID already exists")
	ErrorMissingLocalityID         error = errors.New("Missing 'id' parameter in the request")
	ErrorInvalidLocalityJSONFormat error = errors.New("Invalid JSON format in the request body")
	ErrorInvalidPathParam          error = errors.New("Invalid value for request path parameter")
	ErrorNullLocalityAttribute     error = errors.New("Invalid request body: received empty or null value")	
)
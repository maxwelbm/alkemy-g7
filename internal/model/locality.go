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
	Sellers  string `json:"sellers_count"`
}

type LocalitiesJSONCarriers struct {
	ID       string `json:"locality_id"`
	Locality string `json:"locality_name"`
	Carriers string `json:"carriers_count"`
}

func (s *Locality) ValidateEmptyFields(l *Locality) error {
	if l.Locality == "" || l.Province == "" || l.Country == "" {
		return er.ErrorNullLocalityAttribute
	}
	return nil
}

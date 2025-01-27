package model

import (
	er "github.com/maxwelbm/alkemy-g7.git/pkg/customError"
)

type Locality struct {
	ID       int    `json:"id"`
	Locality string `json:"locality_name"`
	Province string `json:"province_name"`
	Country  string `json:"country_name"`
}

type LocalityJSON struct {
	ID       *int    `json:"id"`
	Locality *string `json:"locality_name"`
	Province *string `json:"province_name"`
	Country  *string `json:"country_name"`
}

type LocalitiesJSONSellers struct {
	ID       string `json:"locality_id"`
	Locality string `json:"locality_name"`
	Sellers  int    `json:"sellers_count"`
}

type LocalitiesJSONCarriers struct {
	ID       string `json:"locality_id"`
	Locality string `json:"locality_name"`
	Carriers int    `json:"carriers_count"`
}

func (s *Locality) ValidateEmptyFields(l *Locality) error {
	localityJSON := LocalityJSON{
		Locality: &l.Locality,
		Province: &l.Province,
		Country:  &l.Country,
	}

	if *localityJSON.Locality == "" || *localityJSON.Province == "" || *localityJSON.Country == "" {
		return er.ErrNullLocalityAttribute
	}

	if localityJSON.Locality == nil || localityJSON.Province == nil || localityJSON.Country == nil {
		return er.ErrInvalidLocalityJSONFormat
	}

	return nil
}

type LocalityResponseSwagger struct {
	Data []Locality `json:"data"`
}

type LocalitySellersResponseSwagger struct {
	Data []LocalitiesJSONSellers `json:"data"`
}

type LocalityCarriersResponseSwagger struct {
	Data []LocalitiesJSONCarriers `json:"data"`
}

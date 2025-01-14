package model

import (
	er 	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type Seller struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	Locality    int    `json:"locality_id"`
}

func (s *Seller) ValidateUpdateFields(sl *Seller, existSeller *Seller) {
	if sl.CID == 0 {
		sl.CID = existSeller.CID
	}
	if sl.Address == "" {
		sl.Address = existSeller.Address
	}
	if sl.CompanyName == "" {
		sl.CompanyName = existSeller.CompanyName
	}
	if sl.Telephone == "" {
		sl.Telephone = existSeller.Telephone
	}
	if sl.Locality == 0 {
		sl.Locality = existSeller.Locality
	}
}

func (s *Seller) ValidateEmptyFields(sl *Seller) error {
	if sl.CID == 0 || sl.Address == "" || sl.CompanyName == "" || sl.Telephone == "" || sl.Locality == 0 {
		return er.ErrorNullSellerAttribute
	}
	return nil
}

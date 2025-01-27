package model

import (
	er "github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type Seller struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	Locality    int    `json:"locality_id"`
}

type SellerJSON struct {
	ID          *int    `json:"id"`
	CID         *int    `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *string `json:"telephone"`
	Locality    *int    `json:"locality_id"`
}

func (s *Seller) ValidateUpdateFields(sl *Seller, existSeller *Seller) error {
	if sl.CID == 0 && sl.Address == "" && sl.CompanyName == "" && sl.Telephone == "" && sl.Locality == 0 {
		return er.ErrNullLocalityAttribute
	}

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

	return nil
}

func (s *Seller) ValidateEmptyFields(sl *Seller) error {
	sellerJSON := SellerJSON{
		CID:         &sl.CID,
		CompanyName: &sl.CompanyName,
		Address:     &sl.Address,
		Telephone:   &sl.Telephone,
		Locality:    &sl.Locality,
	}

	if *sellerJSON.CID == 0 || *sellerJSON.Address == "" || *sellerJSON.CompanyName == "" || *sellerJSON.Telephone == "" || *sellerJSON.Locality == 0 {
		return er.ErrNullSellerAttribute
	}

	if sellerJSON.CID == nil || sellerJSON.Address == nil || sellerJSON.CompanyName == nil || sellerJSON.Telephone == nil || sellerJSON.Locality == nil {
		return er.ErrInvalidSellerJSONFormat
	}

	return nil
}

type SellerResponseSwagger struct {
	Data []Seller `json:"data"`
}

package model

import (
	"fmt"
	"strings"
)

type Product struct {
	ID                             int     `json:"id"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"width"`
	Height                         float64 `json:"height"`
	Length                         float64 `json:"length"`
	NetWeight                      float64 `json:"net_weight"`
	ExpirationRate                 float64 `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductTypeID                  int     `json:"product_type_id"`
	SellerID                       int     `json:"seller_id"`
}

func (p *Product) Validate() error {
	var errors []string

	productCode := p.ProductCode
	if productCode == "" {
		errors = append(errors, "ProductCode is required")
	}

	description := p.Description
	if description == "" {
		errors = append(errors, "Description cannot be empty")
	}

	width := p.Width
	if width <= 0 {
		errors = append(errors, "Width must be greater than zero")
	}

	height := p.Height
	if height <= 0 {
		errors = append(errors, "Height must be greater than zero")
	}

	length := p.Length
	if length <= 0 {
		errors = append(errors, "Length must be greater than zero")
	}

	netWeight := p.NetWeight
	if netWeight <= 0 {
		errors = append(errors, "NetWeight must be greater than zero")
	}

	expirationRate := p.ExpirationRate
	if expirationRate < 0 {
		errors = append(errors, "ExpirationRate cannot be negative")
	}

	productTypeID := p.ProductTypeID
	if productTypeID <= 0 {
		errors = append(errors, "ProductTypeID must be greater than zero")
	}

	if len(errors) > 0 {
		return fmt.Errorf("validation errors: %s", strings.Join(errors, "; "))
	}

	return nil
}

type ProductResponseSwagger struct {
	Data []Product `json:"data"`
}
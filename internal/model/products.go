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
	var erros []string
    
	if p.ProductCode == "" {
		erros = append(erros, "ProductCode is required")
	}
	if p.Description == "" {
		erros = append(erros, "Description não pode estar vazia")
	}
	if p.Width <= 0 {
		erros = append(erros, "Width deve ser maior que zero")
	}
	if p.Height <= 0 {
		erros = append(erros, "Height deve ser maior que zero")
	}
	if p.Length <= 0 {
		erros = append(erros, "Length deve ser maior que zero")
	}
	if p.NetWeight <= 0 {
		erros = append(erros, "NetWeight deve ser maior que zero")
	}
	if p.ExpirationRate < 0 {
		erros = append(erros, "ExpirationRate não pode ser negativo")
	}

	if p.ProductTypeID <= 0 {
		erros = append(erros, "ProductTypeID deve ser maior que zero")
	}

	if len(erros) > 0 {
		return fmt.Errorf("erros de validação: %s", strings.Join(erros, "; "))
	}

	return nil
}

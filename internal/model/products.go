package model

import (
	"fmt"
	"strings"
)

type Product struct {
	ID                             int     `json:"id"`
	ProductCode                    string  `json:"product_code"`
	Description                    string  `json:"description"`
	Width                          float64 `json:"largura"`
	Height                         float64 `json:"altura"`
	Length                         float64 `json:"comprimento"`
	NetWeight                      float64 `json:"peso_liquido"`
	ExpirationRate                 float64 `json:"expiration_rate"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature"`
	FreezingRate                   float64 `json:"freezing_rate"`
	ProductTypeID                  int     `json:"product_type_id"`
	SellerID                       int     `json:"seller_id"`
}

func (p *Product) Validate() error {
	var erros []string
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("ProductCode: %s\n", p.ProductCode)
	fmt.Printf("Description: %s\n", p.Description)
	fmt.Printf("Width: %.2f\n", p.Width)
	fmt.Printf("Height: %.2f\n", p.Height)
	fmt.Printf("Length: %.2f\n", p.Length)
	fmt.Printf("NetWeight: %.2f\n", p.NetWeight)
	fmt.Printf("ExpirationRate: %.2f\n", p.ExpirationRate)
	fmt.Printf("RecommendedFreezingTemperature: %.2f\n", p.RecommendedFreezingTemperature)
	fmt.Printf("FreezingRate: %.2f\n", p.FreezingRate)
	fmt.Printf("ProductTypeID: %d\n", p.ProductTypeID)
	fmt.Printf("SellerID: %d\n", p.SellerID)
	if p.ProductCode == "" {
		erros = append(erros, "ProductCode não pode estar vazio")
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

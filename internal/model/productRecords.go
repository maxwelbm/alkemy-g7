package model

import (
	"fmt"
	"strings"
	"time"
)

type ProductRecords struct {
	ID             int       `json:"id"`
	LastUpdateDate time.Time `json:"last_update_date"`
	PurchasePrice  float64   `json:"purchase_price"`
	SalePrice      float64   `json:"sale_price"`
	ProductId      int       `json:"product_id"`
}

type ProductRecordsReport struct {
	ProductId    int    `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int    `json:"records_count"`
}

func (p *ProductRecords) Validate() error {
	var erros []string
	if p.PurchasePrice == 0.0 || p.PurchasePrice < 0.0 {
		erros = append(erros, "Purchase price invalid")
	} else if p.SalePrice == 0.0 || p.SalePrice < 0.0 {
		erros = append(erros, "ProductCode não pode estar vazio")
	}
	if len(erros) > 0 {
		return fmt.Errorf("Validations errors: %s", strings.Join(erros, "; "))
	}
	return nil
}

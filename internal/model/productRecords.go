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
	ProductID      int       `json:"product_id"`
}

type ProductRecordsReport struct {
	ProductID    int    `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int    `json:"records_count"`
}

func (p *ProductRecords) Validate() error {
	var errors []string

	purchasePrice := p.PurchasePrice
	if purchasePrice == 0.0 || purchasePrice < 0.0 {
		errors = append(errors, "Purchase price is invalid")
	}

	salePrice := p.SalePrice
	if salePrice == 0.0 || salePrice < 0.0 {
		errors = append(errors, "Sale Price is invalid")
	}

	if len(errors) > 0 {
		return fmt.Errorf("%s", strings.Join(errors, "; "))
	}
	
	return nil
}

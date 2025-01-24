package model

import (
	"fmt"
	"strings"
	"time"
)

type PurchaseOrder struct {
	ID              int       `json:"id" example:"1"`
	OrderNumber     string    `json:"order_number" example:"ON001"`
	OrderDate       time.Time `json:"order_date" example:"2025-01-01T00:00:00Z"`
	TrackingCode    string    `json:"tracking_code" example:"TC001"`
	BuyerID         int       `json:"buyer_id" example:"1"`
	ProductRecordID int       `json:"product_record_id" example:"1"`
}

func (p *PurchaseOrder) ValidateEmptyFields() error {
	var fieldsEmpty []string

	if p.OrderNumber == "" {
		fieldsEmpty = append(fieldsEmpty, "order_number")
	}

	if p.OrderDate.IsZero() {
		fieldsEmpty = append(fieldsEmpty, "order_date")
	}

	if p.TrackingCode == "" {
		fieldsEmpty = append(fieldsEmpty, "tracking_code")
	}

	if p.BuyerID == 0 {
		fieldsEmpty = append(fieldsEmpty, "buyer_id")
	}

	if p.ProductRecordID == 0 {
		fieldsEmpty = append(fieldsEmpty, "product_record_id")
	}

	if len(fieldsEmpty) > 0 {
		return fmt.Errorf("Field(s) %s cannot be empty", strings.Join(fieldsEmpty, ","))
	}

	return nil
}

type PurchaseOrderResponseSwagger struct {
	Data []PurchaseOrder `json:"data"`
}

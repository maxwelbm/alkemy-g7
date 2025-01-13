package model

import (
	"fmt"
	"strings"
	"time"
)

type PurchaseOrder struct {
	Id              int       `json:"id"`
	OrderNumber     string    `json:"order_number"`
	OrderDate       time.Time `json:"order_date"`
	TrackingCode    string    `json:"tracking_code"`
	BuyerId         int       `json:"buyer_id"`
	ProductRecordId int       `json:"product_record_id"`
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
	if p.BuyerId == 0 {
		fieldsEmpty = append(fieldsEmpty, "buyer_id")
	}
	if p.ProductRecordId == 0 {
		fieldsEmpty = append(fieldsEmpty, "product_record_id")
	}

	if len(fieldsEmpty) > 0 {
		return fmt.Errorf("Field(s) %s cannot be empty", strings.Join(fieldsEmpty, ","))
	}

	return nil

}

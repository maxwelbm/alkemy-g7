package model

import (
	"fmt"
	"strings"
)

type Buyer struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerPurchaseOrder struct {
	Id                  int    `json:"id"`
	CardNumberId        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}

func (b *Buyer) ValidateEmptyFields(isPatch bool) error {
	var fieldsEmpty []string

	if b.CardNumberId == "" {
		fieldsEmpty = append(fieldsEmpty, "card_number_id")
	}
	if b.FirstName == "" {
		fieldsEmpty = append(fieldsEmpty, "first_name")
	}
	if b.LastName == "" {
		fieldsEmpty = append(fieldsEmpty, "last_name")
	}

	if !isPatch && len(fieldsEmpty) > 0 {
		return fmt.Errorf("Field(s) %s cannot be empty", strings.Join(fieldsEmpty, ","))
	}

	if isPatch && len(fieldsEmpty) == 3 {
		return fmt.Errorf("At least one field must be filled in")
	}

	return nil

}

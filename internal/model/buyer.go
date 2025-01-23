package model

import (
	"fmt"
	"strings"
)

type Buyer struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerPurchaseOrder struct {
	ID                  int    `json:"id"`
	CardNumberID        string `json:"card_number_id"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	PurchaseOrdersCount int    `json:"purchase_orders_count"`
}

func (b *Buyer) ValidateEmptyFields(isPatch bool) error {
	var fieldsEmpty []string

	if b.CardNumberID == "" {
		fieldsEmpty = append(fieldsEmpty, "card_number_id")
	}

	if b.FirstName == "" {
		fieldsEmpty = append(fieldsEmpty, "first_name")
	}

	if b.LastName == "" {
		fieldsEmpty = append(fieldsEmpty, "last_name")
	}

	if !isPatch && len(fieldsEmpty) > 0 {
		return fmt.Errorf("field(s) %s cannot be empty", strings.Join(fieldsEmpty, ", "))
	}

	if isPatch && len(fieldsEmpty) == 3 {
		return fmt.Errorf("at least one field must be filled in")
	}

	return nil
}

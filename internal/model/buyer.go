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

func (b *Buyer) ValidateEmptyFields() error {
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

	if len(fieldsEmpty) == 0 {
		return nil
	} else if len(fieldsEmpty) == 1 {
		return fmt.Errorf("Field %s cannot be empty", strings.Join(fieldsEmpty, ","))
	}

	return fmt.Errorf("Fields %s cannot be empty", strings.Join(fieldsEmpty, ","))

}

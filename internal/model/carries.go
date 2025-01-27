package model

import (
	"fmt"
	"strings"
)

type Carries struct {
	ID          int    `json:"id"`
	CID         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}

func (c *Carries) ValidateEmptyFields(isPatch bool) error {
	var fieldsEmpty []string

	if c.CID == "" {
		fieldsEmpty = append(fieldsEmpty, "cid")
	}

	if c.CompanyName == "" {
		fieldsEmpty = append(fieldsEmpty, "company_name")
	}

	if c.Address == "" {
		fieldsEmpty = append(fieldsEmpty, "address")
	}

	if c.Telephone == "" {
		fieldsEmpty = append(fieldsEmpty, "telephone")
	}

	if c.LocalityID == 0 {
		fieldsEmpty = append(fieldsEmpty, "locality_id")
	}

	if len(fieldsEmpty) > 0 {
		if isPatch {
			return fmt.Errorf("the following fields are empty: %v", strings.Join(fieldsEmpty, ", "))
		} else {
			return fmt.Errorf("the following fields are required: %v", strings.Join(fieldsEmpty, ", "))
		}
	}

	return nil
}

type CarrierResponseSwagger struct {
	Data []Carries `json:"data"`
}

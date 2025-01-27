package model

import (
	"fmt"
	"strings"
)

type WareHouse struct {
	Id                 int    `json:"id"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WareHouseCode      string `json:"warehouse_code"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}

func (w *WareHouse) ValidateEmptyFields(isPatch bool) error {
	var fieldsEmpty []string

	if w.Address == "" {
		fieldsEmpty = append(fieldsEmpty, "address")
	}

	if w.Telephone == "" {
		fieldsEmpty = append(fieldsEmpty, "telephone")
	}

	if w.WareHouseCode == "" {
		fieldsEmpty = append(fieldsEmpty, "warehouse_code")
	}

	if w.MinimunCapacity <= 0 {
		fieldsEmpty = append(fieldsEmpty, "minimun_capacity")
	}

	if w.MinimunTemperature <= 0 {
		fieldsEmpty = append(fieldsEmpty, "minimun_temperature")
	}

	if !isPatch {
		if len(fieldsEmpty) > 0 {
			return fmt.Errorf("Field(s) %s cannot be empty or invalid", strings.Join(fieldsEmpty, ", "))
		}
	} else {
		if len(fieldsEmpty) == 5 {
			return fmt.Errorf("at least one field must be filled in")
		}
	}

	return nil
}

package model

import (
	"errors"
	"strings"
)

type Section struct {
	ID                 int     `json:"id"`
	SectionNumber      string  `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseID        int     `json:"warehouse_id"`
	ProductTypeID      int     `json:"product_type_id"`
}

type SectionProductBatches struct {
	ID            int    `json:"id"`
	SectionNumber string `json:"section_number"`
	ProductsCount int    `json:"products_count"`
}

func (s *Section) Validate() error {
	var errorMessages []string

	if s.SectionNumber == "" {
		errorMessages = append(errorMessages, "SectionNumber não pode ser vazio")
	}

	if s.CurrentTemperature == 0 {
		errorMessages = append(errorMessages, "CurrentTemperature não pode ser vazio")
	}

	if s.MinimumTemperature == 0 {
		errorMessages = append(errorMessages, "MinimumTemperature não pode ser vazio")
	}

	if s.CurrentCapacity == 0 {
		errorMessages = append(errorMessages, "CurrentCapacity não pode ser vazio")
	}

	if s.MinimumCapacity == 0 {
		errorMessages = append(errorMessages, "MinimumCapacity não pode ser vazio")
	}

	if s.MaximumCapacity == 0 {
		errorMessages = append(errorMessages, "MaximumCapacity não pode ser vazio")
	}

	if s.WarehouseID == 0 {
		errorMessages = append(errorMessages, "WarehouseID não pode ser vazio")
	}

	if s.ProductTypeID == 0 {
		errorMessages = append(errorMessages, "ProductTypeID não pode ser vazio")
	}

	if len(errorMessages) > 0 {
		return errors.New(strings.Join(errorMessages, "; "))
	}

	return nil
}

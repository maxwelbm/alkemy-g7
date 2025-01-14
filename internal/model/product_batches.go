package model

import (
	"errors"
	"strings"
	"time"
)

type ProductBatches struct {
	ID                 int
	BatchNumber        string
	CurrentQuantity    int
	CurrentTemperature float64
	MinimumTemperature float64
	DueDate            time.Time
	InitialQuantity    int
	ManufacturingDate  time.Time
	ManufacturingHour  int
	ProductID          int
	SectionID          int
}

func (pb *ProductBatches) Validate() error {
	var errorMessages []string

	if pb.BatchNumber == "" {
		errorMessages = append(errorMessages, "BatchNumber cannot be empty")
	}
	if pb.CurrentQuantity <= 0 {
		errorMessages = append(errorMessages, "CurrentQuantity must be greater than zero")
	}
	if pb.DueDate.IsZero() {
		errorMessages = append(errorMessages, "DueDate cannot be empty")
	}
	if pb.InitialQuantity < 0 {
		errorMessages = append(errorMessages, "InitialQuantity must be greater than zero")
	}
	if pb.ManufacturingDate.IsZero() {
		errorMessages = append(errorMessages, "ManufacturingDate cannot be empty")
	}
	if pb.ManufacturingHour < 0 || pb.ManufacturingHour > 23 {
		errorMessages = append(errorMessages, "ManufacturingHour must be between 0 and 23")
	}
	if pb.ProductID <= 0 {
		errorMessages = append(errorMessages, "ProductID must be greater than zero")
	}
	if pb.SectionID <= 0 {
		errorMessages = append(errorMessages, "SectionID must be greater than zero")
	}

	if len(errorMessages) > 0 {
		return errors.New(strings.Join(errorMessages, "; "))
	}
	return nil
}

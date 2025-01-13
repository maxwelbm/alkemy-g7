package model

import "time"

type ProductBatches struct {
	ID                 int
	BatchNumber        string
	CurrentQuantity    int
	CurrentTemperature float64
	MinimumTeperature  float64
	DueDate            time.Time
	InitialQuantity    int
	ManufacturingDate  time.Time
	ManufacturingHour  time.Time
	ProductID          int
	SectionID          int
}

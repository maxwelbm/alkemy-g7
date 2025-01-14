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
	ManufacturingHour  int
	ProductID          int
	SectionID          int
}

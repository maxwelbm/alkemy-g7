package model

type Section struct {
	ID                 int
	SectionNumber      int
	CurrentTemperature int
	MinimumTemperature int
	CurrentCapacity    int
	MinimumCapacity    int
	MaximumCapacity    int
	WarehouseID        int
	ProductTypeID      int
}

package model

type WareHouse struct {
	Id                 int    `json:"id"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WareHouseCode      string `json:"warehouse_code"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}

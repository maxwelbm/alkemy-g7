package model

type WareHouse struct {
	Id                 int
	Adress             string
	Telephone          string
	WareHouseCode      string
	MinimunCapacity    int
	MinimunTemperature int
}

type WareHouseJson struct {
	Id                 int    `json:"id"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WareHouseCode      string `json:"warehouse_code"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}

type WareHouseRes struct {
	Data []WareHouseJson `json:"data"`
}

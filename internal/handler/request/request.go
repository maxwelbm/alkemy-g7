package request

type RequestBody struct {
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	WareHouseCode      string `json:"warehouse_code"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature int    `json:"minimun_temperature"`
}

func IsValidateFields(reqBody RequestBody) bool {
	if reqBody.Address == "" || reqBody.Telephone == "" || reqBody.WareHouseCode == "" || reqBody.MinimunCapacity == 0 || reqBody.MinimunTemperature == 0 {
		return false
	}
	return true
}

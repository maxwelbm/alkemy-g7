package model

type Buyer struct {
	Id           int    `json: "id"`
	CardNumberId int    `json: "card_number_id"`
	FirstName    string `json: "first_name"`
	LastName     string `json: "last_name"`
}

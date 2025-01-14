package model

import (
	"errors"
)

type Locality struct {
	ID       string `json:"id"`
	Locality string `json:"locality_name"`
	Province string `json:"province_id"`
}

type LocalitiesJSONSellers struct {
	ID       string `json:"locality_id"`
	Locality string `json:"locality_name"`
	Sellers string  `json:"sellers_count"`
}

type LocalitiesJSONCarries struct {
	ID       string `json:"locality_id"`
	Locality string `json:"locality_name"`
	Carries string  `json:"carries_count"`
}

var (
	ErrorLocalityNotFound          error = errors.New("Locality not found")
	ErrorIDAlreadyExist            error = errors.New("Locality ID already exist")
	ErrorMissingLocalityID         error = errors.New("Missing int ID")
	ErrorInvalidLocalityJSONFormat error = errors.New("Invalid JSON request format attribute")
	ErrorNullLocalityAttribute     error = errors.New("Invalid body, empty value received.")
)

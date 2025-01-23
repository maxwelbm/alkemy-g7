package responses

import "github.com/maxwelbm/alkemy-g7.git/internal/model"

type ResponseBody struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func CreateResponseBody(m string, d any) *ResponseBody {
	return &ResponseBody{
		Message: m,
		Data:    d,
	}
}

type BuyerResponseSwagger struct {
	Data []model.Buyer `json:"data"`
}

type ErrorResponseSwagger struct {
	Message string `json:"message" example:"Error message"`
}

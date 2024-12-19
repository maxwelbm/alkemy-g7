package responses

type ResponseBody struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func CreateResponseBody(m string, d any) *ResponseBody {
	return &ResponseBody{
		Message: m,
		Data:    d,
	}
}

package domain

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(message string, data interface{}) *Response {
	return &Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, data interface{}) *Response {
	return &Response{
		Status:  false,
		Message: message,
		Data:    data,
	}
}

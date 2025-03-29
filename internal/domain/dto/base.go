package dto

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ValidationError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Field   string `json:"field"`
}

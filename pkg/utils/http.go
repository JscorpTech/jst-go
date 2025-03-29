package utils

import "github.com/JscorpTech/jst-go/internal/domain/dto"

func SuccessResponse(message string, data interface{}) *dto.Response {
	return &dto.Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, data interface{}) *dto.Response {
	return &dto.Response{
		Status:  false,
		Message: message,
		Data:    data,
	}
}

package services

import (
	"fmt"
	"github.com/JscorpTech/jst-go/domain"
)

type SmsService struct{}

func NewSmsService() domain.SmsService {
	return &SmsService{}
}

func (s *SmsService) SendSMS(from, to string) error {
	fmt.Println("Sending sms to " + from + " to " + to)
	return nil
}

package services

import (
	"fmt"
)

type SmsService struct{}

func NewSmsService() *SmsService {
	return &SmsService{}
}

func (s *SmsService) SendSMS(from, to string) error {
	fmt.Println("Sending sms to " + from + " to " + to)
	return nil
}

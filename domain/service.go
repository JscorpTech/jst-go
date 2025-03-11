package domain

type SmsService interface {
	SendSMS(string, string) error
}

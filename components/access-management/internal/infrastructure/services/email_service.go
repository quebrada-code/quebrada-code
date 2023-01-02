package services

type IEmailService interface {
	SendEmail(to []string, subject string, template string)
}

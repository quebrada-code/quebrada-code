package senders

type ISender interface {
	Send(to []string, subject string, template string, data interface{}) error
}

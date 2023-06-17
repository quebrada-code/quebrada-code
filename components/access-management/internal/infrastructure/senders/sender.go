package senders

import "quebrada_api/resources"

type ISender interface {
	Send(to []string, subject string, template resources.TemplateEmail, data interface{}) error
}

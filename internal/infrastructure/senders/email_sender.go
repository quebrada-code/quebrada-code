package senders

import (
	"bytes"
	"github.com/go-gomail/gomail"
	"github.com/pkg/errors"
	"html/template"
	"quebrada_api/internal/config"
	"quebrada_api/pkg/logger"
)

type EmailSender struct {
	config config.STMPConfig
}

func NewEmailSender(config config.STMPConfig) *EmailSender {
	return &EmailSender{
		config: config,
	}
}

func (e *EmailSender) generateBodyFromHTML(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		logger.Errorf("failed to parse file %s:%s", templateFileName, err.Error())

		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	return body, nil
}

func (e *EmailSender) validate() error {
	return nil
}

func (e *EmailSender) Send(to []string, subject string, template string, data interface{}) error {
	if err := e.validate(); err != nil {
		return err
	}

	body, err := e.generateBodyFromHTML(template, data)
	if err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", e.config.User)
	msg.SetHeader("To", to[0])
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.config.Host, e.config.Port, e.config.User, e.config.Password)
	if err := dialer.DialAndSend(msg); err != nil {
		return errors.Wrap(err, "failed to sent email via smtp")
	}

	return nil
}

package senders

import (
	"github.com/stretchr/testify/assert"
	config2 "quebrada_api/internal/config"
	"quebrada_api/internal/infrastructure/senders"
	"quebrada_api/resources"
	"testing"
)

type TestEmail struct {
	VerificationLink string
}

func TestSendEmail(t *testing.T) {

	var config = config2.STMPConfig{
		Host:     "smtp.mailgun.org",
		Port:     25,
		User:     "teste@marcosmota.com",
		Password: "e167be3236b6486ee734f23cf5530bb7-48d7d97c-5de5a08f",
	}
	emailSender := senders.NewEmailSender(config)
	err := resources.LoadTemplates()
	assert.Nil(t, err)

	test := TestEmail{VerificationLink: "teste"}

	err = emailSender.Send([]string{"marcos.mota287@gmail.com"}, "TESTE QUEBRADA", resources.WelcomeTemplate, test)
	assert.Nil(t, err)
}

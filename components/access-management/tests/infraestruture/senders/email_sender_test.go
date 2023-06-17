package senders

import (
	"github.com/stretchr/testify/assert"
	"quebrada_api/resources"
	"testing"
)

type TestEmail struct {
	VerificationLink string
}

func TestSendEmail(t *testing.T) {
	/*
		var config = config2.STMPConfig{
			Host:     "smtp.mailgun.org",
			Port:     25,
			User:     "teste@marcosmota.com",
			Password: "teste",
		}*/
	//emailSender := senders.NewEmailSender(config)
	err := resources.LoadTemplates()
	assert.Nil(t, err)

	//test := TestEmail{VerificationLink: "teste"}

	//err = emailSender.Send([]string{"marcos.mota287@gmail.com"}, "TESTE QUEBRADA", resources.VerificationEmailTemplate, test)
	//assert.Nil(t, err)
}

package senders

type TestEmail struct {
	VerificationLink string
}

//
//func TestSendEmail(t *testing.T) {
//	var (
//		from     string = "teste@marcosmota.com"
//		host     string = "smtp.mailgun.org"
//		port     int    = 25
//		password string = "e167be3236b6486ee734f23cf5530bb7-48d7d97c-5de5a08f"
//	)
//	emailSender := senders.NewEmailSender(host, from, password, port)
//
//	test := TestEmail{VerificationLink: "teste"}
//
//	err := emailSender.Send([]string{"marcos.mota287@gmail.com"}, "TESTE QUEBRADA", "/Users/marcos.lopes/projects/pessoal/plataform/quebrada_api/resources/welcome.html", test)
//	assert.Nil(t, err)
//}

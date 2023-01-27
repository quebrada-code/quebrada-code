package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"
)

type ResetPassword struct {
	Email string
}

func (a ResetPassword) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
	)
}

type ConfirmResetPassword struct {
	Email       string
	Token       string
	NewPassword string
}

func (a ConfirmResetPassword) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Token, validation.Required),
		validation.Field(&a.NewPassword, validation.Required),
	)
}

type UserVerificationCode struct {
	UserId uint
	Code   string
}

func (a UserVerificationCode) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.UserId, validation.Required),
		validation.Field(&a.Code, validation.Required),
	)
}

type SignUpModel struct {
	Name            string
	Email           string
	Nickname        string
	ZipCode         string
	Password        string
	ConfirmPassword string
}

func (a SignUpModel) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required.Error("Campo nome é obrigatório."), validation.Length(3, 50)),
		validation.Field(&a.Email, validation.Required.Error("Campo e-mail é obrigatório."), is.Email.Error("E-mail é inválido")),
		validation.Field(&a.Nickname, validation.Required.Error("Campo nickname é obrigatório."), is.LowerCase, validation.Length(3, 10)),
		validation.Field(&a.ZipCode, validation.Required.Error("Campo CEP é obrigatório."), validation.Match(regexp.MustCompile("[0-9]{5}-[0-9]{3}"))),
	)
}

type EmailCredential struct {
	Email    string
	Password string
}

func (a EmailCredential) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, validation.Required),
	)
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
}

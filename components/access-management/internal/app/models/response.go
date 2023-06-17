package models

type Response struct {
	Message string `json:"message"`
}

type BadRequestMessage struct {
	Errors map[string]error `json:"errors"`
}

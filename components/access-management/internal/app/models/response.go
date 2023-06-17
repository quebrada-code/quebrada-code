package models

type Response struct {
	Message string `json:"message"`
}

type ResponseModel struct {
	Data interface{} `json:"data,omitempty"`
}

type ResponseList struct {
	Data []interface{} `json:"data"`
}

type BadRequestMessage struct {
	Errors map[string]error `json:"errors"`
}

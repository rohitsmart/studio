package model

type ErrorResponse struct {
	ErrorDescription string `json:"error_description"`
	UserDescription  string `json:"user_description"`
	Code             string `json:"code"`
}

package models

type CommonResponse struct {
	Message string `json:"message"`
	Succes  bool   `json:"success"`
}

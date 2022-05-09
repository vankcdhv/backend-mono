package types

import (
	"backend-mono/cmd/database/model"
)

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CreateUserResponse struct {
	Code  int         `json:"code"`
	Data  *model.User `json:"data"`
	Error string      `json:"error"`
}

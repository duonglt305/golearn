package users

import (
	"golearn/common"
	"net/http"
)

func NewLoginError(err error) *common.GenericError {
	ge := common.GenericError{
		Code:    http.StatusUnauthorized,
		Message: "The credentials invalid.",
	}
	ge.SetError(err)
	return &ge
}

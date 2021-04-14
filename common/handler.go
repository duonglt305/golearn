package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type GenericError struct {
	error   `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (g *GenericError) SetError(err error) {
	g.error = err
}

func NewGenericError(err error) *GenericError {
	return &GenericError{
		error:   err,
		Code:    http.StatusBadRequest,
		Message: "Oops! Something went wrong. Please try again in a bit.",
	}
}

func NewUnauthenticatedError(err error) *GenericError {
	return &GenericError{
		error:   err,
		Code:    http.StatusUnauthorized,
		Message: "Unauthenticated.",
	}
}

type ValidationError struct {
	*GenericError
	Errors map[string]interface{} `json:"errors"`
}

func NewValidationError(err error) *ValidationError {
	return &ValidationError{
		GenericError: &GenericError{
			error:   err,
			Code:    http.StatusUnprocessableEntity,
			Message: "The given data was invalid.",
		},
		Errors: GetValidationMessages(err),
	}
}

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			errors := c.Errors.ByType(gin.ErrorTypeAny)
			err := errors[0].Err
			switch err.(type) {
			case *GenericError:
				parsed := err.(*GenericError)
				c.AbortWithStatusJSON(parsed.Code, parsed)
			case validator.ValidationErrors:
				parsed := NewValidationError(err)
				c.AbortWithStatusJSON(parsed.Code, parsed)
			default:
				parsed := NewGenericError(err)
				c.AbortWithStatusJSON(parsed.Code, parsed)
			}
			return
		}

	}
}

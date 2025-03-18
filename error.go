package rush

import (
	"errors"
	"log"
	"net/http"
)

type ErrorRush struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorHandlerFunc func(ctx *Context, err error)

func (e ErrorRush) Error() string {
	return e.Message
}

func (ctx *Context) NewError(status int, err string) error {
	return ErrorRush{
		Code:    status,
		Message: err,
	}
}

func defaultErrorHandler(ctx *Context, err error) {
	var apiErr ErrorRush
	if errors.As(err, &apiErr) {
		if jsonErr := ctx.Response.Status(apiErr.Code).JSON(apiErr); jsonErr != nil {
			log.Println(jsonErr)
		}
	} else {
		if jsonErr := ctx.Response.Status(http.StatusInternalServerError).JSON(ErrorRush{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}); jsonErr != nil {
			log.Println(jsonErr)
		}
	}
	return
}

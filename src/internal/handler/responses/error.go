package responses

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}

type ValidationErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func NewValidationErrorResponse(
	ctx *gin.Context,
	message string,
	err error,
) {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		errorsMap := getErrorsMap(ve)

		logrus.Error(fmt.Sprintf("%s : %s", message, errorsMap))
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ValidationErrorResponse{message, errorsMap})

		return
	}

	logrus.Error(message)
	ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{message})
}

func getErrorsMap(ve validator.ValidationErrors) map[string]string {
	out := map[string]string{}

	for _, fe := range ve {
		field := fe.Field()
		errMsg := fe.Tag()
		out[field] = errMsg
	}

	return out
}

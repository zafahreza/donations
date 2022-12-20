package exception

import (
	"donations/model/client"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func NewNotFoundError(c *gin.Context, err error) {

	body := client.HttpResponse{
		Code:   http.StatusNotFound,
		Status: "not found",
		Data:   string(err.Error()),
	}

	c.JSON(http.StatusNotFound, body)
}

func validatorError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func Unprocessable(c *gin.Context, err error) {
	_, ok := err.(validator.ValidationErrors)

	if err != nil && ok {
		errors := validatorError(err)

		body := client.HttpResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: "not found",
			Data:   gin.H{"error": errors},
		}

		c.JSON(http.StatusUnprocessableEntity, body)
	} else {
		InternalServerError(c, err)
	}
	return
}

func InternalServerError(c *gin.Context, err error) {
	if err != nil {
		body := client.HttpResponse{
			Code:   http.StatusInternalServerError,
			Status: "not found",
			Data:   err.Error(),
		}

		c.JSON(http.StatusInternalServerError, body)
	}
	return
}

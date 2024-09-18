package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonRespose struct {
	Status  int    `json:"Status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()

func (app *Config) validateBody(c *gin.Context, data any) error {
	if err := c.BindJSON(&data); err != nil {
		return err
	}
	if err := validate.Struct(&data); err != nil {
		return err
	}
	return nil
}

func (app *Config) writJson(c *gin.Context, status int, data any) {
	c.JSON(status, jsonRespose{Status: status, Message: "success", Data: data})
}

func (app *Config) errorJson(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.JSON(statusCode, jsonRespose{Status: statusCode, Message: err.Error()})
}

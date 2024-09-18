package api

import "github.com/gin-gonic/gin"

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.POST("/otp", app.HandleSendSMS())
	app.Router.POST("/verifyOTP", app.HandleVerifySMS())
}

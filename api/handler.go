package api

import (
	"context"
	"net/http"
	"otp/auth/api/data"
	"time"

	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func (app *Config) HandleSendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}

		_, err := app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJson(c, err)
			return
		}
		app.writJson(c, http.StatusAccepted, "OTP send successfully")
	}
}

func (app *Config) HandleVerifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()

		app.validateBody(c, &payload)

		newData := data.VerifyData{
			User: payload.User,
			Code: payload.Code,
		}

		err := app.twilioVerifyOTP(newData.User.PhoneNumber, newData.Code)
		if err != nil {
			app.errorJson(c, err)
			return
		}

		app.writJson(c, http.StatusAccepted, "OTP verified successfully")
	}
}

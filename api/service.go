package api

import (
	helperfunc "otp/auth/api/helperFunc"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: helperfunc.EnvAccountSID(),
	Password: helperfunc.EnvAuthToken(),
})

func (app *Config) twilioSendOTP(phoneNumber string) (string, error) {
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(helperfunc.EnvServiseSID(), params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil

}

func (app *Config) twilioVerifyOTP(phoneNumber string, code string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(helperfunc.EnvServiseSID(), params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}
	return nil
}

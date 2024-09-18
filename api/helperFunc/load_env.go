package helperfunc

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func AbsPath() string {
	abspath, _ := filepath.Abs("../.env")
	return abspath
}

func EnvAccountSID() string {
	abspath := AbsPath()
	err := godotenv.Load(abspath)
	if err != nil {
		log.Fatalln(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func EnvAuthToken() string {
	abspath := AbsPath()
	err := godotenv.Load(abspath)
	if err != nil {
		log.Fatal("Error loading the env file")
	}
	return os.Getenv("TWILIO_AUTHTOKEN")
}

func EnvServiseSID() string {
	abspath := AbsPath()
	err := godotenv.Load(abspath)
	if err != nil {
		log.Fatal("Error loading the env file")
	}
	return os.Getenv("TWILIO_SERVICES_SID")
}

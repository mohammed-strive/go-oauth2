package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mohammed-strive/go-oauth2/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleCredentials models.GoogleCredentials

var DEFAULT_SCOPES = []string{
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/userinfo.profile",
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading env file: %v", err)
	}
	credentialsFilePath := os.Getenv("GOOGLE_CREDENTIALS_FILE")
	credentialsFile, err := ioutil.ReadFile(credentialsFilePath)
	if err != nil {
		log.Fatalf("unable to read google credentials file: %v", err)
	}

	err = json.Unmarshal(credentialsFile, &googleCredentials)
	if err != nil {
		log.Fatalf("unable to unmarshal json data: %v", err)
	}
}

type Config struct {
	GoogleLoginConfig oauth2.Config
}

var AppConfig Config

func GoogleConfig() oauth2.Config {
	AppConfig.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  googleCredentials.Web.RedirectUrls[0],
		ClientID:     googleCredentials.Web.ClientID,
		ClientSecret: googleCredentials.Web.ClientSecret,
		Scopes:       DEFAULT_SCOPES,
		Endpoint:     google.Endpoint,
	}

	return AppConfig.GoogleLoginConfig
}

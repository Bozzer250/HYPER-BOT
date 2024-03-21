package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvIsProd() bool {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	isProd := os.Getenv("ENV") == "production"
	return isProd
}

func GetGoogleSecretString() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("GOOGLE_SECRET_STRING")
}

func GetRedisUrl() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("REDIS_URL")
}

func EnvMongoURI() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("MONGODB_URI")
}

func EnvPort() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("PORT")
}

func EnvPindoToken() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("PINDO_TOKEN")
}

func EnvNewRelicKey() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("NEW_RELIC_KEY")
}

func EnvDdApiKey() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("DD_API_KEY")
}

func GetGoogleClientId() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("GOOGLE_CLIENT_ID")
}

func GetGoogleSecret() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("GOOGLE_CLIENT_SECRET")
}

func GetGoogleOauthRedirectUrl() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")
}

func GetIntouchUserName() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("INTOUCH_USERNAME")
}

func GetIntouchAccountNbr() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("INTOUCH_ACCOUNT_NUMBER")
}

func GetIntouchPartnerPassword() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("INTOUCH_PARTNER_PSSWORD")
}

func GetIntouchCallBackUrl() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("INTOUCH_CALLBACK_URL")
}

func GetSlackNotifyChannel() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("SLACK_NOTIFY_CHANNEL")
}

func GetSlackFailChannel() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("SLACK_FAIL_CHANNEL")
}

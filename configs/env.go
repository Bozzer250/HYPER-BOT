package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func AppEnv() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("APP_ENV")
}

func EnvIsProd() bool {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	isProd := os.Getenv("ENV") == "production"
	return isProd
}

func GetUsdToRwfRate() int {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	rate, err := strconv.Atoi(os.Getenv("USD_TO_RWF"))
	if err != nil {
		return 1300
	}
	return rate
}

func GetPaypackId() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("PAYPACK_CLIENT_ID")
}

func GetPaypackSecret() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("PAYPACK_CLIENT_SECRET")
}

func GetRedisUrl() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("REDIS_URL")
}

func GetSessionKey() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("SESSION_KEY")
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

func EnvNenoApiKey() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	return os.Getenv("NENO_API_KEY")
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

func GetCashinUsdRate() int {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	rate, err := strconv.Atoi(os.Getenv("USD_TO_RWF"))
	if err != nil {
		return 1300
	}
	return rate
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

package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	AppName      string
	IsProduction string
	AppUrl       string
	AppPort      string

	DBHost                    string
	DBUser                    string
	DBPassword                string
	DBName                    string
	DBPort                    string
	DBMaxIdleConnection       int
	DBMaxOpenConnection       int
	DBConnectionMaxLifeMinute time.Duration

	RedisPassword string
	RedisHost     string
	RedisPort     string

	APISecret     string
	TokenDuration string

	AccessSecret  string
	RefreshSecret string
	AccessMinute  time.Duration
	RefreshMinute time.Duration

	MerchId              string
	MidtransIsProduction bool
	ClientKeySandbox     string
	ServerKeySandbox     string
	ClientKeyProduction  string
	ServerKeyProduction  string

	MailMailer      string
	MailHost        string
	MailPort        string
	MailUsername    string
	MailPassword    string
	MailEncryption  string
	MailFromAddress string
	MailFromName    string

	TransactionMinute time.Duration

	AdminName  string
	AdminEmail string
	AdminPhone string

	IsOpenGate      bool
	QrScanBehaviour string
}

func NewAppConfig() *AppConfig {
	midtransIsProduction, _ := strconv.ParseBool(getEnv("MIDTRANS_IS_PRODUCTION", "0"))
	accessMinute, _ := time.ParseDuration(getEnv("ACCESS_MINUTE", "15m"))
	refreshMinute, _ := time.ParseDuration(getEnv("ACCESS_MINUTE", "120m"))
	transactionMinute, _ := time.ParseDuration(getEnv("TRANSACTION_MINUTE", "15m"))
	dbMaxIdleConnection, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNECTION", "10"))
	dbMaxOpenConnection, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNECTION", "10"))
	dbConnectionMaxLifeMinute, _ := time.ParseDuration(getEnv("DB_CONNECTION_MAX_LIFE_MINUTE", "60m"))

	var appConfig = AppConfig{
		AppName:      getEnv("APP_NAME", "gmcgo"),
		IsProduction: getEnv("IS_PRODUCTION", "false"),
		AppUrl:       getEnv("APP_URL", "127.0.0.1"),
		AppPort:      getEnv("APP_PORT", "8080"),

		DBHost:                    getEnv("DB_HOST", "localhost"),
		DBUser:                    getEnv("DB_USER", "root"),
		DBPassword:                getEnv("DB_PASSWORD", "root"),
		DBName:                    getEnv("DB_NAME", "chatin_db"),
		DBPort:                    getEnv("DB_PORT", "5432"),
		DBMaxIdleConnection:       dbMaxIdleConnection,
		DBMaxOpenConnection:       dbMaxOpenConnection,
		DBConnectionMaxLifeMinute: dbConnectionMaxLifeMinute,

		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisHost:     getEnv("REDIS_HOST", "127.0.0.1"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),

		AccessSecret:  getEnv("ACCESS_SECRET", ""),
		RefreshSecret: getEnv("REFRESH_SECRET", ""),
		AccessMinute:  accessMinute,
		RefreshMinute: refreshMinute,

		MerchId:              getEnv("MERCH_ID", ""),
		MidtransIsProduction: midtransIsProduction,
		ClientKeySandbox:     getEnv("CLIENT_KEY_SANDBOX", ""),
		ServerKeySandbox:     getEnv("SERVER_KEY_SANDBOX", ""),
		ClientKeyProduction:  getEnv("CLIENT_KEY_PRODUCTION", ""),
		ServerKeyProduction:  getEnv("SERVER_KEY_PRODUCTION", ""),

		MailMailer:      getEnv("MAIL_MAILER", "smtp"),
		MailHost:        getEnv("MAIL_HOST", "smtp.gmail.com"),
		MailPort:        getEnv("MAIL_PORT", "465"),
		MailUsername:    getEnv("MAIL_USERNAME", ""),
		MailPassword:    getEnv("MAIL_PASSWORD", ""),
		MailEncryption:  getEnv("MAIL_ENCRYPTION", "ssl"),
		MailFromAddress: getEnv("MAIL_FROM_ADDRESS", ""),
		MailFromName:    getEnv("MAIL_FROM_NAME", "gmco"),

		TransactionMinute: transactionMinute,

		AdminName:  getEnv("ADMIN_NAME", ""),
		AdminEmail: getEnv("ADMIN_EMAIL", ""),
		AdminPhone: getEnv("ADMIN_PHONE", ""),

		IsOpenGate:      true,
		QrScanBehaviour: "open_gate", //open_gate, ticket_exchanging, default
	}
	return &appConfig

}

func getEnv(key, fallback string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

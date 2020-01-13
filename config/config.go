package config

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// DevelopmentMode indicates the backend is running in development mode
// This mode trades security for ease of use (Such as using unsecure cookies)
const DevelopmentMode = "development"

// ProductionMode indicates the backend is running in production mode
// This mode has the highest level of security
const ProductionMode = "production"

// DefaultMode is the default mode of the backend
const DefaultMode = ProductionMode

// Ctx is the context type used to reference the current configuration
type Ctx struct{}

// Configuration is the main configuration file loader. The format is expected to be a toml file
type Configuration struct {
	AppName         string
	Mode            string
	Domain          string
	AllowableOrigin string

	Address string
	Port    int32

	DBHost     string
	DBPort     int32
	DBName     string
	DBUser     string
	DBPassword string
	DBSSLMode  string

	JWTSecret   string
	JWTExpireIn time.Duration

	DebugMode bool
	LogFormat string

	MailJetPublicKey  string
	MailJetPrivateKey string

	PostMarkServerToken  string
	PostMarkAccountToken string

	EmailVerificationExpiresIn time.Duration
	PasswordResetExpiresIn     time.Duration

	TestDatabaseConnection string

	ReadTimeoutInSeconds  time.Duration
	WriteTimeoutInSeconds time.Duration
}

// LoadConfig loads the toml file located in the given path with the given basename
func LoadConfig(path string, configFileName string) *Configuration {
	config := viper.New()

	config.AutomaticEnv()
	config.SetEnvPrefix("DIGIDOC")
	config.SetDefault("APP_NAME", "Digidoc")
	config.SetDefault("MODE", DefaultMode)
	config.SetDefault("DOMAIN", "https://digidocapp.com")
	config.SetDefault("ALLOWABLE_ORIGIN", "digidocapp.com")

	config.SetDefault("ADDRESS", "127.0.0.1")
	config.SetDefault("PORT", "8080")

	config.SetDefault("DB_HOST", "http://localhost")
	config.SetDefault("DB_PORT", "5432")
	config.SetDefault("DB_NAME", "digidoc")
	config.SetDefault("DB_USER", "postgres")
	config.SetDefault("DB_PASSWORD", "postgres")
	config.SetDefault("DB_SSL_MODE", "disable")

	config.SetDefault("TEST_DB_HOST", "localhost")
	config.SetDefault("TEST_DB_PORT", "5433")
	config.SetDefault("TEST_DB_NAME", "digidoc")
	config.SetDefault("TEST_DB_USER", "postgres")
	config.SetDefault("TEST_DB_PASSWORD", "postgres")

	config.SetDefault("JWT_SECRET", "")
	config.SetDefault("JWT_EXPIRE_IN", "40000s")
	config.SetDefault("DEBUG_MODE", "false")
	config.SetDefault("LOG_FORMAT", "%{color}%{time:2006/01/02 15:04:05 -05:00 EST} [%{level:.6s}] %{shortfile} : %{color:reset}%{message}")

	config.SetDefault("MAILJET_PUBLIC_KEY", "")
	config.SetDefault("MAILJET_PRIVATE_KEY", "")

	config.SetDefault("POSTMARK_SERVER_TOKEN", "POSTMARK_API_TEST")
	config.SetDefault("POSTMARK_ACCOUNT_TOKEN", "POSTMARK_API_TEST")

	config.SetDefault("READ_TIMEOUT_IN_SECONDS", "15s")
	config.SetDefault("WRITE_TIMEOUT_IN_SECONDS", "15s")

	// Default 7 days
	config.SetDefault("EMAIL_VERIFICATION_EXPIRES_IN", "604800s")

	// Default 30 minutes
	config.SetDefault("PASSWORD_RESET_EXPIRES_IN", "1800s")

	return &Configuration{
		AppName:         config.GetString("APP_NAME"),
		Mode:            config.GetString("MODE"),
		Domain:          config.GetString("DOMAIN"),
		AllowableOrigin: config.GetString("ALLOWABLE_ORIGIN"),

		Address: config.GetString("ADDRESS"),
		Port:    config.GetInt32("PORT"),

		DBHost:     config.GetString("DB_HOST"),
		DBPort:     config.GetInt32("DB_PORT"),
		DBName:     config.GetString("DB_NAME"),
		DBUser:     config.GetString("DB_USER"),
		DBPassword: config.GetString("DB_PASSWORD"),
		DBSSLMode:  config.GetString("DB_SSL_MODE"),

		JWTSecret: config.GetString("JWT_SECRET"),

		DebugMode:   config.GetBool("DEBUG_MODE"),
		JWTExpireIn: config.GetDuration("JWT_EXPIRE_IN"),

		MailJetPublicKey:  config.GetString("MAILJET_PUBLIC_KEY"),
		MailJetPrivateKey: config.GetString("MAILJET_PRIVATE_KEY"),

		PostMarkServerToken:  config.GetString("POSTMARK_SERVER_TOKEN"),
		PostMarkAccountToken: config.GetString("POSTMARK_ACCOUNT_TOKEN"),

		EmailVerificationExpiresIn: config.GetDuration("EMAIL_VERIFICATION_EXPIRES_IN"),
		PasswordResetExpiresIn:     config.GetDuration("PASSWORD_RESET_EXPIRES_IN"),

		ReadTimeoutInSeconds:  config.GetDuration("READ_TIMEOUT_IN_SECONDS"),
		WriteTimeoutInSeconds: config.GetDuration("WRITE_TIMEOUT_IN_SECONDS"),

		TestDatabaseConnection: fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			config.GetString("TEST_DB_USER"), config.GetString("TEST_DB_PASSWORD"), config.GetString("TEST_DB_HOST"),
			config.GetInt32("TEST_DB_PORT"), config.GetString("TEST_DB_NAME"),
		),
	}
}

// Config returns the context configuration
// A config object must be set in every context, so if one does not exist, panic
func Config(ctx context.Context) *Configuration {
	config, ok := ctx.Value(Ctx{}).(*Configuration)
	if !ok {
		panic("Error: config not set in context")
	}
	return config
}

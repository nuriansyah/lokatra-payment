package configs

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"sync"
	"time"
)

// Config is a struct that will receive configuration options via environment
// variables.
type Config struct {
	App struct {
		CORS struct {
			AllowCredentials bool     `mapstructure:"ALLOW_CREDENTIALS"`
			AllowedHeaders   []string `mapstructure:"ALLOWED_HEADERS"`
			AllowedMethods   []string `mapstructure:"ALLOWED_METHODS"`
			AllowedOrigins   []string `mapstructure:"ALLOWED_ORIGINS"`
			Enable           bool     `mapstructure:"ENABLE"`
			MaxAgeSeconds    int      `mapstructure:"MAX_AGE_SECONDS"`
		}
		Name     string `mapstructure:"NAME"`
		Revision string `mapstructure:"REVISION"`
		URL      string `mapstructure:"URL"`
	}
	Cache struct {
		Redis struct {
			Expired struct {
			}
			Primary struct {
				DB       int    `mapstructure:"DB"`
				Host     string `mapstructure:"HOST"`
				Port     string `mapstructure:"PORT"`
				Password string `mapstructure:"PASSWORD"`
			}
		}
	}
	DB struct {
		Postgres struct {
			Read struct {
				Host            string        `mapstructure:"HOST"`
				Port            string        `mapstructure:"PORT"`
				User            string        `mapstructure:"USER"`
				Password        string        `mapstructure:"PASSWORD"`
				Name            string        `mapstructure:"NAME"`
				SSLMode         string        `mapstructure:"SSLMODE"`
				MaxConnLifetime time.Duration `mapstructure:"MAX_CONNECTION_LIFETIME"`
				MaxIdleConn     int           `mapstructure:"MAX_IDLE_CONNECTION"`
				MaxOpenConn     int           `mapstructure:"MAX_OPEN_CONNECTION"`
			}
			Write struct {
				Host            string        `mapstructure:"HOST"`
				Port            string        `mapstructure:"PORT"`
				User            string        `mapstructure:"USER"`
				Password        string        `mapstructure:"PASSWORD"`
				Name            string        `mapstructure:"NAME"`
				SSLMode         string        `mapstructure:"SSLMODE"`
				MaxConnLifetime time.Duration `mapstructure:"MAX_CONNECTION_LIFETIME"`
				MaxIdleConn     int           `mapstructure:"MAX_IDLE_CONNECTION"`
				MaxOpenConn     int           `mapstructure:"MAX_OPEN_CONNECTION"`
			}
		} `mapstructure:"PG"`
	}
	Externals struct {
		Providers struct {
			Midtrans struct {
				AccountID      string `mapstructure:"ACCOUNT_ID"`
				AccountLabel   string `mapstructure:"ACCOUNT_LABEL"`
				Enabled        bool   `mapstructure:"ENABLED"`
				BaseURL        string `mapstructure:"BASE_URL"`
				ChargePath     string `mapstructure:"CHARGE_PATH"`
				ServerKey      string `mapstructure:"SERVER_KEY"`
				TimeoutSeconds int    `mapstructure:"TIMEOUT_SECONDS"`
			} `mapstructure:"MIDTRANS"`
			Xendit struct {
				AccountID      string `mapstructure:"ACCOUNT_ID"`
				AccountLabel   string `mapstructure:"ACCOUNT_LABEL"`
				Enabled        bool   `mapstructure:"ENABLED"`
				BaseURL        string `mapstructure:"BASE_URL"`
				ChargePath     string `mapstructure:"CHARGE_PATH"`
				SecretKey      string `mapstructure:"SECRET_KEY"`
				WebhookToken   string `mapstructure:"WEBHOOK_TOKEN"`
				WebhookSecret  string `mapstructure:"WEBHOOK_SECRET"`
				TimeoutSeconds int    `mapstructure:"TIMEOUT_SECONDS"`
			} `mapstructure:"XENDIT"`
			Durianpay struct {
				AccountID      string `mapstructure:"ACCOUNT_ID"`
				Enabled        bool   `mapstructure:"ENABLED"`
				BaseURL        string `mapstructure:"BASE_URL"`
				APIKey         string `mapstructure:"API_KEY"`
				WebhookSecret  string `mapstructure:"WEBHOOK_SECRET"`
				TimeoutSeconds int    `mapstructure:"TIMEOUT_SECONDS"`
			} `mapstructure:"DURIANPAY"`
		} `mapstructure:"PROVIDERS"`
	} `mapstructure:"EXTERNALS"`
	Internal struct {
		Payment struct {
			AdminToken string `mapstructure:"ADMIN_TOKEN"`
			Routing    struct {
				RulesJSON          string `mapstructure:"RULES_JSON"`
				DefaultProviders   string `mapstructure:"DEFAULT_PROVIDERS"`
				MaxAttempts        int    `mapstructure:"MAX_ATTEMPTS"`
				FailureThreshold   int    `mapstructure:"FAILURE_THRESHOLD"`
				CooldownSeconds    int    `mapstructure:"COOLDOWN_SECONDS"`
				RetryBackoffMillis int    `mapstructure:"RETRY_BACKOFF_MILLIS"`
			} `mapstructure:"ROUTING"`
		} `mapstructure:"PAYMENT"`
	} `mapstructure:"INTERNAL"`
	Server struct {
		Env      string `mapstructure:"ENV"`
		LogLevel string `mapstructure:"LOG_LEVEL"`
		Port     string `mapstructure:"PORT"`
		Shutdown struct {
			CleanupPeriodSeconds int64 `mapstructure:"CLEANUP_PERIOD_SECONDS"`
			GracePeriodSeconds   int64 `mapstructure:"GRACE_PERIOD_SECONDS"`
		}
		Prometheus struct {
			Enable bool `mapstructure:"ENABLE"`
		} `mapstructure:"PROMETHEUS"`
	}
}

var (
	conf Config
	once sync.Once
)

// Get are responsible to load env and get data an return the struct
func Get() *Config {

	once.Do(func() {
		bindPaymentEnvironment()
		viper.AutomaticEnv()
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()

		if err != nil {
			log.Fatal().Err(err).Msg("Failed reading config file")
		}
		log.Info().Msg("Service configuration initialized.")
		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to unmarshal config")
		}
	})

	return &conf
}

func bindPaymentEnvironment() {
	bindings := map[string]string{
		"internal.payment.admin_token":                  "INTERNAL_PAYMENT_ADMIN_TOKEN",
		"internal.payment.routing.rules_json":           "INTERNAL_PAYMENT_ROUTING_RULES_JSON",
		"internal.payment.routing.default_providers":    "INTERNAL_PAYMENT_ROUTING_DEFAULT_PROVIDERS",
		"internal.payment.routing.max_attempts":         "INTERNAL_PAYMENT_ROUTING_MAX_ATTEMPTS",
		"internal.payment.routing.failure_threshold":    "INTERNAL_PAYMENT_ROUTING_FAILURE_THRESHOLD",
		"internal.payment.routing.cooldown_seconds":     "INTERNAL_PAYMENT_ROUTING_COOLDOWN_SECONDS",
		"internal.payment.routing.retry_backoff_millis": "INTERNAL_PAYMENT_ROUTING_RETRY_BACKOFF_MILLIS",
		"externals.providers.midtrans.enabled":          "EXTERNALS_PROVIDERS_MIDTRANS_ENABLED",
		"externals.providers.midtrans.account_id":       "EXTERNALS_PROVIDERS_MIDTRANS_ACCOUNT_ID",
		"externals.providers.midtrans.base_url":         "EXTERNALS_PROVIDERS_MIDTRANS_BASE_URL",
		"externals.providers.midtrans.server_key":       "EXTERNALS_PROVIDERS_MIDTRANS_SERVER_KEY",
		"externals.providers.xendit.enabled":            "EXTERNALS_PROVIDERS_XENDIT_ENABLED",
		"externals.providers.xendit.account_id":         "EXTERNALS_PROVIDERS_XENDIT_ACCOUNT_ID",
		"externals.providers.xendit.base_url":           "EXTERNALS_PROVIDERS_XENDIT_BASE_URL",
		"externals.providers.xendit.secret_key":         "EXTERNALS_PROVIDERS_XENDIT_SECRET_KEY",
		"externals.providers.xendit.webhook_token":      "EXTERNALS_PROVIDERS_XENDIT_WEBHOOK_TOKEN",
		"externals.providers.xendit.webhook_secret":     "EXTERNALS_PROVIDERS_XENDIT_WEBHOOK_SECRET",
		"externals.providers.durianpay.enabled":         "EXTERNALS_PROVIDERS_DURIANPAY_ENABLED",
		"externals.providers.durianpay.account_id":      "EXTERNALS_PROVIDERS_DURIANPAY_ACCOUNT_ID",
		"externals.providers.durianpay.base_url":        "EXTERNALS_PROVIDERS_DURIANPAY_BASE_URL",
		"externals.providers.durianpay.api_key":         "EXTERNALS_PROVIDERS_DURIANPAY_API_KEY",
		"externals.providers.durianpay.webhook_secret":  "EXTERNALS_PROVIDERS_DURIANPAY_WEBHOOK_SECRET",
	}
	for key, environment := range bindings {
		_ = viper.BindEnv(key, environment)
	}
}

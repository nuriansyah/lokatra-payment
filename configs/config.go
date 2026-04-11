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
				TimeoutSeconds int    `mapstructure:"TIMEOUT_SECONDS"`
			} `mapstructure:"XENDIT"`
		} `mapstructure:"PROVIDERS"`
	} `mapstructure:"EXTERNALS"`
	Internal struct {
		Payment struct {
			Routing struct {
				DefaultStrategy     string `mapstructure:"DEFAULT_STRATEGY"`
				DefaultUseCase      string `mapstructure:"DEFAULT_USE_CASE"`
				UseDatabaseFallback bool   `mapstructure:"USE_DATABASE_FALLBACK"`
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

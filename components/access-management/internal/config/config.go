package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

const (
	defaultHTTPPort = "8000"
	EnvLocal        = "local"
)

type (
	Config struct {
		Environment string
		Database    DatabaseConfig
		SMPT        STMPConfig
	}

	DatabaseConfig struct {
		ConnectionString string
	}

	KafkaConfig struct {
		ConnectionString string
	}

	STMPConfig struct {
		Host     string
		Port     int
		User     string
		Password string
	}
)

// Init populates Config struct with values from config file
// located at filepath and environment variables.
func Init(configsDir string) (*Config, error) {

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func unmarshal(cfg *Config) error {

	return viper.UnmarshalKey("database", &cfg.Database)
}

func setFromEnv(cfg *Config) {
	cfg.Database.ConnectionString = fmt.Sprintf("host=%s port=5432 user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	cfg.Environment = os.Getenv("APP_ENV")

	cfg.SMPT.Host = os.Getenv("SMPT_HOST")
	cfg.SMPT.User = os.Getenv("SMPT_USER")
	cfg.SMPT.Password = os.Getenv("SMPT_PASS")
	cfg.SMPT.Port, _ = strconv.Atoi(os.Getenv("SMPT_PORT"))

}

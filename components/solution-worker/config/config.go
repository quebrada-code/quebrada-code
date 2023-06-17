package config

import (
	"github.com/spf13/viper"
	"os"
)

type (
	Config struct {
		Environment   string
		MessageBroker MessageBrokerConfig
	}

	MessageBrokerConfig struct {
		BootstrapServers string `json:"bootstrap.servers"`
		SecurityProtocol string `json:"security.protocol"`
		GroupId          string `json:"group.id"`
		SaslMechanism    string `json:"sasl.mechanism"`
		SaslUsername     string `json:"sasl.username"`
		SaslPassword     string `json:"sasl.password"`
		SslCaLocation    string `json:"ssl.ca.location"`
	}
)

// Init populates Config struct with values from config file
// located at filepath and environment variables.
func Init() (*Config, error) {

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func unmarshal(cfg *Config) error {

	return viper.UnmarshalKey("message_broker", &cfg.MessageBroker)
}

func setFromEnv(cfg *Config) {

	cfg.Environment = os.Getenv("APP_ENV")
	cfg.MessageBroker.BootstrapServers = os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	cfg.MessageBroker.GroupId = os.Getenv("KAFKA_GROUP_ID")
	cfg.MessageBroker.SaslMechanism = os.Getenv("KAFKA_SASL_MECHANISM")
	cfg.MessageBroker.SaslPassword = os.Getenv("KAFKA_SASL_PASSWORD")
	cfg.MessageBroker.SaslUsername = os.Getenv("KAFKA_SASL_USERNAME")
	cfg.MessageBroker.SecurityProtocol = os.Getenv("KAFKA_SECURITY_PROTOCOL")
	cfg.MessageBroker.SslCaLocation = os.Getenv("KAFKA_SSL_CA_LOCATION")

}

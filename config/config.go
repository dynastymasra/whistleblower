package config

import (
	"log"

	"github.com/dynastymasra/whistleblower/infrastructure/provider"

	"github.com/dynastymasra/cookbook"
	"github.com/spf13/viper"
)

type Config struct {
	serverPort string
	logger     LoggerConfig
	postgres   provider.Postgres
}

var config *Config

func Load() {
	viper.SetDefault(envServerPort, "8080")
	viper.SetDefault(envLogLevel, "debug")
	viper.SetDefault(envLogFormat, "text")

	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AddConfigPath("../../../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	config = &Config{
		serverPort: getString(envServerPort),
		logger: LoggerConfig{
			format: getString(envLogFormat),
			level:  getString(envLogLevel),
		},
		postgres: provider.Postgres{
			DatabaseName: getString(envPostgresName),
			Address:      getString(envPostgresAddress),
			Username:     getString(envPostgresUsername),
			Password:     getString(envPostgresPassword),
			MaxIdleConn:  getInt(envPostgresMaxIdleConn),
			MaxOpenConn:  getInt(envPostgresMaxOpenConn),
			LogEnabled:   getBool(envPostgresLogEnable),
		},
	}
}

func ServerPort() string {
	return config.serverPort
}

func Logger() LoggerConfig {
	return config.logger
}

func Postgres() provider.Postgres {
	return config.postgres
}

func getString(key string) string {
	value, err := cookbook.StringEnv(key)
	if err != nil {
		log.Fatalf("%v env key is not set", key)
	}
	return value
}

func getInt(key string) int {
	value, err := cookbook.IntEnv(key)
	if err != nil {
		log.Fatalf("%v env key is not set", key)
	}
	return value
}

func getBool(key string) bool {
	value, err := cookbook.BoolEnv(key)
	if err != nil {
		log.Fatalf("%v env key is not set", key)
	}
	return value
}

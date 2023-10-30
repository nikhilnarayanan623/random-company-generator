package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	CompanyServiceHost string `mapstructure:"COMPANY_SERVICE_HOST"`
	CompanyServicePort string `mapstructure:"COMPANY_SERVICE_PORT"`

	DBProtocol string `mapstructure:"DB_PROTOCOL"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBOptions  string `mapstructure:"DB_OPTIONS"`
}

var envs = []string{
	"COMPANY_SERVICE_HOST", "COMPANY_SERVICE_PORT",
	"DB_PROTOCOL", "DB_HOST", "DB_NAME", "DB_USER", "DB_PASSWORD", "DB_OPTIONS",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}

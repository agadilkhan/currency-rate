package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HttpServer `yaml:"HttpServer"`
	Database   `yaml:"Database"`
	Transport  `yaml:"Transport"`
	Job        `yaml:"Job"`
}

type HttpServer struct {
	Port            string        `yaml:"Port"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
}

type Database struct {
	Url string `yaml:"Url"`
}

type Transport struct {
	Host string `yaml:"Host"`
}

type Job struct {
	UpdateInterval    time.Duration `yaml:"UpdateInterval"`
	CurrencyVariation string        `yaml:"CurrencyVariation"`
}

func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to ReadInConfig err: %v", err)
	}

	var cfg Config

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to Unmarshal err: %v", err)
	}

	return &cfg, nil
}

package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	HttpServer `yaml:"HttpServer"`
	Database   `yaml:"Database"`
}

type HttpServer struct {
	Host            string        `yaml:"Host"`
	ShutdownTimeout time.Duration `yaml:"ShutdownTimeout"`
}

type Database struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `env:"DATABASE_PASSWORD"`
	Name     string `yaml:"Name"`
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

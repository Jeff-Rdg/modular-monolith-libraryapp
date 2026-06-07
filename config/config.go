package config

import "github.com/spf13/viper"

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func Load() (*Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	viper.SetDefault("SERVER_PORT", "3000")
	viper.SetDefault("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/library?sslmode=disable")

	cfg := &Config{
		DatabaseURL: viper.GetString("DATABASE_URL"),
		ServerPort:  viper.GetString("SERVER_PORT"),
	}

	return cfg, nil
}

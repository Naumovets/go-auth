package auth

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	RefreshTokenSecretKey string `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenSecretKey  string `mapstructure:"ACCESS_TOKEN_SECRET"`
}

func NewConfig(path string) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("cannot read config from %s: %w", path, err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("cannot unmarshal config: %w", err)
	}
	return cfg, nil
}

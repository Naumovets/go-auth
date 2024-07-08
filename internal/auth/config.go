package auth

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	refreshTokenSecretKey string `mapstructure:"REFRESH_TOKEN_SECRET"`
	accessTokenSecretKey  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	GrpcAddress           string `mapstructure:"GRPC_ADDRESS"`
	HttpAddress           string `mapstructure:"HTTP_ADDRESS"`
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

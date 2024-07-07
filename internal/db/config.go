package db

type Config struct {
	DB_NAME  string `mapstructure:"PG_NAME"`
	USER     string `mapstructure:"PG_USER"`
	PASSWORD string `mapstructure:"PG_PASSWORD"`
	DB_HOST  string `mapstructure:"PG_HOST"`
	DB_PORT  string `mapstructure:"PG_PORT"`
}

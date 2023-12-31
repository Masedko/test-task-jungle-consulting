package configuration

import (
	"github.com/spf13/viper"
	"time"
)

type EnvConfigModel struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	SSLMode        string `mapstructure:"SSL_MODE"`

	AllowOrigins string `mapstructure:"ALLOW_ORIGINS"`
	AllowHeaders string `mapstructure:"ALLOW_HEADERS"`

	APIPort string `mapstructure:"CONTAINER_API_PORT"`

	JwtSecret    string        `mapstructure:"JWT_SECRET_KEY"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_SECRET_KEY_EXPIRES_IN"`
}

var EnvConfig EnvConfigModel

func LoadConfig(filePath string) (err error) {
	viper.SetConfigType("env")
	viper.SetConfigFile(filePath)

	viper.AutomaticEnv()

	if viper.ReadInConfig() != nil {
		return
	}

	return viper.Unmarshal(&EnvConfig)
}

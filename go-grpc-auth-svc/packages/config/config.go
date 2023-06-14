package config

import "github.com/spf13/viper"

type Config struct {
	Port                 string `mapstructure:"PORT"`
	DBUrl                string `mapstructure:"DB_URL"`
	DBHost               string `mapstructure:"DB_HOST"`
	DBPort               string `mapstructure:"DB_PORT"`
	DBName               string `mapstructure:"DB_NAME"`
	DBUser               string `mapstructure:"DB_USER"`
	DBPassword           string `mapstructure:"DB_PASSWORD"`
	DBTZ                 string `mapstructure:"DB_TZ"`
	JWTSecretKey         string `mapstructure:"AUTH_SECRET_KEY"`
	AccessTokenMinuteTTL int64  `mapstructure:"AUTH_ACCESS_TOKEN_MINUTE_TTL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./packages/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}

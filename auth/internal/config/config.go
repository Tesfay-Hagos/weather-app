package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	DATABASEURL   string `mapstructure:"DATABASE_URL"`
	JWTSecretKey  string `mapstructure:"JWT_SECRET_KEY"`
	MigrationPath string `mapstructure:"MIGRATION_PATH"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
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

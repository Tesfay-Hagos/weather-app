package config

import "github.com/spf13/viper"

// Config holds the application configuration.
type Config struct {
	Port          string `mapstructure:"PORT"`
	DATABASEURL   string `mapstructure:"DATABASE_URL"`
	WeatherURL    string `mapstructure:"WEATHER_URL"`
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
	DbName        string `mapstructure:"DB_NAME"`
}

// LoadConfig loads the application configuration from the environment and the specified file.
func LoadConfig(path string) (config Config, err error) {
	// Set the path to look for the configurations file
	viper.AddConfigPath(path)
	// Set the name of the configuration file
	viper.SetConfigName("dev")
	// Set the configuration file type
	viper.SetConfigType("env")

	// Automatically override values from the file with the values from the environment
	viper.AutomaticEnv()

	// Read the configuration file
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// Unmarshal the configuration into the Config struct
	err = viper.Unmarshal(&config)

	return
}

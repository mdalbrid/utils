package config

import (
	"github.com/mdalbrid/utils/logger"
	"os"
	"path"

	"github.com/spf13/viper"
)

var (
	config *viper.Viper
)

// GetViper - *viper.Viper - Возвращает объект Viper
func GetViper() *viper.Viper {
	if nil == config {
		setupViper()
	}
	return config
}

// setupViper - читает данные из *.ini файла и сохраняет их в переменных окружения
func setupViper() {
	config = viper.New()

	configPath := os.Getenv("AWS_CONFIG_FILE")
	if configPath != "" {
		config.SetConfigType(configPath)
	} else {
		config.SetConfigName("config")
		config.SetConfigType("ini")
		config.AddConfigPath("/etc/aws")
		config.AddConfigPath("/opt/aws")
		config.AddConfigPath("/aws")
		config.AddConfigPath(".")
		config.AddConfigPath(path.Join(os.Getenv("HOME"), ".aws"))
	}

	err := config.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
		panic("Error read config")
	}

}

package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Host     string
		Port     string
		DBName   string
		User     string
		Password string
	}
}

var c *viper.Viper

// Init initializes config
func Init(env string) {
	c = viper.New()
	c.SetConfigType("yaml")
	c.SetConfigName(env)
	c.AddConfigPath("src/config/environments/")
	if err := c.ReadInConfig(); err != nil {
		fmt.Println("config file read error")
		fmt.Println(err)
		os.Exit(1)
	}
}

// GetConfig returns config
func GetConfig() *viper.Viper {
	return c
}

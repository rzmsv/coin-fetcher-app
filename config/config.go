package config

import (
	dotenv "github.com/joho/godotenv"
	"log"
)

type AppConfig struct{}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}

func (c *AppConfig) Configs(configKey string) string {
	envs, err := dotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error message: %s\n", err)
	}
	if len(envs) == 0 {
		log.Fatal(".env file is empty!")
	}
	configs := make(map[string]string)
	for key, value := range envs {
		configs[key] = value
	}
	if configs[configKey] == "" {
		log.Fatalf("%s is not exists in .env file!", configKey)
	}
	return configs[configKey]

}

package configs

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Host string `yaml:"host" env:"host" env-default:"localhost"`
		Port int    `yaml:"port" env:"port" env-default:"8080"`
	} `yaml:"server"`

	DB struct {
		Host    string `yaml:"host" env:"host" env-default:"localhost"`
		Port    int    `yaml:"port" env:"port" env-default:"5432"`
		SSLMode string `yaml:"sslmode" env:"sslmode" env-default:"disable"`
	} `yaml:"db"`
}

func GetConfig() Config {
	var cfg Config
	_ = godotenv.Load(".env")
	err := cleanenv.ReadConfig("configs/config.yml", &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func GetDbUrl() string {
	_ = godotenv.Load(".env")
	return os.Getenv("DB_URL")

}

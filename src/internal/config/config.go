package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// const CONFIG_PATH = "./config/local.yaml"
const CONFIG_PATH = "/usr/local/etc/person-service/config.yaml"

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	HTTPServer `yaml:"http_server"`
	DB         `yaml:"db"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhsot:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func Load() Config {
	if _, err := os.Stat(CONFIG_PATH); os.IsNotExist(err) {
		log.Fatalf("config file doesnot exist: %s", CONFIG_PATH)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(CONFIG_PATH, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return cfg
}

package config

import (
	"github.com/pkg/errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

var config *Config

// Config .
type Config struct {
	Port  string `envconfig:"PORT" yaml:"port"`
	Debug bool   `envconfig:"DEBUG" yaml:"debug"`

	MySQL struct {
		Host   string `envconfig:"DB_HOST"`
		Port   string `envconfig:"DB_PORT"`
		DBName string `envconfig:"DB_NAME"`
		User   string `envconfig:"DB_USER"`
		Pass   string `envconfig:"DB_PASSWORD"`
	} `yaml:"mysql"`
}

func init() {
	config = &Config{}

	// read from config file
	f, err := os.Open("config.yml")
	if err != nil {
		log.Println("Failed to read config", err)
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(config)
	if err != nil {
		log.Println("Failed to decode config yml", err)
	}

	// read from env
	err = godotenv.Load()
	if err != nil {
		log.Println("Failed to load godotenv", err)
	}

	err = envconfig.Process("", config)
	if err != nil {
		err = errors.Wrap(err, "Failed to decode config env")
		log.Println(err)
	}

	// default value
	if len(config.Port) == 0 {
		config.Port = "3000"
	}
}

// GetConfig .
func GetConfig() *Config {
	return config
}

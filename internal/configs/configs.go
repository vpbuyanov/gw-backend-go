package configs

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server   Server   `yaml:"server"`
	Kafka    Kafka    `yaml:"kafka"`
	Logger   Logger   `yaml:"logger"`
	JWT      JWT      `yaml:"jwt"`
	Postgres Postgres `yaml:"postgres"`
}

func New() *Config {
	return &Config{
		Server:   Server{},
		Postgres: Postgres{},
		Kafka:    Kafka{},
		JWT:      JWT{},
		Logger:   Logger{},
	}
}

func MustConfig(p *string) *Config {
	var path string
	if p == nil {
		path = fetchConfigPath()
	}

	if path == "" {
		path = "./config.yaml"
	}

	if _, ok := os.Stat(path); os.IsNotExist(ok) {
		panic("Config file does not exist: " + path)
	}

	cfg := New()

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return cfg
}

func fetchConfigPath() string {
	var res string

	// --config="path/to/config.yaml"
	flag.StringVar(&res, "config", "", "path to config")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

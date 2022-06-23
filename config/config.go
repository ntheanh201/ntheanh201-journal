package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Config struct {
	App  `yaml:"app"`
	Http `yaml:"http"`
}

type Http struct {
	Port string `yaml:"port"`
}

type Log struct {
	Level string `yaml:"log_level"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}

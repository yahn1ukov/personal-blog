package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP struct {
		Port int `yaml:"port"`
	} `yaml:"http"`

	Database struct {
		Mongo struct {
			URL string `yaml:"url"`
		} `yaml:"mongo"`
	} `yaml:"database"`
}

func New(path string) (*Config, error) {
	var cfg Config

	if path != "" {
		if err := cleanenv.ReadConfig(path, &cfg); err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

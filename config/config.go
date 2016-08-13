package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Github struct {
		Tokens []string
	}
}

func Load() Config {
	file, err := ioutil.ReadFile("./config/secrets.yml")

	if err != nil {
		log.Fatal("failed to read config file")
	}

	config := &Config{}
	_ = yaml.Unmarshal([]byte(file), config)
	return *config
}

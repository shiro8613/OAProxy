package modules

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`

	Session []interface{} `yaml:"session"`

	Prefix string `yaml:"prefix"`

	Oauth2 []interface{} `yaml:"oauth2"`

	Server []interface{} `yaml:"server"`
}

func GetConfig() config{
	config := config{}
	b, err := os.ReadFile("./config.yml")
	if err != nil {
		fmt.Println(err)
	}

	yaml.Unmarshal(b, &config)
	return config
}
package modules

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`

	Session map[string]interface{} `yaml:"session"`
	Redis map[string]interface{} `yaml:"redis"`

	Prefix string `yaml:"prefix"`

	Oauth2 map[string]interface{} `yaml:"oauth2"`

	Server map[interface{}]interface{} `yaml:"server"`
}

func ConfigLoad() config{
	config := config{}
	b, err := os.ReadFile("./config.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	yaml.Unmarshal(b, &config)
	return config

}

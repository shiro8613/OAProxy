package modules

import (
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`

	//Https map[string]interface{} `yaml:"https"` <- https用

	Session map[string]interface{} `yaml:"session"`
	Redis map[string]interface{} `yaml:"redis"`

	Prefix string `yaml:"prefix"`

	Oauth2 map[string]interface{} `yaml:"oauth2"`

	Server map[interface{}]interface{} `yaml:"server"`
}

var configer []config

func ConfigLoad() {
	config := config{}
	path := GetFlag().Config_path
	b, err := os.ReadFile(path)
	if err != nil {
		Logger("error", err.Error())
	}

	yaml.Unmarshal(b, &config)
	
	configer = append(configer, config)

}

func GetConfig() config {
	for _, v := range configer {
		return v
	}
	return configer[0]
}

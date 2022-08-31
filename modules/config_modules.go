package modules

import (
	"os"
	"regexp"
	"gopkg.in/yaml.v2"
)

type https_conf struct {
	Enable	bool	`map:"enable"`
	Port	int		`map:"port"`
	Cert	string	`map:"cert"`
	Key		string	`map:"key"`
}

type config struct {
	Host	string	`yaml:"host"`
	Port	int		`yaml:"port"`
	Domain	string	`yaml:"domain"`

	Https	https_conf	`yaml:"https"`

	Session map[string]interface{}		`yaml:"session"`
	Redis	map[string]interface{}		`yaml:"redis"`

	Prefix	string `yaml:"prefix"`

	Oauth2	map[string]interface{}		`yaml:"oauth2"`

	Server	map[interface{}]interface{} `yaml:"server"`
}

var configer []config

func ConfigLoad() {
	config := config{}
	rg := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	path := GetFlag().Config_path
	b, err := os.ReadFile(path)
	if err != nil {
		Logger("error", err.Error())
	}

	yaml.Unmarshal(b, &config)

	if rg.MatchString(config.Domain) {
		Logger("error", "Enter your domain in this format example.com or xxx.example.com")
	}
	
	configer = append(configer, config)

}

func GetConfig() config {
	for _, v := range configer {
		return v
	}
	return configer[0]
}

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

type session_conf struct {
	MaxAge		int		`map:"maxAge"`
	Secure		bool	`map:"secure"`
	HttpOnly	bool	`map:"httpOnly"`
	Mode		string	`map:"mode"`
}

type redis_conf struct {
	Host		string	`map:"host"`
	Port		int		`map:"port"`
	Password	string	`map:"password"`
}

type oauth2_conf struct	 {
	Client_id		int64				`map:"client_id"`
	Client_secret	string				`map:"client_secret"`
	Callback		string				`map:"callback"`
	Guild_id		int64				`map:"guild_id"`
	Roles			map[string]int64	`map:"roles"`
}

type server_conf struct {
	Location		string		`map:"location"`
	Address			string		`map:"address"`
	Privart			bool		`map:"privart"`
	Access_roles	[]string	`map:"access_roles"`
}

type config struct {
	Host	string					`yaml:"host"`
	Port	int						`yaml:"port"`
	Domain	string					`yaml:"domain"`
	Https	https_conf				`yaml:"https"`
	Session session_conf			`yaml:"session"`
	Redis	redis_conf				`yaml:"redis"`
	Prefix	string 					`yaml:"prefix"`
	Oauth2	oauth2_conf				`yaml:"oauth2"`
	Server	map[string]server_conf	`yaml:"server"`
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

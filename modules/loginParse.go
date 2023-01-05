package modules

import (
	"encoding/json"
)


type user struct {
	Username string			`map:"username"`
	Discriminator string	`map:"discriminator"`
	Id string				`map:"id"`
}

type LoginUserConf struct {
	Nick	string			`json:"nick"`
	User	user 			`json:"user"`
	Roles 	[]string	`json:"roles"`
}

func LoginUserParse(data string) LoginUserConf {
	Luc := LoginUserConf{}
	err := json.Unmarshal([]byte(data), &Luc)
	if err != nil {
		Logger("error", err.Error())
	}

	return Luc
}

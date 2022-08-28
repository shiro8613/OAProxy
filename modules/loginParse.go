package modules

import (
	"encoding/json"
	"fmt"
)

type LoginUserConf struct {
	Nick string `json:"nick"`
	User interface{} `json:"user"`
	Roles []interface{} `json:"roles"`
}

func LoginUserParse(data string) LoginUserConf {
	Luc := LoginUserConf{}
	err := json.Unmarshal([]byte(data), &Luc)
	if err != nil {
		Logger("error", err.Error())
	}

	fmt.Println(Luc)
	return Luc
}

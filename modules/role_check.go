package modules

import (
	"encoding/json"
)

func CheckRole(inter []string) string{
	roles := GetConfig().Oauth2.Roles
	var myroles []string

	for _, v := range inter {
		for key, value := range roles {
			if value == v {
				myroles = append(myroles, key)
			}
		}
	}
	
	e, err := json.Marshal(myroles)
    if err != nil {
        Logger("error", err.Error())
    }
    return string(e)
}

func RoleTest(inter []string, data string) bool{
	var rolejs []string
	err := json.Unmarshal([]byte(data), &rolejs)
	if err != nil {
		Logger("error", err.Error())
	}

	for _, v := range rolejs {
		for _, val := range inter {
			if val == v {
				return true
			}
		}
	}
	return false
}
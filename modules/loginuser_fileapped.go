package modules

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type UserLogJson struct {
	Time	string 	`json:"time"`
	Name	string 	`json:"name"`
	Nick	string 	`json:"nick"`
	Id		string 	`json:"id"`
	Ip		string 	`json:"ip"`
}

func remove(s []UserLogJson, i int) []UserLogJson {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func Exists(name string) bool {
    _, err := os.Stat(name)
    return os.IsNotExist(err)
}

func AppendUser(ip string,	id string, name string, disc string, nick string ) {
	path := GetFlag().UserList_path
	time := time.Now().Local().Format(time.RFC1123)

	if Exists(path) {
		os.Create(path)
	}

	userJsonbs := []UserLogJson{}

	rfile, err := os.ReadFile(path)
	if err != nil {
		Logger("error", err.Error())
	}

	json.Unmarshal(rfile, &userJsonbs)

	os.Remove(path)
	
	for i, userJsonb := range userJsonbs {
		if userJsonb.Id == id {
			userJsonbs = remove(userJsonbs, i)				
		}
	}

	userLog := UserLogJson {
		Time: time,
		Name: fmt.Sprintf("%s#%s", name, disc),
		Nick: nick,
		Id: id,
		Ip: ip,
	}

	userJsonbs = append(userJsonbs, userLog)

	b, err := json.Marshal(userJsonbs)
	if err != nil {
		Logger("error",	err.Error())
	}

	file, err := os.Create(path)
	if err != nil {
		Logger("error", err.Error())
	}

	defer file.Close()

	file.Write(b)
}
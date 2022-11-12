package modules

import (
	"bufio"
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

func AppendUser(ip string,	id string, name string, disc string, nick string ) {
	path := GetFlag().UserList_path
	time := time.Now().Local().Format(time.RFC1123)

	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
    if err != nil {
		Logger("error", err.Error())
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jdata := &UserLogJson{}
		err = json.Unmarshal(scanner.Bytes(), jdata)
		if err != nil {
			Logger("error", err.Error())
		}

		if jdata.Id == id {
			return 
			
		} else {
			userlog := UserLogJson{
				Time: time,
				Name: fmt.Sprintf("%s#%s",name ,disc),
				Nick: nick,
				Id: id,
				Ip: ip,
			}

			json_, err := json.Marshal(userlog)
			if err != nil {
				Logger("error", err.Error())
			}

			fmt.Fprintln(file, string(json_))

		}
	}
}
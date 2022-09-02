package modules

import (
	"fmt"
	"os"
	"time"
)

func AppendUser(ip string,	id string, name string, disc string, nick string ) {
	path := GetFlag().UserList_path
	time := time.Now().Local().Format(time.RFC1123)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
		Logger("error", err.Error())
    }
    defer file.Close()
	d := fmt.Sprintf("ip:%s, id:%s, Name:%s, Nick:%s",ip, id, fmt.Sprintf("%s#%s",name ,disc),nick)
	fmt.Fprintln(file, fmt.Sprintf("%s => %s", time, d))

}
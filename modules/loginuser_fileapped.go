package modules

import (
	"time"
	"os"
	"fmt"
)

func AppendUser(content string) {
	path := GetFlag().UserList_path
	time := time.Now().Local().Format(time.RFC1123)

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
		Logger("error", err.Error())
    }
    defer file.Close()
    fmt.Fprintln(file, fmt.Sprintf("%s => %s", time, content))

}
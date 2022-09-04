package modules

import (
	"fmt"
	"os"
	"time"
)

func LogFolderCreate() {
	path := GetFlag().Logsdir_path
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {	
		Logger("warn", "NotDirectory..., DirectoryCreated!")
		if err := os.Mkdir(path, 0755); err != nil {
			Logger("error", err.Error())
		}
	}
}

func LogWrite(content string) {
	path := GetFlag().Logsdir_path
	time := time.Now().Local().Format("2006-01-02")

	file, err := os.OpenFile(fmt.Sprintf("./%s/%s.log",path, time), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
    if err != nil {
		Logger("error", err.Error())
    }
    defer file.Close()
    fmt.Fprintln(file, content)
}
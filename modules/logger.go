package modules

import (
	"fmt"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/labstack/echo/v4"
)

func EchoLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		res := c.Response()
		req := c.Request()

		status := res.Status

		//810.114.514.194  [YJ/Yajusenpai#1919] - - [Mon, 8 Oct 1919 11:45:14 JST] Method URL Proto statusCode
		time := time.Now().Local().Format(time.RFC1123)
		name := ReadSession(c, "name")
		content := fmt.Sprintf("%s [%s] - - [%s] %s %s %s %d",c.RealIP(), name, time, req.Method, req.RequestURI, req.Proto, status) 
		Logger("info", content)
		return next(c)
	}
}

func Logger(mode string, content string) {
	switch mode {
	case "error":
		go LogWrite(content)
		color.Print("<red>[ERROR]</>"+content+"\n")
		os.Exit(0)	

	case "info":
		go LogWrite(content)
		color.Print("<green>[INFO]</>"+content+"\n")
		return
	
	case "warn":
		color.Print("<yellow>[WARN]</>"+content+"\n")
	}
}

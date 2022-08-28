package modules

import (
	"os"

	"github.com/gookit/color"
	"github.com/labstack/echo/v4"
)

func EchoLogger(next echo.HandlerFunc) echo.HandlerFunc{
	return func (c echo.Context) error {
		content := ""
		Logger("info", content)
		return next(c)
	}
}

func Logger(mode string, content string) {
	switch mode {
	case "error":
		color.Print("<red>[ERROR]</>"+content+"\n")
		os.Exit(0)	

	case "info":
		color.Print("<orange>[INFO]</>"+content+"\n")
		return
	}
}

package main

import (
	"fmt"

	"github.com/flan0910/OAProxy/handler"
	"github.com/flan0910/OAProxy/modules"
	"github.com/flan0910/OAProxy/middler"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	modules.ParseFlag()
	modules.LogFolderCreate()
	modules.ConfigLoad()

	config := modules.GetConfig()

	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(middler.DomainCheck)
	e.Use(session.Middleware(modules.StoreCreate()))
	middler.MiddleProx(*e)
	e.Use(modules.EchoLogger)
	//e.Use(middleware.Logger())

	e.GET("/", handler.SlashAccess)
	e.GET(fmt.Sprintf("/%s/login", config.Prefix), handler.Login)
	e.GET(fmt.Sprintf("/%s/logout", config.Prefix), handler.Logout)
	e.GET(fmt.Sprintf("/%s/after", config.Prefix), handler.LoginAfter)

	if config.Https.Enable {
		if config.Domain == ""{
			modules.Logger("error", "You need a domain to use https")
		}
		modules.Logger("error", e.StartTLS(fmt.Sprintf("%s:%d", config.Domain, config.Https.Port), config.Https.Cert, config.Https.Key).Error())
	}else {
		modules.Logger("error", e.Start(fmt.Sprintf("%s:%d", config.Host, config.Port)).Error())
	}

}
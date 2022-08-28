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
	config := modules.ConfigLoad()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(modules.StoreCreate()))



	e.GET("/", handler.SlashAccess)
	e.GET(fmt.Sprintf("/%s/login", config.Prefix), handler.Login)
	e.GET(fmt.Sprintf("/%s/logout", config.Prefix), handler.Logout)
	e.GET(fmt.Sprintf("/%s/after", config.Prefix), handler.LoginAfter)
	e.GET(fmt.Sprintf("/%s/guild", config.Prefix), handler.LoginGuild)
	e.GET(fmt.Sprintf("/%s/user", config.Prefix), handler.LoginUser)

	middler.MiddleProx(*e)

	modules.Logger("error", e.Start(fmt.Sprintf("%s:%d", config.Host, config.Port)).Error())

}
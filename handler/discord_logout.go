package handler

import (

	"github.com/flan0910/OAProxy/modules"

	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	seslogin := modules.ReadSession(c, "login")
	sesguild := modules.ReadSession(c, "guild")

	if seslogin == "true" || sesguild == "true" {
		modules.DeleteSession(c)
		return modules.AfterLogout(c)
	} else {
		return modules.GuildErrorPages(c)
	}
}

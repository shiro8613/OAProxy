package handler

import (
	"fmt"
	"net/http"

	"github.com/flan0910/OAProxy/modules"

	"github.com/labstack/echo/v4"
)

func SlashAccess(c echo.Context) error {
	config := modules.GetConfig()
	seslogin := modules.ReadSession(c, "login")
	sesguild := modules.ReadSession(c, "guild")
	if seslogin == "true" {
		if sesguild == "true" {
			return c.Redirect(http.StatusOK, config.Redirectafter)
		}else {
			return modules.GuildErrorPages(c)
		}
	}else {
		modules.WriteSession(c, "urled", c.Request().URL.Path)
		return c.Redirect(http.StatusFound, fmt.Sprintf("/%s/login", config.Prefix))	
	}
}

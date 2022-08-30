package handler

import (
	"fmt"
	"net/http"

	"github.com/flan0910/OAProxy/modules"

	"github.com/labstack/echo/v4"
)

func SlashAccess(c echo.Context) error {
	prefix := modules.GetConfig().Prefix
	seslogin := modules.ReadSession(c, "login")
	sesguild := modules.ReadSession(c, "guild")
	if seslogin == "true" {
		if sesguild == "true" {
			return c.String(http.StatusOK, "LoginOK")
		}else {
			return c.String(http.StatusForbidden, "JoinGuild")
		}
	}else {
		modules.WriteSession(c, "urled", c.Request().URL.Path)
		return c.Redirect(http.StatusFound, fmt.Sprintf("/%s/login", prefix))	
	}}

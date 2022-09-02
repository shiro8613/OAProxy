package middler

import (
	"net/http"
	"net/url"
	"fmt"

	"github.com/flan0910/OAProxy/modules"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MiddleProx(e echo.Echo) {
	config := modules.GetConfig()

	for _, serverMap := range config.Server {

		g := e.Group(serverMap.Location)
		urls, err := url.Parse(serverMap.Address)
		if err != nil {
			modules.Logger("error", err.Error())
		}

		target := []*middleware.ProxyTarget{
			{
				URL: urls,
			},
		}

		g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error { //session
				seslogin := modules.ReadSession(c, "login")
				sesguild := modules.ReadSession(c, "guild")
				
				if seslogin == "true" {
					if sesguild == "true" {
						return next(c)
					}else {
						return c.String(http.StatusForbidden, "NoGuilds")
					}
				}else {
					modules.WriteSession(c, "urled", c.Request().URL.Path)
					return c.Redirect(http.StatusFound, fmt.Sprintf("/%s/login",config.Prefix))
				}
			}

		}, func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error { //privert
				if serverMap.Privart == true {
					if modules.RoleTest(serverMap.Access_roles, modules.ReadSession(c, "role")) {
						return next(c)
					}else {
						return c.String(http.StatusForbidden, "NoRoles")
					}
				}
				return next(c)
			}
			
		}, middleware.Proxy(middleware.NewRandomBalancer(target)))
	}
}


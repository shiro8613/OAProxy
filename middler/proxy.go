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

	for key, serv := range config.Server {
		serverMap := serv
		keysd := key

		if serverMap.Location == "/" {
			modules.Logger("error", "/ is not supported! If you want to set /, enable slash_access of external_host_pages and set the address")
		}

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

		for _, v := range config.NeedLogin {
			if keysd == v {
				g.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
					return func(c echo.Context) error { //session
						seslogin := modules.ReadSession(c, "login")
						sesguild := modules.ReadSession(c, "guild")
						
						if seslogin == "true" {
							if sesguild == "true" {
								return next(c)
							}else {
								return modules.GuildErrorPages(c)
							}
						}else {
							modules.WriteSession(c, "urled", c.Request().URL.Path)
							return c.Redirect(http.StatusFound, fmt.Sprintf("/%s/login",config.Prefix))
						}
					}

				}, func(next echo.HandlerFunc) echo.HandlerFunc {
					return func(c echo.Context) error { //private
						if serverMap.Private {
							if modules.RoleTest(serverMap.Access_roles, modules.ReadSession(c, "role")) {
								return next(c)
							}else {
								return modules.PrivateErrorPages(c)
							}
						} else {
							return next(c)
						}
					}
					
				}, middleware.Proxy(middleware.NewRandomBalancer(target)))
			}
		}
		g.Use(middleware.Proxy(middleware.NewRandomBalancer(target)))
	}
}


package middler

import (
	"net/http"
	"net/url"

	"github.com/flan0910/OAProxy/handler"
	"github.com/flan0910/OAProxy/modules"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ExternalSlash(e echo.Echo) {
	config := modules.GetConfig().ExternalHost["slash_access"]
	
	if config.Enable {
		if config.Redirect { 
			e.GET("/", func(c echo.Context) error {
				return c.Redirect(http.StatusFound, config.Address)
			})
		} else {
			urls, err := url.Parse(config.Address)
			if err != nil {
				modules.Logger("error", err.Error())
			}
	
			target := []*middleware.ProxyTarget{
				{
					URL: urls,
				},
			}

			e.Group("/", middleware.Proxy(middleware.NewRandomBalancer(target)))
		}
	} else {
		e.GET("/", handler.SlashAccess)
	}
}
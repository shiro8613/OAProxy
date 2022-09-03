package modules

import (
	"net/http"
	"net/http/httputil"

	"github.com/labstack/echo/v4"
)

func AfterLogout(c echo.Context) error {
	config := GetConfig().ExternalHost["after_logout"]
	
	if config.Enable {
		if config.Redirect {
			return c.Redirect(http.StatusFound, config.Address)
		} else {
				c.SetHandler(echo.WrapHandler(&httputil.ReverseProxy{Director: func(request *http.Request) {
					request.URL.RawQuery = ""
					request.URL.Scheme = "http"
					request.URL.Host = parse(config.Address)
					request.URL.Path = "/"
				}}))
				
				return c.Handler()(c)
		}
	} else {
		return c.String(http.StatusOK, "LogoutComplited! You can close the page")
	}
}

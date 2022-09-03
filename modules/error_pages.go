package modules

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
)

func parse(urls string) string {
	h, _ := url.Parse(urls)
	return h.Host
}

func GuildErrorPages(c echo.Context) error{
	config := GetConfig().ExternalHost["guild_error"]
	
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
		return c.String(http.StatusForbidden, "JoinGuild")
	}
}

func LoginErrorPages(c echo.Context) error{
	config := GetConfig().ExternalHost["login_error"]

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
		return c.String(http.StatusForbidden, "NeedDiscordLogin")
	}
}

func PrivartErrorPages(c echo.Context) error{
	config := GetConfig().ExternalHost["privart_error"]

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
		return c.String(http.StatusForbidden, "Permission denied")
	}
}
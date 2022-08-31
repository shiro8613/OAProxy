package middler

import (
	"fmt"
	"net/http"

	"github.com/flan0910/OAProxy/modules"
	"github.com/labstack/echo/v4"
)

func DomainCheck(next echo.HandlerFunc) echo.HandlerFunc {
	config := modules.GetConfig()
	return func (c echo.Context) error {
		if config.Domain == "" {
			return next(c)
		}
		if config.Port == 80 || config.Port == 443 {
			if c.Request().Host == config.Domain {
				return next(c)
			}
		} else {
			if c.Request().Host == fmt.Sprintf("%s:%d", config.Domain,config.Port) {
				return next(c)
			}
		}
		return c.String(http.StatusNotFound, "")
	}
}
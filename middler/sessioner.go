package middler

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func ReadSession(c echo.Context, key string) string{
	sess, _ := session.Get("session", c)
	ret := sess.Values[key]
	if ret != nil {
		return ret.(string)
	}
	return "false"
}

func WriteSession(c echo.Context, key string, value string) {
	sess, _ := session.Get("session", c)
	sess.Values[key] = value
	sess.Save(c.Request(), c.Response())
}

func DeleteSession(c echo.Context) {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
}

func FalseToSlash(path string) string {
	if path == "false" {
		return "/"
	}
	return path
}
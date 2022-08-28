package handler

import (
	"net/http"
	
	"github.com/flan0910/OAProxy/middler"

	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	middler.DeleteSession(c)
	return c.String(http.StatusOK, "ログアウトが完了しました。ページを閉じてください。")
}
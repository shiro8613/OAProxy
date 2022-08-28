package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/flan0910/OAProxy/middler"
	"github.com/flan0910/OAProxy/modules"

	"github.com/labstack/echo/v4"
)


func Login(c echo.Context) error {
	config := modules.ConfigLoad().Oauth2
	return c.Redirect(http.StatusFound, fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=identify guilds guilds.members.read", config["client_id"].(string),config["callback"].(string)))
}

func LoginAfter(c echo.Context) error {
	r := c.Request()
	config := modules.ConfigLoad()
	code := r.URL.Query().Get("code")

	pData := url.Values{}
	pData.Add("client_id", (config.Oauth2)["client_id"].(string))
	pData.Add("client_secret", (config.Oauth2)["client_secret"].(string))
	pData.Add("grant_type", "authorization_code")
	pData.Add("code", code)
	pData.Add("redirect_uri", (config.Oauth2)["callback"].(string))
	res := modules.XPoster("https://discordapp.com/api/oauth2/token", pData)
	ddata := modules.Decoder(res)
	token := ddata.(map[string]interface{})["access_token"].(string)
	middler.WriteSession(c, "login", "true")
	return c.Redirect(http.StatusFound, fmt.Sprintf("/%s/guild?token=%s", config.Prefix, token))
}

func LoginGuild(c echo.Context) error {
	r := c.Request()
	config := modules.ConfigLoad()
	token := r.URL.Query().Get("token")
	heads := http.Header{}
	heads.Add("Authorization", fmt.Sprintf("Bearer %s",token))
	res := modules.XGet("https://discordapp.com/api/v6/users/@me/guilds", heads)
	ddata := modules.Decoder_in(res)

	if modules.Filter(ddata, "id", (config.Oauth2)["guild_id"].(string)) {
		middler.WriteSession(c, "guild", "true")
		return c.Redirect(http.StatusFound, fmt.Sprintf("/%s/user?token=%s", config.Prefix, token))
	} else {
		return c.String(http.StatusForbidden, "JoinGuild")
	}

}

func LoginUser(c echo.Context) error {
	r := c.Request()
	config := modules.ConfigLoad()
	token := r.URL.Query().Get("token")
	heads := http.Header{}
	heads.Add("Authorization", "Bearer "+token)
	res := modules.XGet(fmt.Sprintf("https://discordapp.com/api/v6/users/@me/guilds/%s/member", (config.Oauth2)["guild_id"].(string)), heads)
	jdata := modules.LoginUserParse(res)
	middler.WriteSession(c, "name", fmt.Sprintf("%s/%s\n", jdata.Nick, jdata.User.(map[string]interface{})["username"].(string)))
	middler.WriteSession(c, "id", fmt.Sprintln(jdata.User.(map[string]interface{})["id"].(string)))
	middler.WriteSession(c, "role", modules.CheckRole(jdata.Roles))
	return c.Redirect(http.StatusFound, fmt.Sprintf("%v", middler.FalseToSlash(middler.ReadSession(c, "urled"))))
}
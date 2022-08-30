package handler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/flan0910/OAProxy/modules"

	"github.com/labstack/echo/v4"
)


func Login(c echo.Context) error {
	config := modules.GetConfig().Oauth2
	return c.Redirect(http.StatusFound, fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=identify guilds guilds.members.read", config["client_id"].(string),config["callback"].(string)))
}

func LoginAfter(c echo.Context) error {
	r := c.Request()
	config := modules.GetConfig()
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
	modules.WriteSession(c, "login", "true")

	heads := http.Header{}
	heads.Add("Authorization", fmt.Sprintf("Bearer %s",token))
	res1 := modules.XGet("https://discordapp.com/api/v6/users/@me/guilds", heads)
	ddata1 := modules.Decoder_in(res1)

	if modules.Filter(ddata1, "id", (config.Oauth2)["guild_id"].(string)) {
		modules.WriteSession(c, "guild", "true")
			
		heads := http.Header{}
		heads.Add("Authorization", "Bearer "+token)
		res := modules.XGet(fmt.Sprintf("https://discordapp.com/api/v6/users/@me/guilds/%s/member", (config.Oauth2)["guild_id"].(string)), heads)
		jdata := modules.LoginUserParse(res)
		ip := c.RealIP()
		name := jdata.User.(map[string]interface{})["username"].(string)
		id := jdata.User.(map[string]interface{})["id"].(string)
		disc := jdata.User.(map[string]interface{})["discriminator"].(string)
		modules.AppendUser(fmt.Sprintf("ip:%s, id:%s, Name:%s, Nick:%s",ip, id, fmt.Sprintf("%s#%s",name,disc), jdata.Nick))
		modules.WriteSession(c, "name", fmt.Sprintf("%s/%s#%s", jdata.Nick, name, disc))
		modules.WriteSession(c, "id", id)
		modules.WriteSession(c, "role", modules.CheckRole(jdata.Roles))
		return c.Redirect(http.StatusFound, fmt.Sprintf("%v", modules.FalseToSlash(modules.ReadSession(c, "urled"))))
	
	} else {
		return c.String(http.StatusForbidden, "JoinGuild")
	}
}
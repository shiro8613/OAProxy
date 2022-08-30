package modules

import (
	"fmt"

	"github.com/gorilla/sessions"
	session "github.com/ipfans/echo-session"
)

func StoreCreate() sessions.Store {
	config := GetConfig()
	key := Cryper()
	Age := config.Session["maxAge"].(int) * 86400
	if config.Session["mode"].(string) == "redis" {
		addr := fmt.Sprintf("%s:%d", config.Redis["host"].(string), config.Redis["port"].(int))
		storeRedis, err := session.NewRedisStore(32, "tcp", addr, config.Redis["password"].(string), key)
		storeRedis.Options(session.Options{
			MaxAge: Age,
			Secure: config.Session["secure"].(bool),
			HttpOnly: config.Session["httpOnly"].(bool),
		})
		if err != nil {
			Logger("error", err.Error())
		}	

		return storeRedis
	
	}else {
		
		storeC := sessions.NewCookieStore(key)
		storeC.Options.MaxAge = Age
		storeC.Options.Secure = config.Session["secure"].(bool)
		storeC.Options.HttpOnly = config.Session["httpOnly"].(bool)
		return storeC
	}
}
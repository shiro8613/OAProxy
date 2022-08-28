package modules

import (
	"github.com/gorilla/sessions"
	session "github.com/ipfans/echo-session"
)

func StoreCreate() sessions.Store {
	config := ConfigLoad()
	key := Cryper()
	Age := config.Session["maxAge"].(int) * 86400
	if config.Session["mode"].(string) == "redis" {
		addr := ""
		storeRedis, err := session.NewRedisStore(32,"tcp",addr,"",key)
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
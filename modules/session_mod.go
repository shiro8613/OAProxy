package modules

import (
	"fmt"

	"github.com/gorilla/sessions"
	session "github.com/ipfans/echo-session"
)

func StoreCreate() sessions.Store {
	config := GetConfig()
	key := Cryper()
	Age := config.Session.MaxAge * 86400
	if config.Session.Mode == "redis" {
		addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
		storeRedis, err := session.NewRedisStore(32, "tcp", addr, config.Redis.Password, key)
		storeRedis.Options(session.Options{
			MaxAge: Age,
			Secure: config.Session.Secure,
			HttpOnly: config.Session.HttpOnly,
		})
		if err != nil {
			Logger("error", err.Error())
		}	

		return storeRedis
	
	}else {
		
		storeC := sessions.NewCookieStore(key)
		storeC.Options.MaxAge = Age
		storeC.Options.Secure = config.Session.Secure
		storeC.Options.HttpOnly = config.Session.HttpOnly
		return storeC
	}
}
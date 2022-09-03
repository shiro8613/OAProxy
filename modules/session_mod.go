package modules

import (
	"fmt"

	"github.com/gorilla/sessions"
	"gopkg.in/boj/redistore.v1"

)

func StoreCreate() sessions.Store {
	config := GetConfig()
	key := Cryper()
	Age := config.Session.MaxAge * 86400
	if config.Session.Mode == "redis" {
		addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
		store, err := redistore.NewRediStore(10, "tcp", addr, config.Redis.Password, key)
		if err != nil {
			Logger("error", err.Error())
		}		
		defer store.Close()

		store.DefaultMaxAge = Age
		store.Options.Secure = config.Session.Secure
		store.Options.HttpOnly = config.Session.HttpOnly
		
		return store
	}else {
		store := sessions.NewCookieStore(key)
		store.Options.MaxAge = Age
		store.Options.Secure = config.Session.Secure
		store.Options.HttpOnly = config.Session.HttpOnly

		return store
	}
}
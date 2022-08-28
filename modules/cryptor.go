package modules

import (
	"crypto/rand"
	"fmt"
)

func Cryper() []byte {
	key := [256]byte{}
	_, err := rand.Read(key[:])
	if err != nil {
		Logger("error", err.Error())
	}
		return []byte(fmt.Sprintf("%v",key))
}
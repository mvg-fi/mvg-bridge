package constants

import (
	"github.com/fox-one/mixin-sdk-go"
)

type User struct {
	*mixin.User
	Key *mixin.Keystore
	PIN string
}

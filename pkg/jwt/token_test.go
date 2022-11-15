package jwt

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	t.Run("token", func(t *testing.T) {
		token, err := GenToken(16, "test")
		if err != nil {
			t.Fail()
		}
		claims, err := DecodeToken(token)
		if err != nil {
			t.Fail()
		}
		fmt.Println(claims.UserName, claims.UserId)
	})
}

package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var t *jwt.Token

type jwtBuilder struct {
	key string
	jwt.SigningMethod
	useStore bool
}

func New() *jwtBuilder {
	return &jwtBuilder{
		key:           viper.GetString("app.secret"),
		useStore:      viper.GetBool("jwt.use_store"),
		SigningMethod: jwt.SigningMethodHS256,
	}
}

// IssueById issue jwt token with jti from id
func (b *jwtBuilder) IssueById(id int64) (string, error) {
	// use store: check token exists in storage

	token := jwt.NewWithClaims(b.SigningMethod, jwt.MapClaims{
		"jti": id, "sub": id, "iss": viper.GetString("app.name"),
		"exp": time.Now().Add(time.Hour).UnixMicro(),
	})

	return token.SignedString([]byte(b.key))
}

// Validate determines give token: invalid or expired.
func (b *jwtBuilder) Validate(s string) (error, bool) {
	return nil, false
}

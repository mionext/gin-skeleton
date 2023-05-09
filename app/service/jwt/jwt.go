package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type jwtBuilder struct {
	ttl time.Duration
	key string
	jwt.SigningMethod
	useStore bool
}

func New() *jwtBuilder {
	return &jwtBuilder{
		ttl:           viper.GetDuration("jwt.ttl"),
		key:           viper.GetString("jwt.secret"),
		useStore:      viper.GetBool("jwt.use_store"),
		SigningMethod: jwt.SigningMethodHS256,
	}
}

// WithTTL make jwt with ttl
func WithTTL(ttl time.Duration) *jwtBuilder {
	return &jwtBuilder{
		ttl:           ttl,
		key:           viper.GetString("app.secret"),
		useStore:      viper.GetBool("jwt.use_store"),
		SigningMethod: jwt.SigningMethodHS256,
	}
}

func (b *jwtBuilder) issue(claims jwt.MapClaims) (string, error) {
	claims["iss"] = viper.GetString("app.name")
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(b.ttl * time.Second).Unix()
	token := jwt.NewWithClaims(b.SigningMethod, claims)

	return token.SignedString([]byte(b.key))
}

// IssueById issue jwt token with jti from id
func (b *jwtBuilder) IssueById(id int64) (string, error) {
	return b.issue(jwt.MapClaims{"jti": id})
}

// Validate determines give token: invalid or expired.
func (b *jwtBuilder) Validate(s string) (error, bool) {
	token, err := b.Parse(s)
	if err != nil || !token.Valid {
		return errors.New("Invalid token."), false
	}

	claims, ok := token.Claims.(jwt.Claims)
	fmt.Println(claims, ok)
	if !ok {
		return errors.New("Invalid token."), false
	}

	exp, err := claims.GetExpirationTime()
	if err == nil && exp.Unix() <= time.Now().Unix() {
		return errors.New("Token is expired."), false
	}

	return nil, true
}

// Parse parsing jwt token string.
func (b *jwtBuilder) Parse(s string) (*jwt.Token, error) {
	return jwt.Parse(s, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", t.Header["alg"]))
		}

		return []byte(b.key), nil
	})
}

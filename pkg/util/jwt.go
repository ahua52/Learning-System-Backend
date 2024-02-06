package util

import (
	"HMS/payment/middlewares"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	j := &middlewares.JWT{
		SigningKey: []byte("gp7LI7u5nvqHoTmzRe2iLKexAJAAhZwZwCsCowO39WvJBHktueLpSV6uPHJot0FPKUUOohqGQkCufg3tvhb1BscJyREJGAM7dLdqmDqIEkMNtWnGSkKDcLje5N0KXk5Mx6Z7PqGqHgB3wcqGPjhJhEaYN3VqRURhtynPFhC1JMTem7ovIafp2oaxTfRcvm0vgl39MAB5MJOI4U167orzazR1BQpYmveyZjAp50OLgeUzMO1ditDS3FQSx9XEoRjFr83yrBlOtmAmkprMbUX7hQ2zt3LDJWosAEJNTxylXoSvOMsElN0019IknR6iiDYsT342VrvnIf2q3cOWtMy2LtnElWzjWJ7cnhjCRGetZXkF2ZOcz0fXqLr7fM0kqH2Cu1CwV8MZGf1OoBm6u7Db7uxCzQzSOUwQcEFGpKlZ5lUFUkithcynMBbSJNkBABeaYk96rTqhQUZke9FSBYH5hk9OincdgfUXpqrx389NDGgcf4EHlBf8dPLn2iQ0aTJC"),
	}
	claims := middlewares.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIiLCJuYW1lIjoiIiwicGhvbmUiOiIiLCJleHAiOjE3MDE0NTIxNTcsImlzcyI6Im5ld3RyZWtXYW5nIiwibmJmIjoxNzAxNDQ3NTU3fQ.GbkGVqcgjbX81MABOI2bMyfdtULqeWhcOO-s5jXdedE
	token, err := j.CreateToken(claims)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

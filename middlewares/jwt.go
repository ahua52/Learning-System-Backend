package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret = []byte("gp7LI7u5nvqHoTmzRe2iLKexAJAAhZwZwCsCowO39WvJBHktueLpSV6uPHJot0FPKUUOohqGQkCufg3tvhb1BscJyREJGAM7dLdqmDqIEkMNtWnGSkKDcLje5N0KXk5Mx6Z7PqGqHgB3wcqGPjhJhEaYN3VqRURhtynPFhC1JMTem7ovIafp2oaxTfRcvm0vgl39MAB5MJOI4U167orzazR1BQpYmveyZjAp50OLgeUzMO1ditDS3FQSx9XEoRjFr83yrBlOtmAmkprMbUX7hQ2zt3LDJWosAEJNTxylXoSvOMsElN0019IknR6iiDYsT342VrvnIf2q3cOWtMy2LtnElWzjWJ7cnhjCRGetZXkF2ZOcz0fXqLr7fM0kqH2Cu1CwV8MZGf1OoBm6u7Db7uxCzQzSOUwQcEFGpKlZ5lUFUkithcynMBbSJNkBABeaYk96rTqhQUZke9FSBYH5hk9OincdgfUXpqrx389NDGgcf4EHlBf8dPLn2iQ0aTJC")

func ParseJWT(c *gin.Context) {
	tokenString := c.Query("token")
	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// Don't forget to validate the alg is what you expect:
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	return hmacSampleSecret, nil
	// })
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacSampleSecret), nil
	})
	fmt.Println(11111)
	fmt.Println(err)
	fmt.Println(11111)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 500,
			"msg":  "signature is invalid",
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 500,
			"msg":  "signature is invalid",
		})
		c.Abort()
		return
	}
	c.Next()
}

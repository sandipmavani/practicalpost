package middleware

import (
	"fmt"
	"net/http"
	"practicalpost/controller/authentication"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := authentication.ValidateToken((strings.Trim(tokenString, " ")))
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims["id"])
			c.Set("userId", claims["id"])

			c.Next()
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

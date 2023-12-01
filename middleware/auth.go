package middleware

import (
	"fmt"
	"microservice/handler"
	"microservice/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		// Check if the token is present
		if tokenString == "" {
			handler.ErrorHandler(c, fmt.Errorf("token is missing"), http.StatusBadRequest)
			c.Abort()
			return
		}

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// You should replace this with your secret key
			return []byte("your-secret-key"), nil
		})

		if err != nil {
			handler.ErrorHandler(c, err, http.StatusBadRequest)
			c.Abort()
			return
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			fmt.Println(claims.Email)
			c.Next()
		} else {
			handler.ErrorHandler(c, fmt.Errorf("invalid token"), http.StatusBadRequest)
			c.Abort()
		}
	}
}

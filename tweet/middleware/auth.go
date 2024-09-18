package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/patelajay745/Microservice-VideoApp/tweet/utils"
)

func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenstr := getToken(c)
		if tokenstr == "" {
			c.Abort()
			c.JSON(401, utils.ResMessage{
				Success: false,
				Message: "You are not logged in",
			})
			return
		}

		

		token, parseErr := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected sigining method %v", token.Header["alg"])
			}

			return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil

		})

		if parseErr != nil {
			c.Abort()
			c.JSON(401, utils.ResMessage{
				Success: false,
				Message: "Invalid token",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("_id", claims["_id"])
			c.Next()
		} else {
			c.Abort()
			c.JSON(401, utils.ResMessage{
				Success: false,
				Message: "There was a problem",
			})
		}
	}
}

func getToken(c *gin.Context) string {
	cookie, err := c.Cookie("accessToken")
	if err == nil {
		return cookie
	}

	authHeader := c.GetHeader("Authorization")

	if authHeader != "" {
		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) == 2 {
			return parts[1]
		}
	}

	return ""
}

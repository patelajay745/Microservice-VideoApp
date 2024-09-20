package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/patelajay745/Microservice-VideoApp/subscription/utils"
)

func VerifyJWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenStr := getToken(c)
		if tokenStr == "" {
			c.Abort()
			c.JSON(401, utils.ResMessage{
				Success: false,
				Message: "You are not logged in",
			})
			return
		}

		token, parseErr := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
		})
		// parse error aborts and returns
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
	// Try to get the access token from the cookie
	cookie, err := c.Cookie("accessToken")
	if err == nil {
		return cookie
	}

	// If the cookie is not available, try to get the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) == 2 {
			return parts[1]
		}
	}

	// If neither is available, return an empty string or handle it as needed
	return ""
}

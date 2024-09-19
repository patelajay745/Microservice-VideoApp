package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/patelajay745/Microservice-VideoApp/comment/utils"
)

func VerifyJWT() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenStr := getToken(c)
			if tokenStr == "" {
				return c.JSON(http.StatusUnauthorized, utils.ResMessage{
					Success: false,
					Message: "You are not logged in",
				})
			}

			token, parseErr := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
			})

			if parseErr != nil {
				return c.JSON(http.StatusUnauthorized, utils.ResMessage{
					Success: false,
					Message: "Invalid token",
				})
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("_id", claims["_id"])
				return next(c)
			} else {
				return c.JSON(http.StatusUnauthorized, utils.ResMessage{
					Success: false,
					Message: "There was a problem",
				})
			}
		}
	}
}

// getToken extracts the token from the cookie or Authorization header
func getToken(c echo.Context) string {
	// Try to get the access token from the cookie
	cookie, err := c.Cookie("accessToken")
	if err == nil {
		return cookie.Value
	}

	// If the cookie is not available, try to get the token from the Authorization header
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.Split(authHeader, "Bearer ")
		if len(parts) == 2 {
			return parts[1]
		}
	}

	// If neither is available, return an empty string or handle it as needed
	return ""
}

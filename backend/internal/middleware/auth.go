package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/SKjustSK/secure-task-management/backend/internal/utils" // Imported for Error helper
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// RequireAuth protects routes by verifying the JWT token
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, http.StatusUnauthorized, "Authorization header is missing")
			c.Abort() // Stop the request from reaching the handler
			return
		}

		// 2. Check if it's a Bearer token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Error(c, http.StatusUnauthorized, "Invalid authorization format. Use 'Bearer {token}'")
			c.Abort()
			return
		}
		tokenString := parts[1]

		// 3. Parse and validate the token
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			secret = "fallback_secret_for_local_dev"
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		// 4. Handle expired or tampered tokens
		if err != nil || !token.Valid {
			utils.Error(c, http.StatusUnauthorized, "Session expired or invalid token")
			c.Abort()
			return
		}

		// 5. Extract claims and set context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := uint(claims["user_id"].(float64))

			// Attach the userID to the context so Task handlers can use it
			c.Set("userID", userID)

			c.Next() // Pass control to the next middleware or handler
		} else {
			utils.Error(c, http.StatusUnauthorized, "Failed to parse token claims")
			c.Abort()
			return
		}
	}
}

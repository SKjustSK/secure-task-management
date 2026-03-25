package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a new JWT for a specific user ID
func GenerateToken(userID uint) (string, error) {
	// 1. Get the secret key from the .env file
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "fallback_secret_for_local_dev" // Just in case!
	}

	// 2. Define the payload (claims)
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// 3. Create the token with the HMAC SHA256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 4. Sign the token with our secret key
	return token.SignedString([]byte(secret))
}

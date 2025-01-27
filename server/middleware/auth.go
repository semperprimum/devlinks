package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofor-little/env"
	"github.com/golang-jwt/jwt"
	"github.com/semperprimum/devlinks/server/models"
	"github.com/semperprimum/devlinks/server/utils"
)

type wrappedRequest struct {
	http.Request
	user models.User
}

type key int

const (
	UserIDKey key = iota
)

func Auth(next http.Handler) http.Handler {
	secret, err := env.MustGet("JWT_SECRET")
	if err != nil {
		panic(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriteJSONError(w, "Authorization header it required", http.StatusBadRequest)
			return
		}

		// The token usually comes in the format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteJSONError(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// Parse and verify the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Make sure the token method conforms to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			utils.WriteJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Check if the token is valid and extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Retrieve the user ID from the token claims
			userID, ok := claims["id"].(string)
			if !ok {
				utils.WriteJSONError(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			// Store the user ID in the request context
			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			utils.WriteJSONError(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})
}

func GetUser(r *http.Request) string {
	if id, ok := r.Context().Value(UserIDKey).(string); ok {
		return id
	}

	return ""
}

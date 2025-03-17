package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"main.go/config"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AuthMiddleware JWT doğrulama ve yetkilendirme yapar
func AuthMiddleware(next http.Handler, requiredRole string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userRole := (*claims)["role"].(string)
		userID := (*claims)["user_id"].(string)

		if requiredRole != "" && userRole != requiredRole {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "role", userRole)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromToken token içinden kullanıcıyı çek
func GetUserFromToken(tokenString string) (bson.M, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	userID := (*claims)["user_id"].(string)

	collection := config.GetCollection("users")
	var user bson.M
	err = collection.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("User not found")
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

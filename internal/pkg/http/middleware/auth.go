package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/PhongVX/taskmanagement/internal/pkg/jwt"
)

func EnsureAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
			return
		}
		//TODO Need to read from env file in future
		os.Setenv("ACCESS_SECRET", "my-access-secret")
		tokenString := authHeader[1]
		claims, err := jwt.VerifyToken(tokenString, os.Getenv("ACCESS_SECRET"))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

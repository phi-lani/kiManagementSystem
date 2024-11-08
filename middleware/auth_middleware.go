package middleware

import (
	"context"
	"net/http"

	"github.com/phi-lani/kimanagementsystem/utils"
)

// TokenValidationMiddleware is a middleware function that checks for a valid JWT token
func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the token from the cookie
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}

		// Verify the token using the VerifyJWT function
		claims, err := utils.VerifyJWT(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract the userID and role from the claims
		userID := claims.UserID
		role := claims.Role

		// Set the userID in the request context
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "role", role)
		r = r.WithContext(ctx)

		// If the token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

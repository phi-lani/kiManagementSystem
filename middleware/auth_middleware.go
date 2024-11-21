package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/phi-lani/kimanagementsystem/utils"
)

// // TokenValidationMiddleware is a middleware function that checks for a valid JWT token
// func TokenValidationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Retrieve the token from the cookie
// 		cookie, err := r.Cookie("token")
// 		if err != nil {
// 			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
// 			return
// 		}

// 		// Verify the token using the VerifyJWT function
// 		claims, err := utils.VerifyJWT(cookie.Value)
// 		if err != nil {
// 			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		// Extract the userID and role from the claims
// 		userID := claims.UserID
// 		role := claims.Role

// 		// Set the userID in the request context
// 		ctx := context.WithValue(r.Context(), "userID", userID)
// 		ctx = context.WithValue(ctx, "role", role)
// 		r = r.WithContext(ctx)

// 		// If the token is valid, proceed to the next handler
// 		next.ServeHTTP(w, r)
// 	})
// }

// AdminOnly ensures that only users with the "admin" role can access the endpoint
func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value("role").(string)
		log.Printf("The current role is %s: ", role)
		if !ok || role != "admin" {
			http.Error(w, "Forbidden: Admins access only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// KeyIndividualOnly checks if the user has a key individual role
func KeyIndividualOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value("role").(string)
		log.Printf("The current role is %s: ", role)
		if !ok || role != "key_individual" {
			http.Error(w, "Forbidden: Key Individual access only", http.StatusForbidden)
			return
		}

		// If the user is a key individual, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// StartupOnly checks if the user has a startup role
func StartupOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value("role").(string)
		if !ok || role != "startup" {
			http.Error(w, "Forbidden: Startup access only", http.StatusForbidden)
			return
		}

		// If the user is a startup, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

// // TokenValidationMiddleware is a middleware function that checks for a valid JWT token
// func TokenValidationMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Retrieve the token from the cookie
// 		cookie, err := r.Cookie("token")
// 		if err != nil {
// 			log.Println("Unauthorized: No token provided")
// 			http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
// 			return
// 		}

// 		// Verify the token using the VerifyJWT function
// 		claims, err := utils.VerifyJWT(cookie.Value)
// 		if err != nil {
// 			log.Printf("Unauthorized: Invalid token, error: %v", err)
// 			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		// Extract the userID and role from the claims
// 		userID := claims.UserID
// 		role := claims.Role

// 		// Log the extracted userID and role for debugging
// 		log.Printf("Token valid: userID=%v, role=%v", userID, role)

// 		// Set the userID and role in the request context
// 		ctx := context.WithValue(r.Context(), "userID", userID)
// 		ctx = context.WithValue(ctx, "role", role)
// 		r = r.WithContext(ctx)

// 		// If the token is valid, proceed to the next handler
// 		next.ServeHTTP(w, r)
// 	})
// }

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string

		// Check for token in the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			// Strip the "Bearer " prefix if present
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// Fallback to retrieving the token from the cookie
			cookie, err := r.Cookie("token")
			if err != nil {
				log.Println("Unauthorized: No token provided")
				http.Error(w, "Unauthorized: No token provided", http.StatusUnauthorized)
				return
			}
			token = cookie.Value
		}

		// Verify the token using the VerifyJWT function
		claims, err := utils.VerifyJWT(token)
		if err != nil {
			log.Printf("Unauthorized: Invalid token, error: %v", err)
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract the userID and role from the claims
		userID := claims.UserID
		role := claims.Role

		// Log the extracted userID and role for debugging
		log.Printf("Token valid: userID=%v, role=%v", userID, role)

		// Set the userID and role in the request context
		ctx := context.WithValue(r.Context(), "userID", userID)
		ctx = context.WithValue(ctx, "role", role)
		r = r.WithContext(ctx)

		// If the token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

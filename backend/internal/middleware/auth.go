package middleware

import (
	"context"
	"strings"

	"net/http"

	"github.com/durianpay/fullstack-boilerplate/internal/entity"
	"github.com/durianpay/fullstack-boilerplate/internal/transport"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
)

func Auth(swagger *openapi3.T, jwtSecret []byte) func(http.Handler) http.Handler {
	router, err := legacy.NewRouter(swagger)
	if err != nil {
		// This should only happen if the swagger spec is invalid
		panic("auth middleware: failed to initialize openapi router: " + err.Error())
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Find the route in the swagger spec
			route, _, err := router.FindRoute(r)
			if err != nil {
				// If route not found, let the standard router handle 404
				next.ServeHTTP(w, r)
				return
			}

			// Check if the operation has security defined
			// In OpenAPI 3, security can be defined at the operation level or globally at the root.
			security := route.Operation.Security
			if security == nil {
				// Fallback to global security if operation-level is not defined
				security = &swagger.Security
			}

			// If no security is defined for this route, skip authentication
			if security == nil || len(*security) == 0 {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				transport.WriteError(w, entity.ErrorUnauthorized("missing authorization header"))
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				transport.WriteError(w, entity.ErrorUnauthorized("invalid authorization header format"))
				return
			}

			tokenString := parts[1]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, entity.ErrorUnauthorized("unexpected signing method")
				}
				return jwtSecret, nil
			})

			if err != nil || !token.Valid {
				transport.WriteError(w, entity.ErrorUnauthorized("invalid or expired token"))
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				transport.WriteError(w, entity.ErrorUnauthorized("invalid token claims"))
				return
			}

			userID, ok := claims["sub"].(string)
			if !ok {
				transport.WriteError(w, entity.ErrorUnauthorized("invalid token subject"))
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

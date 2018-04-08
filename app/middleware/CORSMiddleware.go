package middleware

import "github.com/rs/cors"

// CORSMiddleware ...
func CORSMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		AllowCredentials: true,
	})
}

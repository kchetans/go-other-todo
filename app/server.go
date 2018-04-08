package app

import (
	"github.com/go-to-do/app/middleware"
	"github.com/urfave/negroni"
)

// NewServer ...
func NewServer() *negroni.Negroni {

	// Define the global middlewares
	server := negroni.New()
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.LogMiddleware())

	// Attach app router
	server.UseHandler(AllRoutes())

	return server
}

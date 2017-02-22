package main

import (
	"github.com/dinever/golf"
	s "github.com/jkunii/json-going/services"
)

func main() {
	app := golf.New()
	app.View.SetTemplateLoader("default", ".")
	app.SessionManager = golf.NewMemorySessionManager()
	app.Use(golf.SessionMiddleware)

	app.Get("/", s.MainHandler)
	app.Post("/login", s.LoginHandlerPost)
	app.Get("/login", s.LoginHandler)

	app.Get("/key", s.GoMockHandler)
	app.Post("/key", s.CreateMockResponse)

	app.Run(":9000")
}

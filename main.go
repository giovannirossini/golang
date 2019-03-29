package main

import (
	"os"

	routers "github.com/giovannirossini/curso_web/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ikeikeikeike/pongor"
	"github.com/labstack/echo/middleware"
)

// Start the routes paths
func main() {
	server := routers.App

	// Start the render for views
	p := pongor.GetRenderer()
	p.Directory = "views"

	server.Renderer = p
	port := os.Getenv("PORT")
	server.Use(middleware.Logger())
	server.Logger.Fatal(server.Start(":" + port))
}

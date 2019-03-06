package main

import (
	routers "github.com/giovannirossini/curso/routers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ikeikeikeike/pongor"
	"github.com/labstack/echo/middleware"
)

func main() {
	server := routers.App

	p := pongor.GetRenderer()
	p.Directory = "views"

	server.Renderer = p

	server.Use(middleware.Logger())
	server.Logger.Fatal(server.Start(":3000"))
}

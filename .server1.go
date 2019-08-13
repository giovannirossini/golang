package main

import (
	"net/http"
    "os"
	"github.com/labstack/echo"
)

//  Easy server on Go
func main() {
	server := echo.New()
    port := os.Getenv("PORT")

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	server.Start(":" + port)
}

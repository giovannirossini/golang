package routers

import (
	"github.com/giovannirossini/curso/controllers"
	"github.com/labstack/echo"
)

// App variable to use Echo framework
var App *echo.Echo

// Start server and map the routes
func init() {
	App = echo.New()
	api := App.Group("/api")

	// Controllers calls and paths
	App.GET("/", controllers.Home)
	App.GET("/add", controllers.Add)
	App.GET("/edit/:id", controllers.Edit)

	api.POST("/insert", controllers.Post)
	api.DELETE("/delete/:id", controllers.Delete)
	api.PUT("/put/:id", controllers.Put)

}

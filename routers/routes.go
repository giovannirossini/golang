package routers

import (
	"github.com/giovannirossini/curso/controllers"
	"github.com/labstack/echo"
)

// App variable to init Echo framework
var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controllers.Home)
	App.GET("/add", controllers.Add)
	App.GET("/edit/:id", controllers.Edit)

	api := App.Group("/api")

	api.POST("/insert", controllers.Post)
	api.DELETE("/delete/:id", controllers.Delete)
	api.PUT("/put/:id", controllers.Put)

}

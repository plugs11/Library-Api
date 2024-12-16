package routes

import (
	"github.com/gin-gonic/gin"
	"manan.tola/controllers"
)

func RegisterBookRoutes(router *gin.Engine) {
	cluster := router.Group("/book")
	{
		cluster.POST("/", controllers.CreateBook)
		cluster.GET("/id/:id", controllers.GetBookByID)
		cluster.GET("/", controllers.GetBookByID)
		cluster.PUT("/:id", controllers.UpdateBook)
		cluster.DELETE("/:id", controllers.DeleteBook)
	}

} // calls all the routes depending upon request

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api_gimnasios/controllers"
)

func GimnasioClaseRoutes(router *gin.Engine) {
	grupo := router.Group("/api/v1/clases")
	{
		grupo.GET("/", controllers.GetClases)
		grupo.GET("/:id", controllers.GetClaseByID)
		grupo.POST("/", controllers.CreateClase)
		grupo.PUT("/:id", controllers.UpdateClase)
		grupo.DELETE("/:id", controllers.DeleteClase)
	}
}
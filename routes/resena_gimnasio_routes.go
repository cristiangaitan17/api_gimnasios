package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api_gimnasios/controllers"
)

func ResenaGimnasioRoutes(router *gin.Engine) {
	grupo := router.Group("/api/v1/resenas")
	{
		grupo.GET("/", controllers.GetResenas)
		grupo.GET("/:id", controllers.GetResenaByID)
		grupo.POST("/", controllers.CreateResena)
		grupo.PUT("/:id", controllers.UpdateResena)
		grupo.DELETE("/:id", controllers.DeleteResena)
	}
}
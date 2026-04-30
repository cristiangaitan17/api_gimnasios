package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api_gimnasios/controllers"
)

func SedeRoutes(router *gin.Engine) {
	grupo := router.Group("/api/v1/sedes")
	{
		grupo.GET("/", controllers.GetSedes)
		grupo.GET("/:id", controllers.GetSedeByID)
		grupo.POST("/", controllers.CreateSede)
		grupo.PUT("/:id", controllers.UpdateSede)
		grupo.DELETE("/:id", controllers.DeleteSede)
	}
}
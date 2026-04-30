package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api_gimnasios/config"
	"github.com/cristiangaitan17/api_gimnasios/routes"
)

func main() {
	config.InitDB()

	router := gin.Default()
	
	routes.SedeRoutes(router)
	routes.GimnasioClaseRoutes(router)
	routes.ResenaGimnasioRoutes(router)

	log.Println("🚀 Servidor API Gimnasios corriendo en http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
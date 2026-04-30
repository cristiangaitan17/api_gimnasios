package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api_gimnasios/config"
	"github.com/cristiangaitan17/api_gimnasios/models"
)

// GetResenas obtiene todas las reseñas
func GetResenas(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, gimnasio_id, usuario_id, calificacion, 
		       COALESCE(comentario, ''), activo
		FROM gimnasios.resenas_gimnasio
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var resenas []models.ResenaGimnasio
	for rows.Next() {
		var r models.ResenaGimnasio
		err := rows.Scan(
			&r.ID, &r.GimnasioID, &r.UsuarioID, &r.Calificacion,
			&r.Comentario, &r.Activo,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		resenas = append(resenas, r)
	}
	c.JSON(http.StatusOK, resenas)
}

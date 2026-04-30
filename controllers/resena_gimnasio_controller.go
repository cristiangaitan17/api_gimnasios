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

// GetResenaByID obtiene una reseña por ID
func GetResenaByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var r models.ResenaGimnasio
	row := config.DB.QueryRow(`
		SELECT id, gimnasio_id, usuario_id, calificacion, 
		       COALESCE(comentario, ''), activo
		FROM gimnasios.resenas_gimnasio WHERE id = $1
	`, id)

	err = row.Scan(
		&r.ID, &r.GimnasioID, &r.UsuarioID, &r.Calificacion,
		&r.Comentario, &r.Activo,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reseña no encontrada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, r)
}

// UpdateResena actualiza una reseña existente
func UpdateResena(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var r models.ResenaGimnasio
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar calificación entre 1 y 5
	if r.Calificacion < 1 || r.Calificacion > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La calificación debe ser entre 1 y 5"})
		return
	}

	query := `
		UPDATE gimnasios.resenas_gimnasio 
		SET gimnasio_id = $1, usuario_id = $2, calificacion = $3, 
		    comentario = $4, activo = $5
		WHERE id = $6
	`
	result, err := config.DB.Exec(query, r.GimnasioID, r.UsuarioID, r.Calificacion,
		r.Comentario, r.Activo, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reseña no encontrada"})
		return
	}
	r.ID = id
	c.JSON(http.StatusOK, r)
}
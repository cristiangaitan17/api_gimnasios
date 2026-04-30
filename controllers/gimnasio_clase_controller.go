package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api_gimnasios/config"
	"github.com/cristiangaitan17/api_gimnasios/models"
)

// GetClases obtiene todas las clases
func GetClases(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, gimnasio_id, nombre_clase, activo
		FROM gimnasios.gimnasio_clases
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var clases []models.GimnasioClase
	for rows.Next() {
		var cl models.GimnasioClase
		err := rows.Scan(&cl.ID, &cl.GimnasioID, &cl.NombreClase, &cl.Activo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		clases = append(clases, cl)
	}
	c.JSON(http.StatusOK, clases)
}

// GetClaseByID obtiene una clase por ID
func GetClaseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var cl models.GimnasioClase
	row := config.DB.QueryRow(`
		SELECT id, gimnasio_id, nombre_clase, activo
		FROM gimnasios.gimnasio_clases WHERE id = $1
	`, id)

	err = row.Scan(&cl.ID, &cl.GimnasioID, &cl.NombreClase, &cl.Activo)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cl)
}

// CreateClase crea una nueva clase
func CreateClase(c *gin.Context) {
	var cl models.GimnasioClase
	if err := c.ShouldBindJSON(&cl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO gimnasios.gimnasio_clases (gimnasio_id, nombre_clase, activo)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var id int
	err := config.DB.QueryRow(query, cl.GimnasioID, cl.NombreClase, cl.Activo).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	cl.ID = id
	c.JSON(http.StatusCreated, cl)
}

// UpdateClase actualiza una clase existente
func UpdateClase(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var cl models.GimnasioClase
	if err := c.ShouldBindJSON(&cl); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE gimnasios.gimnasio_clases 
		SET gimnasio_id = $1, nombre_clase = $2, activo = $3
		WHERE id = $4
	`
	result, err := config.DB.Exec(query, cl.GimnasioID, cl.NombreClase, cl.Activo, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}
	cl.ID = id
	c.JSON(http.StatusOK, cl)
}

// DeleteClase elimina una clase
func DeleteClase(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result, err := config.DB.Exec("DELETE FROM gimnasios.gimnasio_clases WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clase no encontrada"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
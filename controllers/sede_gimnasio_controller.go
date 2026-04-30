package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cristiangaitan17/api_gimnasios/config"
	"github.com/cristiangaitan17/api_gimnasios/models"
)

// GetSedes obtiene todas las sedes
func GetSedes(c *gin.Context) {
	rows, err := config.DB.Query(`
		SELECT id, nombre, nit, descripcion, ciudad, departamento, 
		       direccion, correo, telefono, agregar_img, agregar_sede,
		       aprovacion_entrenadores, calificacion_prom, total_resenas, 
		       activo, administrador_id
		FROM gimnasios.sedes_gimnasios
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var sedes []models.SedeGimnasio
	for rows.Next() {
		var s models.SedeGimnasio
		err := rows.Scan(
			&s.ID, &s.Nombre, &s.Nit, &s.Descripcion, &s.Ciudad, &s.Departamento,
			&s.Direccion, &s.Correo, &s.Telefono, &s.AgregarImg, &s.AgregarSede,
			&s.AprovacionEntrenadores, &s.CalificacionProm, &s.TotalResenas,
			&s.Activo, &s.AdministradorID,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		sedes = append(sedes, s)
	}
	c.JSON(http.StatusOK, sedes)
}

// GetSedeByID obtiene una sede por ID
func GetSedeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var s models.SedeGimnasio
	row := config.DB.QueryRow(`
		SELECT id, nombre, nit, descripcion, ciudad, departamento, 
		       direccion, correo, telefono, agregar_img, agregar_sede,
		       aprovacion_entrenadores, calificacion_prom, total_resenas, 
		       activo, administrador_id
		FROM gimnasios.sedes_gimnasios WHERE id = $1
	`, id)

	err = row.Scan(
		&s.ID, &s.Nombre, &s.Nit, &s.Descripcion, &s.Ciudad, &s.Departamento,
		&s.Direccion, &s.Correo, &s.Telefono, &s.AgregarImg, &s.AgregarSede,
		&s.AprovacionEntrenadores, &s.CalificacionProm, &s.TotalResenas,
		&s.Activo, &s.AdministradorID,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sede no encontrada"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, s)
}

// CreateSede crea una nueva sede
func CreateSede(c *gin.Context) {
	var s models.SedeGimnasio
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO gimnasios.sedes_gimnasios 
		(nombre, nit, descripcion, ciudad, departamento, direccion, 
		 correo, telefono, agregar_img, agregar_sede, aprovacion_entrenadores, 
		 calificacion_prom, total_resenas, activo, administrador_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		RETURNING id
	`
	var id int
	err := config.DB.QueryRow(query, s.Nombre, s.Nit, s.Descripcion, s.Ciudad,
		s.Departamento, s.Direccion, s.Correo, s.Telefono, s.AgregarImg,
		s.AgregarSede, s.AprovacionEntrenadores, s.CalificacionProm,
		s.TotalResenas, s.Activo, s.AdministradorID).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.ID = id
	c.JSON(http.StatusCreated, s)
}

// UpdateSede actualiza una sede existente
func UpdateSede(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var s models.SedeGimnasio
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		UPDATE gimnasios.sedes_gimnasios 
		SET nombre = $1, nit = $2, descripcion = $3, ciudad = $4, 
		    departamento = $5, direccion = $6, correo = $7, telefono = $8,
		    agregar_img = $9, agregar_sede = $10, aprovacion_entrenadores = $11,
		    calificacion_prom = $12, total_resenas = $13, activo = $14, administrador_id = $15
		WHERE id = $16
	`
	result, err := config.DB.Exec(query, s.Nombre, s.Nit, s.Descripcion, s.Ciudad,
		s.Departamento, s.Direccion, s.Correo, s.Telefono, s.AgregarImg,
		s.AgregarSede, s.AprovacionEntrenadores, s.CalificacionProm,
		s.TotalResenas, s.Activo, s.AdministradorID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sede no encontrada"})
		return
	}
	s.ID = id
	c.JSON(http.StatusOK, s)
}

// DeleteSede elimina una sede
func DeleteSede(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	result, err := config.DB.Exec("DELETE FROM gimnasios.sedes_gimnasios WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sede no encontrada"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
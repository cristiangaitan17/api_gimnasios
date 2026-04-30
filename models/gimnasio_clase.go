package models

type GimnasioClase struct {
	ID                int    `json:"id"`
	GimnasioID        int    `json:"gimnasio_id"`
	NombreClase       string `json:"nombre_clase"`
	Activo            bool   `json:"activo"`
	FechaModificacion string `json:"fecha_modificacion"`
	FechaCreacion     string `json:"fecha_creacion"`
}

package models

type ResenaGimnasio struct {
	ID                int     `json:"id"`
	GimnasioID        int     `json:"gimnasio_id"`
	UsuarioID         int     `json:"usuario_id"`
	Calificacion      int     `json:"calificacion"`
	Comentario        string  `json:"comentario"`
	Activo            bool    `json:"activo"`
	FechaModificacion string  `json:"fecha_modificacion"`
	FechaCreacion     string  `json:"fecha_creacion"`
}
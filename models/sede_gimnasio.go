package models

type SedeGimnasio struct {
	ID                  int     `json:"id"`
	Nombre              string  `json:"nombre"`
	Nit                 string  `json:"nit"`
	Descripcion         string  `json:"descripcion"`
	Ciudad              string  `json:"ciudad"`
	Departamento        string  `json:"departamento"`
	Direccion           string  `json:"direccion"`
	Correo              string  `json:"correo"`
	Telefono            string  `json:"telefono"`
	AgregarImg          string  `json:"agregar_img"`
	AgregarSede         string  `json:"agregar_sede"`
	AprovacionEntrenadores bool `json:"aprovacion_entrenadores"`
	CalificacionProm    float32 `json:"calificacion_prom"`
	TotalResenas        int     `json:"total_resenas"`
	Activo              bool    `json:"activo"`
	AdministradorID     int     `json:"administrador_id"`
	FechaModificacion   string  `json:"fecha_modificacion"`
	FechaCreacion       string  `json:"fecha_creacion"`
}
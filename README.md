# api_gimnasios

El API provee la gestión de los diferentes procesos relacionados con gimnasios, incluyendo sedes, clases y reseñas.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
- Golang
- Gin Framework
- PostgreSQL
- Git

### Variables de Entorno

| Variable | Descripción |
|----------|-------------|
| `DB_HOST` | Dirección del servidor de base de datos |
| `DB_PORT` | Puerto de conexión con la base de datos |
| `DB_USER` | Usuario con acceso a la base de datos |
| `DB_PASS` | Password del usuario |
| `DB_NAME` | Nombre de la base de datos |

**NOTA:** Las variables se deben configurar en un archivo `.env` en la raíz del proyecto.

## Estructura del Proyecto
api_gimnasios/
├── config/
│ └── db.go # Configuración y conexión a BD
├── controllers/
│ ├── sede_gimnasio_controller.go
│ ├── gimnasio_clase_controller.go
│ └── resena_gimnasio_controller.go
├── models/
│ ├── sede_gimnasio.go
│ ├── gimnasio_clase.go
│ └── resena_gimnasio.go
├── routes/
│ ├── sede_routes.go
│ ├── gimnasio_clase_routes.go
│ └── resena_gimnasio_routes.go
├── main.go
├── go.mod
├── go.sum
├── .env
└── .gitignore

text

## Endpoints de la API

### Sedes (`/api/v1/sedes`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/sedes` | Obtener todas las sedes |
| GET | `/api/v1/sedes/{id}` | Obtener una sede por ID |
| POST | `/api/v1/sedes` | Crear una nueva sede |
| PUT | `/api/v1/sedes/{id}` | Actualizar una sede |
| DELETE | `/api/v1/sedes/{id}` | Eliminar una sede |

### Clases (`/api/v1/clases`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/clases` | Obtener todas las clases |
| GET | `/api/v1/clases/{id}` | Obtener una clase por ID |
| POST | `/api/v1/clases` | Crear una nueva clase |
| PUT | `/api/v1/clases/{id}` | Actualizar una clase |
| DELETE | `/api/v1/clases/{id}` | Eliminar una clase |

### Reseñas (`/api/v1/resenas`)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/api/v1/resenas` | Obtener todas las reseñas |
| GET | `/api/v1/resenas/{id}` | Obtener una reseña por ID |
| POST | `/api/v1/resenas` | Crear una nueva reseña |
| PUT | `/api/v1/resenas/{id}` | Actualizar una reseña |
| DELETE | `/api/v1/resenas/{id}` | Eliminar una reseña |

## Ejecución del Proyecto

### 1. Clonar el repositorio

```bash
git clone https://github.com/cristiangaitan17/api_gimnasios.git
cd api_gimnasios
2. Configurar variables de entorno
Crear un archivo .env en la raíz del proyecto:

env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=tu_contraseña
DB_NAME=gimnasio_db
3. Instalar dependencias
bash
go mod tidy
4. Ejecutar el proyecto
bash
go run main.go
El servidor correrá en http://localhost:8080

Ejemplo de petición POST
bash
curl -X POST http://localhost:8080/api/v1/sedes \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Gimnasio Central",
    "nit": "900123456-1",
    "ciudad": "Bogotá",
    "departamento": "Cundinamarca",
    "activo": true
  }'
Dependencias
Paquete	Propósito
github.com/gin-gonic/gin	Framework web
github.com/lib/pq	Driver para PostgreSQL
github.com/joho/godotenv	Manejo de variables de entorno
Estado del Proyecto
✅ Completo - API con CRUD completo para las tablas:

sedes_gimnasios

gimnasio_clases

resenas_gimnasio

Autor
Cristian Gaitán

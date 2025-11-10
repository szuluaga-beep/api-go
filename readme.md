# API REST en Go

Una API REST simple construida con Go y el framework Gin para la gestión de usuarios.

## Características

- CRUD completo para usuarios (Crear, Leer, Actualizar, Eliminar)
- API RESTful con endpoints JSON
- Almacenamiento en memoria
- Construido con Gin Framework

## Requisitos Previos

- Go 1.25.4 o superior
- Git (opcional, para clonar el repositorio)

## Instalación

### 1. Clonar el repositorio (si aplica)
```bash
git clone https://github.com/szuluaga-beep/mi-api-go.git
cd mi-api-go
```

### 2. Descargar dependencias
```bash
go mod download
```

## Ejecución

### Ejecutar el servidor
```bash
go run main.go
```

El servidor se iniciará en `http://localhost:8080`

### Ejecutar en modo producción
```bash
go build -o mi-api
```

## Endpoints de la API

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/usuarios` | Obtener todos los usuarios |
| GET | `/usuarios/:id` | Obtener un usuario por ID |
| POST | `/usuarios` | Crear un nuevo usuario |
| PUT | `/usuarios/:id` | Actualizar un usuario existente |
| DELETE | `/usuarios/:id` | Eliminar un usuario |

## Ejemplos de Uso

### 1. Obtener todos los usuarios
```bash
curl http://localhost:8080/usuarios
```

### 2. Crear un nuevo usuario
```bash
curl -X POST http://localhost:8080/usuarios \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Juan Pérez",
    "email": "juan@example.com"
  }'
```

### 3. Obtener un usuario específico
```bash
curl http://localhost:8080/usuarios/1
```

### 4. Actualizar un usuario
```bash
curl -X PUT http://localhost:8080/usuarios/1 \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Juan Carlos Pérez",
    "email": "juancarlos@example.com"
  }'
```

### 5. Eliminar un usuario
```bash
curl -X DELETE http://localhost:8080/usuarios/1
```

## Estructura del Proyecto

```
api-go/
├── main.go          # Archivo principal con la lógica de la API
├── go.mod           # Definición del módulo y dependencias
└── readme.md        # Este archivo
```

## Estructura de Datos

### Usuario
```json
{
  "id": 1,
  "nombre": "Juan Pérez",
  "email": "juan@example.com"
}
```

## Comandos Útiles

### Verificar versión de Go
```bash
go version
```

### Limpiar caché de módulos
```bash
go clean -modcache
```

### Verificar dependencias
```bash
go mod verify
```

### Actualizar dependencias
```bash
go mod tidy
```

## Desarrollo

### Ejecutar con recarga automática (opcional)
Para desarrollo, puedes usar `air` para recarga automática:

```bash
# Instalar air
go install github.com/cosmtrek/air@latest

# Ejecutar con recarga automática
air
```

## Notas Importantes

- Los datos se almacenan en memoria, por lo que se perderán al reiniciar el servidor
- El puerto por defecto es 8080
- La API no incluye autenticación ni validación avanzada
- Este es un proyecto de demostración/aprendizaje

## Solución de Problemas

### Puerto en uso
Si el puerto 8080 está ocupado, puedes cambiarlo modificando la línea en `main.go`:
```go
router.Run(":8080") // Cambiar por otro puerto, ej: ":3000"
```

### Dependencias faltantes
Si tienes problemas con las dependencias, ejecuta:
```bash
go mod download
go mod tidy
```

## Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## Licencia

Este proyecto es de código abierto y está disponible bajo la [MIT License](LICENSE).
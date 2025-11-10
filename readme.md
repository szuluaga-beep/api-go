# API REST en Go con Concurrencia

Una API REST avanzada construida con Go y el framework Gin para la gestión de usuarios, implementando **goroutines** y concurrencia para alto rendimiento.

## 🚀 Características

- **CRUD completo** para usuarios (Crear, Leer, Actualizar, Eliminar)
- **API RESTful** con endpoints JSON
- **Concurrencia con Goroutines** - Operaciones asíncronas y thread-safe
- **Almacenamiento en memoria** con protección de concurrencia (sync.RWMutex)
- **Logging asíncrono** para mejor rendimiento
- **Inicialización en background** con datos de ejemplo
- **Endpoint de estadísticas** en tiempo real
- **Construido con Gin Framework**
- **Thread-safe** - Manejo seguro de múltiples requests concurrentes

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
go build -o mi-api main.go
./mi-api
```

## 🎯 Endpoints de la API

| Método | Endpoint | Descripción | Concurrencia |
|--------|----------|-------------|--------------|
| GET | `/usuarios` | Obtener todos los usuarios | ✅ Asíncrono |
| GET | `/usuarios/:id` | Obtener un usuario por ID | ✅ Búsqueda asíncrona |
| POST | `/usuarios` | Crear un nuevo usuario | ✅ Creación asíncrona |
| PUT | `/usuarios/:id` | Actualizar un usuario existente | ✅ Actualización asíncrona |
| DELETE | `/usuarios/:id` | Eliminar un usuario | ✅ Eliminación asíncrona |
| GET | `/usuarios/stats` | **NUEVO** - Estadísticas del sistema | ✅ Procesamiento asíncrono |

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

### 6. **NUEVO** - Obtener estadísticas del sistema
```bash
curl http://localhost:8080/usuarios/stats
```

**Respuesta esperada:**
```json
{
  "total_usuarios": 3,
  "usuario_mas_reciente": {
    "id": 3,
    "nombre": "Carlos López",
    "email": "carlos@example.com",
    "creado": "2025-11-10T10:30:45Z"
  },
  "timestamp": "2025-11-10T10:35:22Z"
}
```

## 📁 Estructura del Proyecto

```
api-go/
├── main.go          # Archivo principal con lógica de API + Goroutines
├── go.mod           # Definición del módulo y dependencias
└── readme.md        # Este archivo
```

## 🏗️ Arquitectura de Concurrencia

### UsuarioService
- **Thread-safe** con `sync.RWMutex`
- **Lecturas concurrentes** permitidas (`RLock()`)
- **Escrituras exclusivas** (`Lock()`)
- **Operaciones asíncronas** con goroutines

### Características de Concurrencia
- ✅ **Inicialización asíncrona** - Datos de ejemplo cargados en background
- ✅ **Logging no bloqueante** - Logs procesados en goroutines separadas
- ✅ **Operaciones de BD asíncronas** - Búsquedas y actualizaciones concurrentes
- ✅ **Canales para comunicación** - `resultChan` y `errorChan`
- ✅ **Select statements** - Manejo de respuestas concurrentes

## 📊 Estructura de Datos

### Usuario (Actualizada)
```json
{
  "id": 1,
  "nombre": "Juan Pérez",
  "email": "juan@example.com",
  "creado": "2025-11-10T10:30:45Z"
}
```

### Estadísticas del Sistema
```json
{
  "total_usuarios": 3,
  "usuario_mas_reciente": {
    "id": 3,
    "nombre": "Carlos López", 
    "email": "carlos@example.com",
    "creado": "2025-11-10T10:30:45Z"
  },
  "timestamp": "2025-11-10T10:35:22Z"
}
```

## ⚡ Beneficios de Performance

### Antes (Versión Síncrona)
- ❌ Una operación a la vez
- ❌ Logging bloqueante
- ❌ Sin protección de concurrencia
- ❌ Potenciales race conditions

### Ahora (Con Goroutines)
- ✅ **Múltiples operaciones simultáneas**
- ✅ **Logging asíncrono** - No afecta tiempo de respuesta
- ✅ **Thread-safe** - Operaciones concurrentes seguras
- ✅ **Mejor escalabilidad** - Manejo eficiente de carga alta
- ✅ **Inicialización no bloqueante** - Servidor inicia más rápido

## 🛠️ Comandos Útiles

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

### Ejecutar tests de concurrencia (opcional)
```bash
# Instalar herramienta de testing de concurrencia
go install golang.org/x/tools/cmd/stress@latest

# Test de stress
stress go test -race ./...
```

## 🔧 Desarrollo

### Ejecutar con recarga automática (opcional)
Para desarrollo, puedes usar `air` para recarga automática:

```bash
# Instalar air
go install github.com/cosmtrek/air@latest

# Ejecutar con recarga automática
air
```

### Debugging de Goroutines
```bash
# Ejecutar con race detection
go run -race main.go

# Compilar con race detection
go build -race -o mi-api main.go
```

## ⚠️ Notas Importantes

- **Almacenamiento en memoria** - Los datos se perderán al reiniciar el servidor
- **Puerto por defecto**: 8080
- **Thread-safe** - Múltiples requests pueden ser procesados simultáneamente
- **Goroutines activas** - El servidor utiliza concurrencia para mejor rendimiento
- **Datos de ejemplo** - Se cargan automáticamente 3 usuarios al iniciar
- **Este es un proyecto de demostración** de concurrencia en Go

## 🛠️ Solución de Problemas

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

### Race Conditions
Si sospechas problemas de concurrencia:
```bash
# Ejecutar con detección de race conditions
go run -race main.go
```

### Problemas de Performance
Para monitorear goroutines:
```bash
# Instalar pprof para profiling
go tool pprof http://localhost:8080/debug/pprof/goroutine
```

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## Licencia

Este proyecto es de código abierto y está disponible bajo la [MIT License](LICENSE).
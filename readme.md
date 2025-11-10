# API REST - Comparativa de Concurrencia en Go

Este proyecto demuestra la diferencia entre una API REST sin concurrencia y otra implementada con Goroutines en Go.

## 📋 Descripción

El proyecto contiene una API REST que gestiona usuarios con 4 endpoints principales. La particularidad es que hay **dos ramas**:
- **`master`**: Implementación sin concurrencia (secuencial)
- **`concurrent`**: Implementación con Goroutines (paralelo)

## 🚀 Endpoints

### 1. GET `/usuarios`
Obtiene la lista completa de 100 usuarios de prueba.

```bash
curl http://localhost:8080/usuarios
```

### 2. GET `/usuarios/:id`
Obtiene un usuario específico por su ID. Simula procesamiento con delays.

```bash
curl http://localhost:8080/usuarios/5
```

**Diferencia de rendimiento:**
- **Sin concurrencia**: ~1000ms (100 usuarios × 10ms)
- **Con Goroutines**: ~10-50ms (ejecución paralela)

### 3. GET `/usuarios/search?nombre=Juan`
Busca usuarios que contengan el nombre especificado en su campo nombre.

```bash
curl "http://localhost:8080/usuarios/search?nombre=Juan"
```

**Diferencia de rendimiento:**
- **Sin concurrencia**: ~500ms (100 usuarios × 5ms)
- **Con Goroutines**: ~5-20ms (ejecución paralela)

### 4. POST `/usuarios/process`
Procesa todos los usuarios de forma secuencial o paralela, retornando el tiempo total.

```bash
curl -X POST http://localhost:8080/usuarios/process
```

**Diferencia de rendimiento:**
- **Sin concurrencia**: ~5000ms (100 usuarios × 50ms)
- **Con Goroutines**: ~50-100ms (ejecución paralela)

## 🔀 Comparación de Ramas

### Rama `master` (Sin Concurrencia)
```go
func processUsuarios(c *gin.Context) {
    for range usuarios {
        time.Sleep(50 * time.Millisecond)  // Secuencial
        procesados++
    }
}
```

**Características:**
- Procesamiento secuencial
- Cada operación espera a que termine la anterior
- Bajo uso de recursos pero lento
- Ideal para operaciones que requieren orden garantizado

### Rama `concurrent` (Con Goroutines)
```go
func processUsuarios(c *gin.Context) {
    var wg sync.WaitGroup
    
    for range usuarios {
        wg.Add(1)
        go func() {
            defer wg.Done()
            time.Sleep(50 * time.Millisecond)  // Paralelo
            procesados++
        }()
    }
    
    wg.Wait()  // Esperar a que terminen todas
}
```

**Características:**
- Procesamiento paralelo con Goroutines
- Múltiples operaciones ejecutándose simultáneamente
- Mayor consumo de recursos pero mucho más rápido
- Sincronización con `sync.WaitGroup` y `sync.Mutex`

## 💻 Cómo usar

### Requisitos
- Go 1.16+
- Módulo: `github.com/gin-gonic/gin`

### Instalación de dependencias
```bash
go mod download
go mod tidy
```

### Ejecutar en rama `master` (sin concurrencia)
```bash
git checkout master
go run main.go
```

### Ejecutar en rama `concurrent` (con Goroutines)
```bash
git checkout concurrent
go run main.go
```

### Compilar ejecutable
```bash
go build -o api-go
./api-go
```

El servidor estará disponible en `http://localhost:8080`

## 📊 Resumen de Diferencias

| Aspecto | Sin Concurrencia | Con Goroutines |
|---------|-----------------|----------------|
| **Búsqueda por ID (100 usuarios)** | ~1000ms | ~50ms |
| **Búsqueda por nombre (100 usuarios)** | ~500ms | ~20ms |
| **Procesamiento (100 usuarios)** | ~5000ms | ~100ms |
| **Complejidad del código** | Simple | Requiere sync |
| **Uso de CPU** | Bajo/Uniforme | Alto/Variable |
| **Escalabilidad** | Limitada | Excelente |

## 🔑 Conceptos Clave

### Goroutines
- Unidades de concurrencia muy ligeras de Go
- No son threads del SO, Go las mapea inteligentemente
- Ideal para operaciones I/O y CPU-bound paralelo

### sync.WaitGroup
- Sincroniza múltiples Goroutines
- `Add()`: incrementa el contador
- `Done()`: decrementa el contador
- `Wait()`: espera a que el contador sea 0

### sync.Mutex
- Mutex (Mutual Exclusion) para proteger datos compartidos
- `Lock()`: adquiere el bloqueo
- `Unlock()`: libera el bloqueo

## 📝 Notas

- Los tiempos simulados (`time.Sleep`) son para demostración
- En producción, estos serían operaciones reales (BD, APIs, cálculos)
- La diferencia de rendimiento es mucho más notable con operaciones más lentas
- El overhead de crear muchas Goroutines es mínimo comparado con threads

## 🎯 Conclusión

Este proyecto demuestra claramente cómo Go's Goroutines permiten escribir código concurrente altamente eficiente. La rama `concurrent` ejecuta la misma lógica pero en paralelo, logrando speedups dramáticos sin complejidad excesiva.

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
go build -o api-go.exe main.go
./api-go.exe
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
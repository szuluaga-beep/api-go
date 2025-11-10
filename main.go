package main

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func init() {
	// Crear 100 usuarios de prueba
	nombres := []string{"Juan", "María", "Carlos", "Ana", "Pedro", "Luis", "Rosa", "Miguel", "Carmen", "Javier",
		"Isabel", "Francisco", "Dolores", "Manuel", "Angela", "Jose", "Elena", "Diego", "Pilar", "Antonio",
		"Beatriz", "Ramon", "Teresa", "Andres", "Marta", "Ricardo", "Francisca", "Fernando", "Victoria", "Enrique",
		"Susana", "Gabriel", "Margarita", "Raul", "Josefa", "Alberto", "Amparo", "Alfonso", "Concepcion", "Arturo",
		"Antonia", "Ruben", "Esperanza", "Salvador", "Consuelo", "Eduardo", "Ascension", "Emilio", "Gloria", "Felipe",
		"Rosario", "Alfredo", "Presentacion", "Guillermo", "Soledad", "Esteban", "Visitacion", "Ignacio", "Asuncion", "Julian"}

	apellidos := []string{"García", "López", "Rodríguez", "Martínez", "Pérez", "Hernández", "Sánchez", "González", "Torres", "Flores",
		"Reyes", "Ruiz", "Morales", "Rivera", "Gutierrez", "Ortiz", "Jimenez", "Diaz", "Cruz", "Vargas",
		"Castro", "Salazar", "Romero", "Aguilar", "Cabrera", "Campos", "Carvajal", "Castillo", "Contreras", "Cordero"}

	for i := 1; i <= 100; i++ {
		nombre := nombres[(i-1)%len(nombres)]
		apellido := apellidos[(i-1)%len(apellidos)]
		usuario := Usuario{
			ID:     i,
			Nombre: nombre + " " + apellido,
			Email:  nombre + "." + apellido + "@example.com",
		}
		usuarios = append(usuarios, usuario)
	}
}

func main() {
	router := gin.Default()

	// Configurar proxies de forma segura
	// Solo confiar en proxies específicos (en este caso ninguno en desarrollo)
	router.SetTrustedProxies([]string{})

	// Endpoints
	router.GET("/usuarios", getUsuarios)
	router.GET("/usuarios/:id", getUsuarioByID)
	router.GET("/usuarios/search", searchUsuarios)
	router.POST("/usuarios/process", processUsuarios)

	router.Run(":8080")
}

// GET /usuarios - Obtiene todos los usuarios
func getUsuarios(c *gin.Context) {
	c.JSON(http.StatusOK, usuarios)
}

// GET /usuarios/:id - Obtiene un usuario por ID con Goroutines
func getUsuarioByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	
	// Canal para recibir el resultado
	resultChan := make(chan *Usuario)
	var wg sync.WaitGroup

	// Lanzar Goroutines para buscar en paralelo
	for _, usuario := range usuarios {
		wg.Add(1)
		go func(u Usuario) {
			defer wg.Done()
			// Simular procesamiento con Goroutines
			time.Sleep(10 * time.Millisecond)
			if u.ID == id {
				resultChan <- &u
			}
		}(usuario)
	}

	// Goroutine para cerrar el canal cuando terminen todas
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Intentar recibir resultado
	for resultado := range resultChan {
		if resultado != nil {
			c.JSON(http.StatusOK, resultado)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "usuario no encontrado"})
}

// GET /usuarios/search?nombre=Juan - Busca usuarios por nombre con Goroutines
func searchUsuarios(c *gin.Context) {
	nombre := c.Query("nombre")
	if nombre == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "parámetro 'nombre' requerido"})
		return
	}

	var resultados []Usuario
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Lanzar Goroutines para búsqueda paralela
	for _, usuario := range usuarios {
		wg.Add(1)
		go func(u Usuario) {
			defer wg.Done()
			// Simular procesamiento de búsqueda con Goroutines
			time.Sleep(5 * time.Millisecond)
			if strings.Contains(strings.ToLower(u.Nombre), strings.ToLower(nombre)) {
				mu.Lock()
				resultados = append(resultados, u)
				mu.Unlock()
			}
		}(usuario)
	}

	// Esperar a que terminen todas las Goroutines
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"total":    len(resultados),
		"usuarios": resultados,
	})
}

// POST /usuarios/process - Procesa usuarios con Goroutines (paralelo)
func processUsuarios(c *gin.Context) {
	inicio := time.Now()
	var wg sync.WaitGroup
	procesados := 0
	var mu sync.Mutex

	// Lanzar Goroutines para procesamiento paralelo
	for range usuarios {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Simular procesamiento pesado con Goroutines
			time.Sleep(50 * time.Millisecond)
			mu.Lock()
			procesados++
			mu.Unlock()
		}()
	}

	// Esperar a que terminen todas las Goroutines
	wg.Wait()

	duracion := time.Since(inicio)
	c.JSON(http.StatusOK, gin.H{
		"mensaje":           "Procesamiento completado",
		"total_procesados":  procesados,
		"duracion_ms":       duracion.Milliseconds(),
		"modo":              "CONCURRENTE (con Goroutines)",
	})
}

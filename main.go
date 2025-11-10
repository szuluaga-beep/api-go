package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID     int       `json:"id"`
	Nombre string    `json:"nombre"`
	Email  string    `json:"email"`
	Creado time.Time `json:"creado"`
}

type UsuarioService struct {
	usuarios []Usuario
	mutex    sync.RWMutex
	nextID   int
}

var usuarioService *UsuarioService

func init() {
	usuarioService = &UsuarioService{
		usuarios: make([]Usuario, 0),
		nextID:   1,
	}

	// Crear 100 usuarios de prueba de forma síncrona
	nombres := []string{"Juan", "María", "Carlos", "Ana", "Pedro", "Luis", "Rosa", "Miguel", "Carmen", "Javier",
		"Isabel", "Francisco", "Dolores", "Manuel", "Angela", "Jose", "Elena", "Diego", "Pilar", "Antonio",
		"Beatriz", "Ramon", "Teresa", "Andres", "Marta", "Ricardo", "Francisca", "Fernando", "Victoria", "Enrique",
		"Susana", "Gabriel", "Margarita", "Raul", "Josefa", "Alberto", "Amparo", "Alfonso", "Concepcion", "Arturo",
		"Antonia", "Ruben", "Esperanza", "Salvador", "Consuelo", "Eduardo", "Ascension", "Emilio", "Gloria", "Felipe",
		"Rosario", "Alfredo", "Presentacion", "Guillermo", "Soledad", "Esteban", "Visitacion", "Ignacio", "Asuncion", "Julian"}

	apellidos := []string{"García", "López", "Rodríguez", "Martínez", "Pérez", "Hernández", "Sánchez", "González", "Torres", "Flores",
		"Reyes", "Ruiz", "Morales", "Rivera", "Gutierrez", "Ortiz", "Jimenez", "Diaz", "Cruz", "Vargas",
		"Castro", "Salazar", "Romero", "Aguilar", "Cabrera", "Campos", "Carvajal", "Castillo", "Contreras", "Cordero"}

	for i := 0; i < 100; i++ {
		nombre := nombres[i%len(nombres)]
		apellido := apellidos[i%len(apellidos)]
		usuario := Usuario{
			Nombre: nombre + " " + apellido + " " + string(rune(i+1)),
			Email:  nombre + "." + apellido + string(rune(i+1)) + "@example.com",
		}
		usuarioService.CrearUsuario(usuario)
	}
	log.Println("100 Usuarios de prueba creados")
}

func main() {
	router := gin.Default()

	// Middleware para logging asíncrono
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// Log asíncrono
		go func() {
			latency := time.Since(start)
			log.Printf("[%s] %s %d - %v",
				c.Request.Method,
				c.Request.URL.Path,
				c.Writer.Status(),
				latency)
		}()
	})

	router.GET("/usuarios", getUsuarios)

	log.Println("Servidor iniciando en puerto 8080...")
	router.Run(":8080")
}

// Métodos del UsuarioService con goroutines
func (us *UsuarioService) CrearUsuario(usuario Usuario) Usuario {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	usuario.ID = us.nextID
	usuario.Creado = time.Now()
	us.nextID++
	us.usuarios = append(us.usuarios, usuario)
	return usuario
}

func (us *UsuarioService) ObtenerUsuarios() []Usuario {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	// Crear copia para evitar modificaciones concurrentes
	usuarios := make([]Usuario, len(us.usuarios))
	copy(usuarios, us.usuarios)
	return usuarios
}

func getUsuarios(c *gin.Context) {
	// Usar goroutine para procesamiento asíncrono si es necesario
	resultChan := make(chan []Usuario)

	go func() {
		usuarios := usuarioService.ObtenerUsuarios()
		resultChan <- usuarios
	}()

	usuarios := <-resultChan
	c.JSON(http.StatusOK, usuarios)
}

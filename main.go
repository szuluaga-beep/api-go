package main

import (
	"log"
	"net/http"
	"strconv"
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
}

func main() {
	// Inicializar algunos usuarios de ejemplo en una goroutine
	go func() {
		usuariosEjemplo := []Usuario{
			{Nombre: "Juan Pérez", Email: "juan@example.com"},
			{Nombre: "María García", Email: "maria@example.com"},
			{Nombre: "Carlos López", Email: "carlos@example.com"},
		}

		for _, usuario := range usuariosEjemplo {
			usuarioService.CrearUsuario(usuario)
			time.Sleep(100 * time.Millisecond) // Simular procesamiento
		}
		log.Println("Usuarios de ejemplo creados")
	}()

	router := gin.Default()

	// Middleware para logging asíncrono
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// Log asíncrono
		go func() {
			latency := time.Since(start)
			log.Printf("[%s] %s %s - %v",
				c.Request.Method,
				c.Request.URL.Path,
				c.Writer.Status(),
				latency)
		}()
	})

	router.GET("/usuarios", getUsuarios)
	router.GET("/usuarios/:id", getUsuario)
	router.POST("/usuarios", crearUsuario)
	router.PUT("/usuarios/:id", actualizarUsuario)
	router.DELETE("/usuarios/:id", eliminarUsuario)
	router.GET("/usuarios/stats", getEstadisticas)

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

func (us *UsuarioService) ObtenerUsuario(id int) (Usuario, bool) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	for _, usuario := range us.usuarios {
		if usuario.ID == id {
			return usuario, true
		}
	}
	return Usuario{}, false
}

func (us *UsuarioService) ActualizarUsuario(id int, usuarioActualizado Usuario) (Usuario, bool) {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	for i, u := range us.usuarios {
		if u.ID == id {
			usuarioActualizado.ID = id
			usuarioActualizado.Creado = u.Creado // Mantener fecha de creación original
			us.usuarios[i] = usuarioActualizado
			return usuarioActualizado, true
		}
	}
	return Usuario{}, false
}

func (us *UsuarioService) EliminarUsuario(id int) bool {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	for i, usuario := range us.usuarios {
		if usuario.ID == id {
			us.usuarios = append(us.usuarios[:i], us.usuarios[i+1:]...)
			return true
		}
	}
	return false
}

func (us *UsuarioService) ObtenerEstadisticas() map[string]interface{} {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	totalUsuarios := len(us.usuarios)
	var usuarioMasReciente Usuario

	if totalUsuarios > 0 {
		usuarioMasReciente = us.usuarios[totalUsuarios-1]
	}

	return map[string]interface{}{
		"total_usuarios":       totalUsuarios,
		"usuario_mas_reciente": usuarioMasReciente,
		"timestamp":            time.Now(),
	}
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

func getUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Usar goroutine para búsqueda asíncrona
	resultChan := make(chan Usuario)
	errorChan := make(chan bool)

	go func() {
		if usuario, found := usuarioService.ObtenerUsuario(id); found {
			resultChan <- usuario
		} else {
			errorChan <- true
		}
	}()

	select {
	case usuario := <-resultChan:
		c.JSON(http.StatusOK, usuario)
	case <-errorChan:
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
	}
}

func crearUsuario(c *gin.Context) {
	var usuario Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear usuario en goroutine
	resultChan := make(chan Usuario)

	go func() {
		nuevoUsuario := usuarioService.CrearUsuario(usuario)
		resultChan <- nuevoUsuario
	}()

	usuarioCreado := <-resultChan
	c.JSON(http.StatusCreated, usuarioCreado)
}

func actualizarUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var usuario Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Actualizar usuario en goroutine
	resultChan := make(chan Usuario)
	errorChan := make(chan bool)

	go func() {
		if usuarioActualizado, found := usuarioService.ActualizarUsuario(id, usuario); found {
			resultChan <- usuarioActualizado
		} else {
			errorChan <- true
		}
	}()

	select {
	case usuarioActualizado := <-resultChan:
		c.JSON(http.StatusOK, usuarioActualizado)
	case <-errorChan:
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
	}
}

func eliminarUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Eliminar usuario en goroutine
	resultChan := make(chan bool)

	go func() {
		eliminado := usuarioService.EliminarUsuario(id)
		resultChan <- eliminado
	}()

	if <-resultChan {
		c.JSON(http.StatusOK, gin.H{"mensaje": "Usuario eliminado correctamente"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
	}
}

func getEstadisticas(c *gin.Context) {
	// Obtener estadísticas en goroutine
	resultChan := make(chan map[string]interface{})

	go func() {
		stats := usuarioService.ObtenerEstadisticas()
		resultChan <- stats
	}()

	estadisticas := <-resultChan
	c.JSON(http.StatusOK, estadisticas)
}

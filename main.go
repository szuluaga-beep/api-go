package main

import (
	"net/http"

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

	router.GET("/usuarios", getUsuarios)

	router.Run(":8080")
}

func getUsuarios(c *gin.Context) {
	c.JSON(http.StatusOK, usuarios)
}

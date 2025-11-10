package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Usuario struct {
	ID    int    `json:"id"`
	Nombre string `json:"nombre"`
	Email string `json:"email"`
}

var usuarios []Usuario

func main() {
	router := gin.Default()

	router.GET("/usuarios", getUsuarios)
	router.GET("/usuarios/:id", getUsuario)
	router.POST("/usuarios", crearUsuario)
	router.PUT("/usuarios/:id", actualizarUsuario)
	router.DELETE("/usuarios/:id", eliminarUsuario)

	router.Run(":8080")
}

func getUsuarios(c *gin.Context) {
	c.JSON(http.StatusOK, usuarios)
}

func getUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, usuario := range usuarios {
		if usuario.ID == id {
			c.JSON(http.StatusOK, usuario)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "No encontrado"})
}

func crearUsuario(c *gin.Context) {
	var usuario Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usuario.ID = len(usuarios) + 1
	usuarios = append(usuarios, usuario)
	c.JSON(http.StatusCreated, usuario)
}

func actualizarUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var usuario Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, u := range usuarios {
		if u.ID == id {
			usuario.ID = id
			usuarios[i] = usuario
			c.JSON(http.StatusOK, usuario)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "No encontrado"})
}

func eliminarUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, usuario := range usuarios {
		if usuario.ID == id {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"mensaje": "Eliminado"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "No encontrado"})
}
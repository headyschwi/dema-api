package controllers

import (
	"dema-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCurso(c *gin.Context) {
	var Curso models.Curso
	err := c.BindJSON(&Curso)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error(),
		})
	}
	models.Database.Create(&Curso)
	c.JSON(http.StatusOK, Curso)
}
func ReadCursos(c *gin.Context) {
	var Cursos []models.Curso
	models.Database.Find(&Cursos)
	if Cursos == nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Cursos not Found"})
		return
	}
	c.JSON(http.StatusOK, Cursos)
}
func ReadOneCurso(c *gin.Context) {
	var Curso models.Curso
	id := c.Param("id")

	models.Database.First(&Curso, id)
	if Curso.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Curso not Found"})
		return
	}

	c.JSON(http.StatusOK, Curso)
}

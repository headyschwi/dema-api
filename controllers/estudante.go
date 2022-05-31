package controllers

import (
	"dema-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CreateEstudante(c *gin.Context) {
	var Estudante models.Estudante
	err := c.BindJSON(&Estudante)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error:": err.Error(),
		})
	}
	models.Database.Create(&Estudante)
	c.JSON(http.StatusOK, Estudante)
}
func ReadEstudantes(c *gin.Context) {
	var Estudantes []models.Estudante
	models.Database.Find(&Estudantes)
	if Estudantes == nil {
		c.JSON(http.StatusNotFound, gin.H{"Error:": "Estudantes not Found"})
		return
	}

	models.Database.Preload(clause.Associations).Find(&Estudantes)
	c.JSON(http.StatusOK, Estudantes)
}
func ReadOneEstudante(c *gin.Context) {
	var Estudante models.Estudante
	id := c.Param("id")

	models.Database.Preload(clause.Associations).First(&Estudante, id)
	if Estudante.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Error:": "Estudante not Found"})
		return
	}
	c.JSON(http.StatusOK, Estudante)
}
func AddCursoToEstudante(c *gin.Context) {
	var Estudante models.Estudante
	var Curso models.Curso

	idEstudante := c.Param("ide")
	models.Database.Preload("CursosFinalizados").First(&Estudante, idEstudante)

	if Estudante.ID == 0 {
		return
	}

	idCurso := c.Param("idc")
	models.Database.First(&Curso, idCurso)
	if Curso.ID == 0 {
		return
	}

	models.Database.Model(&Estudante).Association("CursosFinalizados").Append(&Curso)
	c.JSON(http.StatusOK, Estudante)
}

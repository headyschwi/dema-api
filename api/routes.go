package api

import (
	"dema-api/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	listen(800)
}

func listen(p int) {
	port := fmt.Sprintf(":%d", p)
	r := gin.Default()
	r.POST("/estudantes", controllers.CreateEstudante)
	r.GET("/estudantes", controllers.ReadEstudantes)
	r.GET("/estudantes/:id", controllers.ReadOneEstudante)

	r.PATCH("/estudantes/:ide/cursos/add/:idc", controllers.AddCursoToEstudante)

	r.POST("/cursos", controllers.CreateCurso)
	r.GET("/cursos", controllers.ReadCursos)
	r.GET("/cursos/:id", controllers.ReadOneCurso)
	r.Run(port)
}

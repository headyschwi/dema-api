package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	USER     = "postgres"
	PASSWORD = "root"
	DBNAME   = "dema-api"
	PORT     = 5432
)

var (
	Database *gorm.DB
	err      error
)

func DatabaseConnect() {
	connectString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", HOST, USER, PASSWORD, DBNAME, PORT)
	Database, err = gorm.Open(postgres.Open(connectString))
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println("Conectado com o banco de dados!")
}

func AutoMigrations() {
	if Database == nil {
		log.Fatal("Conecte-se ao banco da dados para utilizar esta função!")
		return
	}
	Database.Debug().Migrator().DropTable(&Estudante{}, &Curso{}, "estudantes_cursos")
	Database.Debug().AutoMigrate(&Estudante{}, &Curso{})
}

package models

import "gorm.io/gorm"

type Estudante struct {
	gorm.Model
	Nome              string   `gorm:"not null" json:"nome" `
	Exp               int      `gorm:"default:0;" json:"exp"`
	Email             string   `gorm:"not null; uniqueIndex" json:"email"`
	Password          string   `gorm:"not null" json:"password"`
	CursosFinalizados []*Curso `gorm:"many2many:estudantes_cursos; constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"cursosFinalizados"`
}

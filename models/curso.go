package models

import (
	"gorm.io/gorm"
)

type Curso struct {
	gorm.Model
	Nome           string `gorm:"not null; uniqueIndex" json:"nome"`
	DuracaoMinutos int    `gorm:"not null" json:"duracao_minutos"`
	Exp            int    `gorm:"not null" json:"exp"`
}

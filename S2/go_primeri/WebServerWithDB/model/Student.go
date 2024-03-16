package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name" gorm:"not null;type:string"`
	Major string    `json:"major"`
}

func (student *Student) BeforeCreate(scope *gorm.DB) error {
	student.ID = uuid.New()
	return nil
}

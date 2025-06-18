package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type StudentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *StudentRepository) FindById(id string) (model.Student, error) {
	student := model.Student{}
	dbResult := repo.DatabaseConnection.First(&student, "id = ?", id)
	if dbResult != nil {
		return student, dbResult.Error
	}
	return student, nil
}

func (repo *StudentRepository) CreateStudent(student *model.Student) error {
	dbResult := repo.DatabaseConnection.Create(student)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

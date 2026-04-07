package db

import (
	"github.com/GabsOnRails/api-estudantes/db/schemas"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to connect to database: %v", err.Error())
	}
	db.AutoMigrate(&schemas.Student{})
	return db, nil
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student schemas.Student) error {
	if result := s.DB.Create(&student); result.Error != nil {
		log.Error().Msgf("Failed to add student: %v", result.Error)
		return result.Error
	}

	log.Info().Msgf("Student %s added successfully", student.Name)
	return nil
}

func (s *StudentHandler) GetStudent(id int) (schemas.Student, error) {
	var student schemas.Student
	err := s.DB.First(&student, id)
	// err := s.DB.Find(&students).Error
	// if err != nil {
	// 	return nil, err
	// }
	return student, err.Error
}

func (s *StudentHandler) GetStudents() ([]schemas.Student, error) {
	students := []schemas.Student{}
	err := s.DB.Find(&students).Error
	return students, err

}

func (s *StudentHandler) GetFilteredStudents(active bool) ([]schemas.Student, error) {
	filteredStudents := []schemas.Student{}
	err := s.DB.Where("active = ?", active).Find(&filteredStudents).Error
	return filteredStudents, err
}

func (s *StudentHandler) UpdateStudent(updatingStudent schemas.Student) error {
	return s.DB.Save(&updatingStudent).Error
}

func (s *StudentHandler) DeleteStudent(student schemas.Student) error {
	return s.DB.Delete(&student).Error
}

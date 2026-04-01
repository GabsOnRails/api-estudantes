package db

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

func Init() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to connect to database: %v", err.Error())
	}
	db.AutoMigrate(&Student{})
	return db, nil
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student Student) error {
	if result := s.DB.Create(&student); result.Error != nil {
		log.Error().Msgf("Failed to add student: %v", result.Error)
		return result.Error
	}

	log.Info().Msgf("Student %s added successfully", student.Name)
	return nil
}

func (s *StudentHandler) GetStudent(id int) (Student, error) {
	var student Student
	err := s.DB.First(&student, id)
	// err := s.DB.Find(&students).Error
	// if err != nil {
	// 	return nil, err
	// }
	return student, err.Error
}

func (s *StudentHandler) GetStudents() ([]Student, error) {
	students := []Student{}
	err := s.DB.Find(&students).Error
	return students, err

}

func (s *StudentHandler) UpdateStudent(updatingStudent Student) error {
	return s.DB.Save(&updatingStudent).Error
}

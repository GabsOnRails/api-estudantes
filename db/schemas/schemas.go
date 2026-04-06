package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

type StudentResponse struct {
	gorm.Model
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
	DeleteAt  time.Time `json:"deleted_at"`
	Name      string    `json:"name"`
	CPF       int       `json:"cpf"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Active    bool      `json:"active"`
}

func NewResponseStudent(students []Student) []StudentResponse {
	studentsResponse := []StudentResponse{}
	for _, student := range students {
		studentResponse := StudentResponse{
			ID:        int(student.ID),
			CreatedAt: student.CreatedAt,
			UpdateAt:  student.UpdatedAt,
			Name:      student.Name,
			CPF:       student.CPF,
			Email:     student.Email,
			Age:       student.Age,
			Active:    student.Active,
		}
		studentsResponse = append(studentsResponse, studentResponse)
	}
	return studentsResponse
}

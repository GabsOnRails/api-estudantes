package api

import (
	"fmt"
)

type StudentRequest struct {
	Name   string `json:"name"`
	CPF    int    `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active *bool  `json:"active"` // Use pointer to distinguish between false and not provided
}

func errParamRequired(param, typ string) error {
	return fmt.Errorf("%s is required and must be a valid %s", param, typ)
}

func (s *StudentRequest) ValidateStudentRequest(student StudentRequest) error {
	if s.Name == "" {
		return errParamRequired("name", "string")
	}
	if s.CPF == 0 {
		return errParamRequired("cpf", "integer")
	}
	if s.Email == "" {
		return errParamRequired("email", "string")
	}
	if s.Age == 0 {
		return errParamRequired("age", "integer")
	}
	if s.Active == nil {
		return errParamRequired("active", "boolean")
	}
	return nil
}

package api

import (
	"net/http"
	"strconv"

	"github.com/GabsOnRails/api-estudantes/db/schemas"
	"github.com/labstack/echo/v5"
	"github.com/rs/zerolog/log"
)

//----------
// Handlers
//----------

func (api *API) createStudent(c *echo.Context) error {
	studentReq := StudentRequest{}

	if err := c.Bind(&studentReq); err != nil {
		return err
	}
	if err := studentReq.ValidateStudentRequest(studentReq); err != nil {
		log.Error().Err(err).Msg("Validation error")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}
	student := schemas.Student{
		Name:   studentReq.Name,
		CPF:    studentReq.CPF,
		Email:  studentReq.Email,
		Age:    studentReq.Age,
		Active: *studentReq.Active,
	}

	err := api.DB.AddStudent(student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error to create student",
		})
	}

	return c.JSON(http.StatusCreated, student)
}

func (api *API) getStudent(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid student ID",
		})
	}

	student, err := api.DB.GetStudent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error to get student",
		})
	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) updateStudent(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid student ID",
		})
	}

	receivedStudent := schemas.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		return err
	}

	updatingStudent, err := api.DB.GetStudent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error to get student",
		})
	}
	updatingStudent = updateStudentInfo(receivedStudent, updatingStudent)
	if err := api.DB.UpdateStudent(updatingStudent); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error to update student",
		})
	}

	return c.JSON(http.StatusOK, updatingStudent)

}

func updateStudentInfo(receivedStudent, updatingStudent schemas.Student) schemas.Student {
	if receivedStudent.Name != "" {
		updatingStudent.Name = receivedStudent.Name
	}
	if receivedStudent.CPF != 0 {
		updatingStudent.CPF = receivedStudent.CPF
	}
	if receivedStudent.Email != "" {
		updatingStudent.Email = receivedStudent.Email
	}
	if receivedStudent.Age != 0 {
		updatingStudent.Age = receivedStudent.Age
	}
	if receivedStudent.Active != updatingStudent.Active {
		updatingStudent.Active = receivedStudent.Active
	}
	return updatingStudent
}

func (api *API) deleteStudent(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid student ID",
		})
	}
	student, err := api.DB.GetStudent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Student not found",
		})
	}

	if err := api.DB.DeleteStudent(student); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete student",
		})
	}

	return c.JSON(http.StatusOK, student)

}

func (api *API) getStudents(c *echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error to get students",
		})
	}

	listOfStudents := map[string][]schemas.StudentResponse{
		"students": schemas.NewResponseStudent(students),
	}
	return c.JSON(http.StatusOK, listOfStudents)
}

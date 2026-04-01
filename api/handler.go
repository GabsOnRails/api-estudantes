package api

import (
	"net/http"
	"strconv"

	"github.com/GabsOnRails/api-estudantes/db"
	"github.com/labstack/echo/v5"
)

//----------
// Handlers
//----------

func (api *API) createStudent(c *echo.Context) error {
	student := db.Student{}

	if err := c.Bind(&student); err != nil {
		return err
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
	lock.Lock()
	defer lock.Unlock()
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	students[id].Name = u.Name
	return c.JSON(http.StatusOK, students[id])
}

func (api *API) deleteStudent(c *echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(students, id)
	return c.NoContent(http.StatusNoContent)
}

func (api *API) getStudents(c *echo.Context) error {
	students, err := api.DB.GetStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error to get students",
		})
	}
	return c.JSON(http.StatusOK, students)
}

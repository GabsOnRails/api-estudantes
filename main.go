package main

import (
	"context"
	"net/http"
	"strconv"
	"sync"

	"github.com/GabsOnRails/api-estudantes/db"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	students = map[int]*user{}
	seq      = 1
	lock     = sync.Mutex{}
)

//----------
// Handlers
//----------

func createStudent(c *echo.Context) error {
	student := db.Student{}

	if err := c.Bind(&student); err != nil {
		return err
	}

	err := db.AddStudent(student)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error to create student",
		})
	}

	return c.JSON(http.StatusCreated, student)
}

func getStudent(c *echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, students[id])
}

func updateStudent(c *echo.Context) error {
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

func deleteStudent(c *echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(students, id)
	return c.NoContent(http.StatusNoContent)
}

func getStudents(c *echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, students)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudents)
	e.POST("/students", createStudent)
	e.GET("/students/:id", getStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	// Start server
	sc := echo.StartConfig{Address: ":8080"}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

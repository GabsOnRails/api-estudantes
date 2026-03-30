package api

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/GabsOnRails/api-estudantes/db"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

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

func NewServer() *API {

	e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	database, err := db.Init()
	if err != nil {
		panic(err)
	}

	return &API{
		Echo: e,
		DB:   database,
	}

}

func (api *API) StartServer() error {
	// Start server
	return api.Echo.Start(":8080")
}

func (api *API) ConfigureRoutes() {
	// Routes
	api.Echo.GET("/students", getStudents)
	api.Echo.POST("/students", createStudent)
	api.Echo.GET("/students/:id", getStudent)
	api.Echo.PUT("/students/:id", updateStudent)
	api.Echo.DELETE("/students/:id", deleteStudent)
}

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
	students, err := db.GetStudents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error to get students",
		})
	}
	return c.JSON(http.StatusOK, students)
}

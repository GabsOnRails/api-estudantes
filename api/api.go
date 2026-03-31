package api

import (
	"sync"

	"github.com/GabsOnRails/api-estudantes/db"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
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
	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}

}

func (api *API) StartServer() error {
	// Start server
	return api.Echo.Start(":8080")
}

func (api *API) ConfigureRoutes() {
	// Routes
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/students/:id", api.getStudent)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
}

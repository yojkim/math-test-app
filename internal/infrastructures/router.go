package infrastructures

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yojkim/math-test-app/internal/infrastructures/sqlite"
	"github.com/yojkim/math-test-app/internal/services"
)

var Router *echo.Echo

func init() {
	router := echo.New()

	router.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	apiGroup := router.Group("/api")

	conn := sqlite.Connect()

	problemService := services.NewProblemService(conn)
	resultService := services.NewResultService(conn)

	apiGroup.POST("/problems", func(c echo.Context) error {
		return problemService.CreateProblems(c)
	})

	apiGroup.GET("/fetchProblem", func(c echo.Context) error {
		return problemService.GetAllProblems(c)
	})

	apiGroup.POST("/submit", func(c echo.Context) error {
		return resultService.CheckAnswer(c)
	})

	Router = router
}

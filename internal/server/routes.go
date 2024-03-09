package server

import (
	"net/http"

	"hyperbot/cmd/web"
	webHandlers "hyperbot/cmd/web/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", s.healthHandler)
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/js/*", echo.WrapHandler(fileServer))
	// web routes
	// e.GET("/", s.HelloWorldHandler)
	e.GET("/dashboard", webHandlers.RenderDashboardLandingPage)
	e.GET("/login", webHandlers.RenderLoginPage)

	return e
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.Health())
}

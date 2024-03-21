package main

import (
	"fmt"
	"log"
	"net/http"

	"hyperbot/configs"
	"hyperbot/web"
	webHandlers "hyperbot/web/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found", err)
	}
	var app = echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	configs.ConnectDB()

	fileServer := http.FileServer(http.FS(web.Files))
	app.GET("/js/*", echo.WrapHandler(fileServer))
	// web routes
	app.GET("/", webHandlers.RenderLandingPage)
	app.GET("/dashboard", webHandlers.RenderDashboardLandingPage)
	app.GET("/login", webHandlers.RenderLoginPage)

	app.Static("/static", "assets")

	port := fmt.Sprintf(":%s", configs.EnvPort())

	log.Fatal(app.Start(port))
}

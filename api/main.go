package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Kiraks200/url_shortning/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)

}

func main() {
	err := godotenv.Load()  
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()  // lets try fibre but gorilla mux is normally prefered

	app.Use(logger.New())

	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

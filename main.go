package main

import (
	"auth/injector"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	app := fiber.New()
	injector.NewAdminInjector(app)
	injector.NewAuthInjector(app)

	fmt.Printf("⚡️[server]: Server is running on port %s\n", port)
	log.Fatal(app.Listen(port))
}

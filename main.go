package main

import (
	"auth/config"
	"auth/injector"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	conns := config.NewConn()
	user := conns.UserConn()
	customer := conns.CustomerConn()

	port := os.Getenv("PORT")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		ExposeHeaders:    "Retry-After",
	}))

	injector.NewAuthInjector(app)
	// some routes use jwt here
	injector.NewAdminInjector(app, user)

	// customer
	api := app.Group("api")
	injector.NewCustomerInjector(api, customer)

	fmt.Printf("⚡️[server]: Server is running on port %s\n", port)
	log.Fatal(app.Listen(port))
}

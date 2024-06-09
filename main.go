package main

import (
	"go-fiber-swagger/configuration"
	ds "go-fiber-swagger/domain/datasources"
	repo "go-fiber-swagger/domain/repositories"
	gw "go-fiber-swagger/src/gateways"
	"go-fiber-swagger/src/middlewares"
	sv "go-fiber-swagger/src/services"
	"log"
	"os"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	// // // remove this before deploy but when developing uncomment ###################
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// /// ############################################

	app := fiber.New(configuration.NewFiberConfiguration())
	middlewares.Logger(app)
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(swagger.New(swagger.Config{
		BasePath: "/api/users/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}))

	mongodb := ds.NewMongoDB(10)

	userMongo := repo.NewUsersRepository(mongodb)

	sv0 := sv.NewUsersService(userMongo)

	gw.NewHTTPGateway(app, sv0)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}

package main

import (
	"encoding/json"
	"log"
	"os"

	db "go-template/data/database"
	repo "go-template/data/repository"
	"go-template/src/gateway"
	"go-template/src/middleware"
	service "go-template/src/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "go-template/docs" 
)
// @title Fiber API MyFavFood
// @version 1.0
// @description This is a api MyFavFood
// @host  humorous-colt-vital.ngrok-free.app
// @BasePath /
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	middleware.LoggerMiddleware(app)
	app.Get("/swagger/*", swagger.HandlerDefault)	
	postgresConfig := db.NewPSQL()
	pg, err := postgresConfig.ConnectGorm()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repo.NewUserRepository(pg)
	logsRepository := repo.NewLogsRepository(pg)
	sv1 := service.NewUserService(userRepository, logsRepository)
	loginSV := service.NewLoginService(userRepository, logsRepository)
	cashcontrolSV := service.NewCashControlService(repo.NewCashControlRepository(pg))
	gateway.HTTPGatewayHandler(app, sv1, loginSV, cashcontrolSV)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set")
	}

	app.Listen(":" + port)
}

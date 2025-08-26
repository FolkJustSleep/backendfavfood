package gateway

import (
	"github.com/gofiber/fiber/v2"
	"go-template/util"
	"go-template/src/middleware"
)

func gatewayUser(gateway HTTPGateway, app *fiber.App){
	api := app.Group("/api/user")
	api.Get("/getall", gateway.GetAllUser)
	api.Post("/create", gateway.CreateUser)
	api.Get("/get", gateway.GetUserByID)
	api.Put("/update", middleware.CheckRole, gateway.UpdateUser)
	api.Delete("/delete", middleware.CheckRole, gateway.DeleteUser)
	api.Get("/ip", util.GetIP)
}

func gatewayLogin(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/login")
	api.Post("/login", gateway.Login)
}

func gatewayCashControl(gateway HTTPGateway, app *fiber.App){
	api := app.Group("/api/cashcontrol")
	api.Get("/getall", gateway.GetAllCashControls)
	api.Post("/create", gateway.CreateCashControl)
	api.Get("/get", gateway.GetCashControlByID)
	api.Put("/update", middleware.CheckRole, gateway.UpdateCashControl)
	api.Delete("/delete", middleware.CheckRole, gateway.DeleteCashControl)
}
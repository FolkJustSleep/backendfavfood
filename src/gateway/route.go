package gateway

import (
	"github.com/gofiber/fiber/v2"
	"go-template/util"
	"go-template/src/middleware"
)

func gatewayUser(gateway HTTPGateway, app *fiber.App){
	api := app.Group("/api/user")
	api.Get("/getall", gateway.GetAllUser)
	api.Get("/get", gateway.GetUserByID)
	api.Put("/update", middleware.CheckRole, gateway.UpdateUser)
	api.Delete("/delete", middleware.CheckRole, gateway.DeleteUser)
	api.Get("/ip", util.GetIP)
}

func gatewayLogin(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/auth")
	api.Post("/register", gateway.Register)
	api.Post("/login", gateway.Login)
	api.Post("/logout", gateway.Logout)
}

func gatewayCashControl(gateway HTTPGateway, app *fiber.App){
	api := app.Group("/api/cashcontrol")
	api.Get("/getall", gateway.GetAllCashControls)
	api.Post("/create", gateway.CreateCashControl)
	api.Get("/get", gateway.GetCashControlByID)
	api.Get("/getbyuserid", gateway.GetCashControlByUserID)
	api.Put("/update", middleware.CheckRole, gateway.UpdateCashControl)
	api.Delete("/delete", middleware.CheckRole, gateway.DeleteCashControl)
}

func gatewayMenu(gateway HTTPGateway, app *fiber.App){
	api := app.Group("/api/menu")
	api.Get("/getall", gateway.GetAllMenus)
	api.Post("/create", middleware.CheckRole, gateway.CreateMenu)
	api.Get("/get", gateway.GetMenuByID)
	api.Put("/update", middleware.CheckRole, gateway.UpdateMenu)
	api.Delete("/delete", middleware.CheckRole, gateway.DeleteMenu)
}

func gatewayTimestamp(gateway HTTPGateway, app *fiber.App){
	api := app.Group("/api/timestamp")
	api.Post("/checkin", gateway.CheckIn)
	api.Post("/checkout", gateway.CheckOut)
	api.Get("/checkedinemployee", gateway.GetCheckedInEmployee)
}
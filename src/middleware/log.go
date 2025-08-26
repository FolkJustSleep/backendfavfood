package middleware

import (

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerMiddleware(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		TimeFormat: "15:04:05",//set time format in log
		TimeZone: "Asia/Bangkok",// set timezone in log
	}))
}

// 11:00:11 | 500 |      732.25µs | 127.0.0.1 | GET | /api/user | - format log use in main 
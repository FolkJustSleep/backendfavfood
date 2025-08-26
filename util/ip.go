package util

import(
	"github.com/gofiber/fiber/v2"
	"go-template/data/model"
)

func GetIP(ctx *fiber.Ctx) error {
	ip := ctx.IP()
	return ctx.JSON(model.Response{
		Status: 200,
		Message: "Successfully get ip",
		Data: ip,
	})
}

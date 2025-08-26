package gateway

import (
	"time"

	"go-template/data/model"

	"github.com/gofiber/fiber/v2"
	
	fiberlog "github.com/gofiber/fiber/v2/log"
)

// Login Godoc
// @Summary Login a user
// @Description Login a user
// @Tags Login
// @Accept json
// @Produce json
// @Param login body model.LoginRequest true "Login credentials"
// @Success 200 {object} model.Response "Successfully logged in"
// @Failure 400 {object} model.Response "Bad Request"
// @Failure 401 {object} model.Response "Unauthorized"
// @Router /api/login/login [post]
func (h *HTTPGateway) Login(ctx *fiber.Ctx) error {
	var loginData model.LoginRequest
	if err := ctx.BodyParser(&loginData); err != nil {
		fiberlog.Error("Error parsing login request: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: 400,
			Message: "Invalid request",
			Data: nil,
		})
	}
	token, err := h.LoginService.Login(loginData.Email, loginData.Password)
	if err != nil {
		fiberlog.Error("Error logging in: ", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: "Unauthorized",
			Data: nil,
		})
	}
	ctx.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return ctx.JSON(model.Response{
		Status: 200,
		Message: "Login successful",
		Data: token,
	})
}
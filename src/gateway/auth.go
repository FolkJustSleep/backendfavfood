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
	token, err := h.AuthService.Login(loginData.Email, loginData.Password)
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
		Expires:  time.Now().Add(time.Hour * 2),
		HTTPOnly: true,
	})
	return ctx.JSON(model.Response{
		Status: 200,
		Message: "Login successful",
		Data: token,
	})
}

// Register Godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User object"
// @Success 201 {object} model.Response{data=model.User} "Successfully created user"
// @Failure 400 {object} model.Response "Bad Request"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/auth/register [post]
func (h *HTTPGateway) Register(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		fiberlog.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: 400,
			Message: "Bad Request",
			Data: nil,
		})
	}
	if user.Email == "" || user.Password == "" || user.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: 400,
			Message: "Email, Password and Role are required",
			Data: nil,
		})
	}
	createdUser, err := h.AuthService.Register(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: err.Error(),
			Data: nil,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Status: 201,
		Message: "Successfully created user",
		Data: createdUser,
	})
}

func (h *HTTPGateway) Logout(c *fiber.Ctx) error {
	if c == nil {
		return fiber.ErrBadRequest
	}
	err := h.AuthService.Logout(c)
	if err != nil {
		fiberlog.Error("Error logging out: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: "Internal Server Error" + err.Error(),
			Data: nil,
		})
	}
	return c.JSON(model.Response{
		Status: 200,
		Message: "Successfully logged out",
		Data: nil,
	})
}
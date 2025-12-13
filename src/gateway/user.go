package gateway

import (
	"go-template/data/model"
	"go-template/src/middleware"

	"github.com/gofiber/fiber/v2"

	fiberlog "github.com/gofiber/fiber/v2/log"
)



// GetAllUser Godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=[]model.User} "Successfully retrieved all users"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/user/getall [get]
func (h *HTTPGateway) GetAllUser(c *fiber.Ctx) error {
	users, err := h.UserService.GetAllUser()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: err.Error(),
			Data: nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: "Successfully get all user",
		Data: users,
	})
}

// GetUserByID Godoc
// @Summary Get users
// @Description Get users information
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=[]model.User} "Successfully retrieved users"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/user/get [get]
func (h *HTTPGateway) GetUserByID(c *fiber.Ctx) error {
	Token ,err := middleware.DecodeCookie(c)
	if err != nil {
		return err
	}
	id := Token.UserID
	user, err := h.UserService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: err.Error(),
			Data: nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: "Successfully retrieved user",
		Data: user,
	})
}


// UpdateUser Godoc
// @Summary Update users
// @Description Update users information
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User object"
// @Success 200 {object} model.Response{data=[]model.User} "Successfully updated users"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/user/update [put]
func (h *HTTPGateway) UpdateUser(c *fiber.Ctx) error {
	Token, err := middleware.DecodeCookie(c)
	if err != nil {
		return err
	}
	id := Token.UserID
	fiberlog.Info("Updating user with ID: ", id)
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: 400,
			Message: "Bad Request",
			Data: nil,
		})
	}
	_, err = h.UserService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: err.Error(),
			Data: nil,
		})
	}
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		fiberlog.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: 400,
			Message: "Bad Request",
			Data: nil,
		})
	}
	user.ID = id
	updatedUser, err := h.UserService.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: err.Error(),
			Data: nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: "Successfully updated user",
		Data: updatedUser,
	})
}

// DeleteUser Godoc
// @Summary Delete users
// @Description Delete users information
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=[]model.User} "Successfully deleted users"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/user/delete [delete]
func (h *HTTPGateway) DeleteUser(c *fiber.Ctx) error {
	Token, err := middleware.DecodeCookie(c)
	if err != nil {
		return err
	}
	id := Token.UserID
	fiberlog.Info("Updating user with ID: ", id)
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: 400,
			Message: "Bad Request",
			Data: nil,
		})
	}
	_, err = h.UserService.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: err.Error(),
			Data: nil,
		})
	}
	err = h.UserService.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: err.Error(),
			Data: nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: "Successfully deleted user",
		Data: nil,
	})
}
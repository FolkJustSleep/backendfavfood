package gateway

import (
	"go-template/data/model"
	"go-template/src/middleware"

	"github.com/gofiber/fiber/v2"

)

func (h *HTTPGateway) CheckIn(c *fiber.Ctx) error {
	Token ,err := middleware.DecodeCookie(c)
	if err != nil {
		return err
	}
	var timestamp model.TimestampModel
	timestamp.UserID = Token.UserID
	resp, err := h.TimestampService.CheckIn(timestamp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: err.Error(),
			Data: nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: "Check-In Successful",
		Data: resp,
	})
}

func (h *HTTPGateway) CheckOut(c *fiber.Ctx) error {
	Token ,err := middleware.DecodeCookie(c)
	if err != nil {
		return err
	}
	resp, err := h.TimestampService.CheckOut(Token.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: err.Error(),
			Data: nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: "Check-Out Successful",
		Data: resp,
	})
}

func (h *HTTPGateway) GetCheckedInEmployee(c *fiber.Ctx) error {
	resp, err := h.TimestampService.GetCheckedInEmployee()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: err.Error(),
			Data: nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: "Successfully retrieved checked-in employees",
		Data: resp,
	})
}

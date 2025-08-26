package gateway

import (
	"github.com/gofiber/fiber/v2"
	// fiberlog "github.com/gofiber/fiber/v2/log"

	"go-template/data/model"
	"go-template/src/middleware"
)

func (h *HTTPGateway) CreateCashControl(ctx *fiber.Ctx) error {
	Token, err := middleware.DecodeCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: " Unauthorized" + err.Error(),
		})
	}
	userID := Token.UserID
	var cashControlModel model.CashControl
	if err := ctx.BodyParser(&cashControlModel); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: 400,
			Message: " Bad Request" + err.Error(),
		})
	}
	cashControlModel.UserID = userID
	cashControl, err := h.CashControlService.CreateCashControl(cashControlModel)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: " Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(model.Response{
		Status: 201,
		Message: " Created Successfully",
		Data: cashControl,
	})
}

func (h *HTTPGateway) GetCashControlByID(ctx *fiber.Ctx) error {
	id := ctx.Query("id")
	cashControl, err := h.CashControlService.GetCashControlByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: " Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: "Successfully retrieved cash control",
		Data: cashControl,
	})
}

func (h *HTTPGateway) GetCashControlByUserID(ctx *fiber.Ctx) error {
	Token, err := middleware.DecodeCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: fiber.StatusUnauthorized,
			Message: "Unauthorized" + err.Error(),
		})
	}
	userID := Token.UserID
	if userID == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: fiber.StatusUnauthorized,
			Message: "Unauthorized",
		})
	}
	cashControl, err := h.CashControlService.GetCashControlByUserID(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status: fiber.StatusOK,
		Message: "Successfully retrieved cash control",
		Data: cashControl,
	})
}

func (h *HTTPGateway) GetAllCashControls(ctx *fiber.Ctx) error {
	cashControls, err := h.CashControlService.GetAllCashControls()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status: fiber.StatusOK,
		Message: "Successfully retrieved all cash controls",
		Data: cashControls,
	})
}

func (h *HTTPGateway) UpdateCashControl(ctx *fiber.Ctx) error {
	var cashControl model.CashControl
	if err := ctx.BodyParser(&cashControl); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: fiber.StatusBadRequest,
			Message: "Bad Request" + err.Error(),
		})
	}
	cashControl, err := h.CashControlService.UpdateCashControl(cashControl)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status: fiber.StatusOK,
		Message: "Successfully updated cash control",
		Data: cashControl,
	})
}

func (h *HTTPGateway) DeleteCashControl(ctx *fiber.Ctx) error {
	id := ctx.Query("id")
	if err := h.CashControlService.DeleteCashControl(id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: fiber.StatusInternalServerError,
			Message: "Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status: fiber.StatusOK,
		Message: "Successfully deleted cash control",
		Data: nil,
	})
}

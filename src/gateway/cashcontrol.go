package gateway

import (
	"github.com/gofiber/fiber/v2"
	// fiberlog "github.com/gofiber/fiber/v2/log"

	"go-template/data/model"
	"go-template/src/middleware"
)

// Create Cashcontrol Godoc
// @Summary Create a new cash control
// @Description Create a new cash control
// @Tags cash
// @Accept json
// @Produce json
// @Param cash body model.CashControl true "CashControl object"
// @Success 201 {object} model.Response{data=model.CashControl} "Successfully created cash control"
// @Failure 400 {object} model.Response "Bad Request"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/cashcontrol/create [post]
func (h *HTTPGateway) CreateCashControl(ctx *fiber.Ctx) error {
	_, err := middleware.DecodeCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: " Unauthorized" + err.Error(),
		})
	}
	// userID := Token.UserID

	var cashControlModel model.CashControl
	if err := ctx.BodyParser(&cashControlModel); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: 400,
			Message: " Bad Request" + err.Error(),
		})
	}
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

// GetCashControlByID goDoc
// @Summary Get cash control by id
// @Description Get cash control by id
// @Tags cash
// @Accept json
// @Produce json
// @Param id query string true "CashControl ID"
// @Success 200 {object} model.Response{data=model.CashControl} "Successfully retrieved cash control"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/cashcontrol/get [get]
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

// GetCashControlByID goDoc
// @Summary Get All cash control by user id
// @Description Get cash control by user id
// @Tags cash
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=model.CashControl} "Successfully retrieved cash control"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/cashcontrol/getbyuserid [get]
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

// GetCashControlByID goDoc
// @Summary Get all cash control
// @Description Get all cash control
// @Tags cash
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=model.CashControl} "Successfully retrieved cash control"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/cashcontrol/getall [get]
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

// UpdateCashControl Godoc
// @Summary Update CashControl
// @Description Update CashControl information
// @Tags cash
// @Accept json
// @Produce json
// @Param cashControl body model.CashControl true "CashControl object"
// @Success 200 {object} model.Response{data=model.CashControl} "Successfully updated CashControl"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/cashcontrol/update [put]
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

// DeleteCashControl Godoc
// @Summary Delete cash control
// @Description Delete cash control information
// @Tags cash
// @Accept json
// @Produce json
// @Param id query string true "CashControl ID"
// @Success 200 {object} model.Response{data=[]model.CashControl} "Successfully deleted cash control"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/cashcontrol/delete [delete]
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

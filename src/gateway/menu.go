package gateway

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	// fiberlog "github.com/gofiber/fiber/v2/log"

	"go-template/data/model"
	"go-template/src/middleware"
)

// Create Menu Godoc
// @Summary Create a new menu
// @Description Create a new menu
// @Tags menu
// @Accept json
// @Produce json
// @Param menu body model.Menu true "Menu object"
// @Success 201 {object} model.Response{data=model.Menu} "Successfully created menu"
// @Failure 400 {object} model.Response "Bad Request"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/menu/create [post]
func (h *HTTPGateway) CreateMenu(ctx *fiber.Ctx) error {
	_, err := middleware.DecodeCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: " Unauthorized" + err.Error(),
		})
	}
	// userID := Token.UserID
	var menuModel model.Menu
	if err := ctx.BodyParser(&menuModel); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: 400,
			Message: " Bad Request" + err.Error(),
		})
	}
	menu, err := h.MenuService.CreateMenu(menuModel)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: " Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(model.Response{
		Status: 201,
		Message: " Created Successfully",
		Data: menu,
	})
}

// GetMenuByID goDoc
// @Summary Get menu by id
// @Description Get menu by id
// @Tags menu
// @Accept json
// @Produce json
// @Param id query string true "Menu ID"
// @Success 200 {object} model.Response{data=model.Menu} "Successfully retrieved menu"
// @Failure 400 {object} model.Response "Bad Request"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/menu/get [get]
func (h *HTTPGateway) GetMenuByID(ctx *fiber.Ctx) error {
	_, err := middleware.DecodeCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: " Unauthorized" + err.Error(),
		})
	}
	// userID := Token.UserID

	id := ctx.Query("id")
	menu, err := h.MenuService.GetMenuByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: " Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: " Retrieved Successfully",
		Data: menu,
	})
}

// GetAllMenus goDoc
// @Summary Get all menus
// @Description Get all menus
// @Tags menu
// @Accept json
// @Produce json
// @Success 200 {object} model.Response{data=[]model.Menu} "Successfully retrieved all menus"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/menu/getall [get]
func (h *HTTPGateway) GetAllMenus(ctx *fiber.Ctx) error {
	_, err := middleware.DecodeCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: " Unauthorized" + err.Error(),
		})
	}
	// userID := Token.UserID

	menus, err := h.MenuService.GetAllMenus()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: " Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: " Retrieved Successfully",
		Data: menus,
	})
}

// UpdateMenu goDoc
// @Summary Update a menu
// @Description Update a menu
// @Tags menu
// @Accept json
// @Produce json
// @Param menu body model.Menu true "Menu object"
// @Success 200 {object} model.Response{data=model.Menu} "Successfully updated menu"
// @Failure 400 {object} model.Response "Bad Request"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/menu/update [put]
func (h *HTTPGateway) UpdateMenu(ctx *fiber.Ctx) error {
	_, err := middleware.DecodeCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: " Unauthorized" + err.Error(),
		})
	}
	// userID := Token.UserID

	var menuModel model.Menu
	if err := ctx.BodyParser(&menuModel); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.Response{
			Status: 400,
			Message: " Bad Request" + err.Error(),
		})
	}
	updatedMenu, err := h.MenuService.UpdateMenu(menuModel)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: " Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: " Updated Successfully",
		Data: updatedMenu,
	})
}

// DeleteMenu goDoc
// @Summary Delete a menu
// @Description Delete a menu
// @Tags menu
// @Accept json
// @Produce json
// @Param id query string true "Menu ID"
// @Success 200 {object} model.Response "Successfully deleted menu"
// @Failure 400 {object} model.Response "Bad Request"
// @Failure 500 {object} model.Response "Internal Server Error"
// @Router /api/menu/delete [delete]
func (h *HTTPGateway) DeleteMenu(ctx *fiber.Ctx) error {
	_, err := middleware.DecodeCookie(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Status: 401,
			Message: " Unauthorized" + err.Error(),
		})
	}
	// userID := Token.UserID

	id := ctx.Query("id")
	fmt.Println("Deleting menu with ID:", id)
	fmt.Println("Type of id:", fmt.Sprintf("%T", id))
	err = h.MenuService.DeleteMenu(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Status: 500,
			Message: " Internal Server Error" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(model.Response{
		Status: 200,
		Message: " Deleted Successfully",
	})
}
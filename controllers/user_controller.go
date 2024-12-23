package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/snykk/go-fx/services"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	user, err := uc.userService.GetUser(idInt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve user",
		})
	}
	return c.JSON(user)
}

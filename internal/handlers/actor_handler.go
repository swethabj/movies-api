package handlers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/swethabj/movies-api/internal/models"
	"github.com/swethabj/movies-api/internal/repository"
	"github.com/swethabj/movies-api/utils"
	"gorm.io/gorm"
)

type CreateActorRequest struct {
	Name string
}

func CreateActors(c *fiber.Ctx) error {
	var req CreateActorRequest
	if err := c.BodyParser(&req); err != nil {
		utils.Logger.Warnf("invalid body %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	newactor := &models.Actor{
		Name: req.Name,
	}

	if err := repository.CreateActor(newactor); err != nil {
		utils.Logger.Errorf("create actor error %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unable to create actor"})
	}

	return c.Status(fiber.StatusCreated).JSON(newactor)
}

func GetActors(c *fiber.Ctx) error {
	actors, err := repository.GetAllActors()
	if err != nil {
		utils.Logger.Errorf("get actors error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch actors"})
	}
	return c.JSON(actors)
}

// GetActorByID
func GetActor(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.Logger.Warnf("invalid id %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	actor, err := repository.GetActorByID(uint(id64))
	if err != nil {
		utils.Logger.Errorf("server error %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	if actor == nil {
		utils.Logger.Warnf("actor not found %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "actor not found"})
	}
	return c.JSON(actor)
}

func DeleteActor(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.Logger.Warnf("invalid id %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := repository.DeleteActorByID(uint(id64)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.Logger.Warnf("actor not found %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "actor not found"})
		} else {
			utils.Logger.Errorf("server error %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "server error"})
		}
	}
	return c.SendStatus(fiber.StatusNoContent)

}

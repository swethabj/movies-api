package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/swethabj/movies-api/internal/models"
	"github.com/swethabj/movies-api/internal/repository"
	"github.com/swethabj/movies-api/utils"
)

type CreateGenreRequest struct {
	Name string
}

func CreateGenre(c *fiber.Ctx) error {
	var req CreateGenreRequest
	if err := c.BodyParser(&req); err != nil {
		utils.Logger.Warnf("invalid body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	newgenre := &models.Genre{
		Name: req.Name,
	}

	if err := repository.CreateGenre(newgenre); err != nil {
		utils.Logger.Errorf("create genre error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unable to create genre"})
	}

	return c.Status(fiber.StatusCreated).JSON(newgenre)
}

func GetGenres(c *fiber.Ctx) error {
	genres, err := repository.GetAllGenres()
	if err != nil {
		utils.Logger.Errorf("get genres error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch genres"})
	}
	return c.JSON(genres)
}

func GetGenre(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.Logger.Warnf("invalid body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	genre, err := repository.GetGenreByID(uint(id64))
	if err != nil {
		utils.Logger.Errorf("invalid body: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	if genre == nil {
		utils.Logger.Warnf("genre not found: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "genre not found"})
	}
	return c.JSON(genre)
}

func DeleteGenre(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.Logger.Warnf("invalid id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := repository.DeleteGenreByID(uint(id64)); err != nil {
		utils.Logger.Errorf("server error : %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/swethabj/movies-api/internal/models"
	"github.com/swethabj/movies-api/internal/repository"
	"github.com/swethabj/movies-api/utils"
)

// Request payloads
type CreateMovieRequest struct {
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_year"`
	GenreID     *uint  `json:"genre_id"`
	ActorIDs    []uint `json:"actor_ids"`
}

type UpdateMovieRequest = CreateMovieRequest

func CreateMovie(c *fiber.Ctx) error {
	var req CreateMovieRequest
	if err := c.BodyParser(&req); err != nil {
		utils.Logger.Warnf("invalid body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	movie := &models.Movie{
		Title:       req.Title,
		ReleaseYear: req.ReleaseYear,
		GenreID:     req.GenreID,
	}

	if err := repository.CreateMovie(movie, req.ActorIDs); err != nil {
		utils.Logger.Errorf("create movie error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unable to create movie"})
	}

	return c.Status(fiber.StatusCreated).JSON(movie)
}

// GetMovies godoc
// @Summary Get all movies
// @Description Get details of all movies
// @Tags movies
// @Accept  json
// @Produce  json
// @Success 200 {array}  models.Movie
// @Failure 400 {object} map[string]string
// @Router /movies [get]
func GetMovies(c *fiber.Ctx) error {
	movies, err := repository.GetAllMovies()
	if err != nil {
		utils.Logger.Errorf("get movies error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch movies"})
	}
	return c.JSON(movies)
}

func GetMovie(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.Logger.Warnf("invalid id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	movie, err := repository.GetMovieByID(uint(id64))
	if err != nil {
		utils.Logger.Errorf("server error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	if movie == nil {
		utils.Logger.Warnf("movie not found: %v", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "movie not found"})
	}
	return c.JSON(movie)
}

func UpdateMovie(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.Logger.Warnf("invalid id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	var req UpdateMovieRequest
	if err := c.BodyParser(&req); err != nil {
		utils.Logger.Warnf("invalid body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := repository.UpdateMovie(uint(id64), &models.Movie{
		Title:       req.Title,
		ReleaseYear: req.ReleaseYear,
		GenreID:     req.GenreID,
	}, req.ActorIDs); err != nil {
		utils.Logger.Errorf("update failed: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "update failed"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func DeleteMovie(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.Logger.Warnf("invalid id: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := repository.DeleteMovie(uint(id64)); err != nil {
		utils.Logger.Errorf("delete failed: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "delete failed"})
	}
	return c.SendStatus(fiber.StatusNoContent)

}

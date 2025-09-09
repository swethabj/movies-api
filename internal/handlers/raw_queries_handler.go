package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/swethabj/movies-api/internal/repository"
	"github.com/swethabj/movies-api/utils"
)

func GetRawMovies(c *fiber.Ctx) error {
	rows, err := RawMoviesJoinHandler() // we'll add this helper below or you can call repo.RawMovieJoin directly
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// Raw join helper for debug route
func RawMoviesJoinHandler() ([]repository.MovieJoinRow, error) {
	return repository.RawMovieJoin()
}

func GetMoviesGenre(c *fiber.Ctx) error {
	genrename := c.Params("name")

	rows, err := repository.GetMoviesByGenreName(genrename)
	if err != nil {
		utils.Logger.Errorf("failed to fetch movies by genre %s: %v", genrename, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}

	if len(rows) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "no movies found for this genre"})
	}

	return c.JSON(rows)
}

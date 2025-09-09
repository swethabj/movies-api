package server

import (
	"github.com/gofiber/fiber/v2"
	v2logger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/swethabj/movies-api/internal/handlers"
)

func NewRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "movies-api",
	})

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Fiber logger middleware writes to stdout (in production pipe to file or aggregator)
	app.Use(v2logger.New())

	v1 := app.Group("/api/v1")
	movies := v1.Group("/movies")
	genres := v1.Group("/genres")
	actors := v1.Group("/actors")
	//movie
	movies.Post("/", handlers.CreateMovie)
	movies.Get("/", handlers.GetMovies)
	movies.Get("/:id", handlers.GetMovie)
	movies.Put("/:id", handlers.UpdateMovie)
	movies.Delete("/:id", handlers.DeleteMovie)
	//genre
	genres.Post("/", handlers.CreateGenre)
	genres.Get("/", handlers.GetGenres)
	genres.Get("/:id", handlers.GetGenre)
	genres.Delete("/:id", handlers.DeleteGenre)
	//actor
	actors.Post("/", handlers.CreateActors)
	actors.Get("/", handlers.GetActors)
	actors.Get("/:id", handlers.GetActor)
	actors.Delete("/:id", handlers.DeleteActor)

	// optional diagnostics route to show raw sql join example
	v1.Get("/raw-queries", handlers.GetRawMovies)
	v1.Get("/raw-queries/:name", handlers.GetMoviesGenre)

	return app

}

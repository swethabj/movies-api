package repository

import (
	"errors"

	"github.com/swethabj/movies-api/config"
	"github.com/swethabj/movies-api/internal/models"
	"github.com/swethabj/movies-api/utils"
	"gorm.io/gorm"
)

// Create a movie and associate actors
func CreateMovie(movie *models.Movie, actorsIDs []uint) error {
	tx := config.DB.Begin()
	// if tx != nil {
	// 	return tx.Error
	// }

	if err := tx.Create(movie).Error; err != nil {
		tx.Rollback()
		utils.Logger.Errorf("create movie failed %v", err)
		return err
	}

	if len(actorsIDs) > 0 {
		var actors []models.Actor
		if err := tx.Find(&actors, actorsIDs).Error; err != nil {
			tx.Rollback()
			return err
		}
		if len(actors) != len(actorsIDs) {
			tx.Rollback()
			return errors.New("some actors not found")
		}
		if err := tx.Model(movie).Association("Actors").Replace(&actors); err != nil {
			tx.Rollback()
			return err
		}

	}
	return tx.Commit().Error
}

// Get all movies with genre and actors (GORM preload)
func GetAllMovies() ([]models.Movie, error) {
	var movies []models.Movie
	if err := config.DB.Preload("Genre").Preload("Actors").Find(&movies).Error; err != nil {
		utils.Logger.Errorf("failed to fetch movies %v", err)
		return nil, err
	}
	return movies, nil
}

// Get detail movie by id
func GetMovieByID(id uint) (*models.Movie, error) {
	var movie models.Movie
	if err := config.DB.Preload("Genre").Preload("Actors").First(&movie, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &movie, nil
}

// Update movie fields and actor associations
func UpdateMovie(id uint, payload *models.Movie, actorIDs []uint) error {
	movie, err := GetMovieByID(id)
	if err != nil {
		return err
	}
	if movie == nil {
		return gorm.ErrRecordNotFound
	}
	tx := config.DB.Begin()
	if err := tx.Model(movie).Updates(models.Movie{
		Title:       payload.Title,
		ReleaseYear: payload.ReleaseYear,
		GenreID:     payload.GenreID,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if actorIDs != nil {
		var actors []models.Actor
		if err := tx.Find(&actors, actorIDs).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Model(movie).Association("Actors").Replace(&actors); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error

}

// Delete Move by Id
func DeleteMovie(id uint) error {
	if err := config.DB.Delete(&models.Movie{}, id).Error; err != nil {
		return err
	}
	return nil
}

// Example raw SQL multi-select (returns flattened rows) - for advanced use
type MovieJoinRow struct {
	MovieID     uint
	Title       string
	ReleaseYear int
	GenreName   *string
	ActorName   *string
}

package repository

import (
	"errors"

	"github.com/swethabj/movies-api/config"
	"github.com/swethabj/movies-api/internal/models"
	"github.com/swethabj/movies-api/utils"
	"gorm.io/gorm"
)

func CreateGenre(newgenre *models.Genre) error {
	tx := config.DB.Begin() //tx := config.DB.Begin() never returns nil. It returns a *gorm.DB.
	// if tx != nil {
	// 	return tx.Error
	// }
	if err := tx.Create(newgenre).Error; err != nil {
		tx.Rollback()
		utils.Logger.Errorf("create genre failed %v", err)
		return err
	}
	return tx.Commit().Error

	// OR
	// if err := config.DB.Create(newgenre).Error; err != nil {
	// 	utils.Logger.Errorf("create genre failed %v", err)
	// 	return err
	// }
	// return nil
}

func GetAllGenres() ([]models.Genre, error) {
	var genres []models.Genre
	if err := config.DB.Find(&genres).Error; err != nil {
		utils.Logger.Errorf("failed to fetch genres %v", err)
		return nil, err
	}
	return genres, nil
}

func GetGenreByID(id uint) (*models.Genre, error) {
	var genre models.Genre
	if err := config.DB.First(&genre, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &genre, nil
}

func DeleteGenreByID(id uint) error {
	var genre models.Genre
	if err := config.DB.Delete(&genre, id).Error; err != nil {
		return err
	}
	return nil
}

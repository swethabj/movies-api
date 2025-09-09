package repository

import (
	"errors"

	"github.com/swethabj/movies-api/config"
	"github.com/swethabj/movies-api/internal/models"
	"github.com/swethabj/movies-api/utils"
	"gorm.io/gorm"
)

func CreateActor(newactor *models.Actor) error {
	if err := config.DB.Create(newactor).Error; err != nil {
		utils.Logger.Errorf("create actor failed %v", err)
		return err
	}
	return nil
}

func GetAllActors() ([]models.Actor, error) {
	var actors []models.Actor
	if err := config.DB.Find(&actors).Error; err != nil {
		utils.Logger.Errorf("failed to fetch actors %v", err)
		return nil, err
	}
	return actors, nil
}

func GetActorByID(id uint) (*models.Actor, error) {
	var actor models.Actor
	if err := config.DB.First(&actor, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &actor, nil
}

func DeleteActorByID(id uint) error {
	result := config.DB.Delete(&models.Actor{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

package models

type MovieActor struct {
	MovieID uint `gorm:"primaryKey"`
	ActorID uint `gorm:"primaryKey"`
}

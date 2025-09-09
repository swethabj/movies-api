package models

import "time"

type Movie struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	ReleaseYear int       `json:"release_year"`
	GenreID     *uint     `json:"genre_id"` // nullable
	Genre       *Genre    `json:"genre,omitempty" gorm:"foreignKey:GenreID;references:ID"`
	Actors      []Actor   `json:"actors,omitempty" gorm:"many2many:movie_actors;constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

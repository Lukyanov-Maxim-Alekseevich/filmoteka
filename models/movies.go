package models

type Movie struct {
	Id          int     `json:"-" db:"id"`
	Name        string  `json:"name" binding:"required" db:"name"`
	Description string  `json:"description" binding:"required" db:"description"`
	ReleaseDate string  `json:"release_date" binding:"required" db:"release_date"`
	Rating      float64 `json:"rating" binding:"required" db:"rating"`
}

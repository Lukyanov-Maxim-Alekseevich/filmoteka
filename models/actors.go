package models

type Actor struct {
	Id         int    `json:"-" db:"id"`
	Name       string `json:"name" binding:"required"`
	Gender     string `json:"gender" binding:"required"`
	Date_birth string `json:"date_birth" binding:"required"`
	Movie_id   int    `json:"movie_id" binding:"required"`
}

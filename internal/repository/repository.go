package repository

import (
	"github.com/MrSaveliy/vk-filmoteca/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	CreateRole(role models.Role) (int, error)
	GetUser(email, password string) (models.User, error)
}

type Movies interface {
	CreateMovie(movie models.Movie) (int, error)
	GetMovieById(id int) (models.Movie, error)
	UpdateMovieById(id int, new_movie models.Movie) (int, error)
	DeleteMovieById(id int) (int, error)
}

type Actors interface {
	CreateActor(actor models.Actor) (int, error)
	GetActorById(id int) (models.Actor, error)
	UpdateActorById(id int, new_actor models.Actor) (int, error)
	DeleteActorById(id int) (int, error)
}

type Repository struct {
	Authorization
	Movies
	Actors
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Movies:        NewMoviesPostgres(db),
		Actors:        NewActorsPostgres(db),
	}
}

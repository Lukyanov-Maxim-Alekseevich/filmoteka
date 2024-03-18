package service

import (
	"github.com/MrSaveliy/vk-filmoteca/internal/repository"
	"github.com/MrSaveliy/vk-filmoteca/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	CreateRole(role models.Role) (int, error)
	ParseToken(token string) (int, error)
	GenerateToken(email, password string) (string, error)
}

type Movies interface {
	CreateMovie(movie models.Movie) (int, error)
	GetMovieById(id int) (models.Movie, error)
	UpdateMovieById(id int, new_movie models.Movie) (int, error)
	DeleteMoviesById(id int) (int, error)
}

type Actors interface {
	CreateActor(actor models.Actor) (int, error)
	GetActorById(id int) (models.Actor, error)
	UpdateActorById(id int, new_actor models.Actor) (int, error)
	DeleteActorById(id int) (int, error)
}

type Service struct {
	Authorization
	Movies
	Actors
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Movies:        NewMoviesService(repos.Movies),
		Actors:        NewActorsService(repos.Actors),
	}
}

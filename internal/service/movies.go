package service

import (
	"github.com/MrSaveliy/vk-filmoteca/internal/repository"
	"github.com/MrSaveliy/vk-filmoteca/models"
)

type MoviesService struct {
	repo repository.Movies
}

func NewMoviesService(repo repository.Movies) *MoviesService {
	return &MoviesService{repo: repo}
}

func (s *MoviesService) CreateMovie(movie models.Movie) (int, error) {
	return s.repo.CreateMovie(movie)
}

func (s *MoviesService) GetMovieById(id int) (models.Movie, error) {
	return s.repo.GetMovieById(id)
}

func (s *MoviesService) UpdateMovieById(id int, new_movie models.Movie) (int, error) {
	return s.repo.UpdateMovieById(id, new_movie)
}

func (s *MoviesService) DeleteMoviesById(id int) (int, error) {
	return s.repo.DeleteMovieById(id)
}

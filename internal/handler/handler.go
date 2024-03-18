package handler

import (
	"net/http"

	"github.com/MrSaveliy/vk-filmoteca/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(serveces *service.Service) *Handler {
	return &Handler{services: serveces}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/sign-up", h.signUpHandler)
	mux.HandleFunc("/auth/sign-in", h.signInHandler)

	mux.HandleFunc("/role/create", h.createRole)

	mux.HandleFunc("/movie/create", h.createMovie)
	mux.HandleFunc("/movie/", h.getMovieById)
	mux.HandleFunc("/movie/update/", h.updateMovieById)
	mux.HandleFunc("/movie/delete/", h.deleteMovieById)

	mux.HandleFunc("/actor/create", h.createActor)
	mux.HandleFunc("/actor/", h.getActorById)
	mux.HandleFunc("/actor/update/", h.updateActorById)
	mux.HandleFunc("/actor/delete/", h.deleteActorById)

	mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger/index.html")
	})

	return mux
}

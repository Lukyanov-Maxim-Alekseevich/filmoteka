package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/MrSaveliy/vk-filmoteca/models"
)

func (h *Handler) createActor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input models.Actor
	if err := json.Unmarshal(body, &input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Actors.CreateActor(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{
		"id": id,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func (h *Handler) getActorById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		newErrorResponse(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	actorIdStr := pathParts[len(pathParts)-1]
	actorId, err := strconv.Atoi(actorIdStr)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Invalid actor ID"+err.Error())
		return
	}

	actor, err := h.services.Actors.GetActorById(actorId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Internal Server Error"+err.Error())
		return
	}

	response := map[string]interface{}{
		"name":       actor.Name,
		"gender":     actor.Gender,
		"date_bitrh": actor.Date_birth,
		"movie_id":   actor.Movie_id,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Internal Server Error"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func (h *Handler) updateActorById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Bad Request"+err.Error())
		return
	}

	var input models.Actor
	if err := json.Unmarshal(body, &input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "Bad Request"+err.Error())
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		newErrorResponse(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	actorIdStr := pathParts[len(pathParts)-1]
	actorId, err := strconv.Atoi(actorIdStr)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Invalid actor ID"+err.Error())
		return
	}

	id, err := h.services.Actors.UpdateActorById(actorId, input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Internal Server Error"+err.Error())
		return
	}

	response := map[string]interface{}{
		"id": id,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Internal Server Error"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func (h *Handler) deleteActorById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		newErrorResponse(w, http.StatusBadRequest, "Invalid URL")
		return
	}

	actorIdStr := pathParts[len(pathParts)-1]
	actorId, err := strconv.Atoi(actorIdStr)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Invalid actor ID"+err.Error())
		return
	}

	id, err := h.services.Actors.DeleteActorById(actorId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Error retrieving actor: "+err.Error())
		return
	}

	response := map[string]interface{}{
		"id": id,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, "Internal Server Error"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

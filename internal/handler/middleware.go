package handler

import (
	"context"
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		newErrorResponse(w, http.StatusUnauthorized, "empty header")
		return
	}
	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}
	ctx := context.WithValue(r.Context(), "userCtx", userId)
	// Обновляем контекст запроса с новым значением
	r = r.WithContext(ctx)
	// Вызываем следующий обработчик с обновленным контекстом
	http.DefaultServeMux.ServeHTTP(w, r)
}

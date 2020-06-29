package auth

import (
	"encoding/json"
	"net/http"

	"github.com/PhongVX/taskmanagement/internal/pkg/http/response"
	"github.com/PhongVX/taskmanagement/internal/pkg/log"
)

func NewHTTPHandler(srv ServiceInterface) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var auth Auth
	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		log.WithContext(r.Context()).Infof("Failed to decode body, err: %v", err)
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	tokens, err := h.srv.Login(r.Context(), &auth)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, tokens)
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var t Tokens
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.WithContext(r.Context()).Infof("Failed to decode body, err: %v", err)
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	tokens, err := h.srv.Refresh(r.Context(), &t)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, tokens)
}

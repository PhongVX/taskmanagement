package task

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/PhongVX/taskmanagement/internal/pkg/http/response"
	"github.com/PhongVX/taskmanagement/internal/pkg/log"
	"github.com/PhongVX/taskmanagement/internal/pkg/types/responsetype"
	"github.com/PhongVX/taskmanagement/internal/pkg/utils/handlerutil"

	"github.com/gorilla/mux"
)

func NewHTTPHandler(srv Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	req := FindingRequestObject{
		Offset:      handlerutil.IntFromQuery("offset", queries, 0),
		Limit:       handlerutil.IntFromQuery("limit", queries, 15),
		SprintID:    queries["sprint_id"],
		CreatedByID: queries.Get("created_by_id"),
		SortBy:      queries["sort_by"],
	}
	tasks, err := h.srv.repo.FindAll(r.Context(), req)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, responsetype.Base{
		Result: tasks,
	})
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		response.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	t, err := h.srv.repo.FindByID(r.Context(), id)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, t)
}

func (h *Handler) Insert(w http.ResponseWriter, r *http.Request) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.WithContext(r.Context()).Infof("Failed to decode body, err: %v", err)
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.repo.Insert(r.Context(), &t); err != nil {
		log.WithContext(r.Context()).Errorf("Could not create article, err: %v", err)
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, t)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		response.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	err := h.srv.repo.Delete(r.Context(), id)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{
		"id": id,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var t Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		log.WithContext(r.Context()).Infof("Failed to decode body, err: %v", err)
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.repo.Update(r.Context(), &t); err != nil {
		log.WithContext(r.Context()).Errorf("Could not update article, err: %v", err)
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, map[string]string{
		"id": t.ID.Hex(),
	})
}

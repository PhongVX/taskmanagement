package task

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PhongVX/taskmanagement/internal/pkg/http/response"
	"github.com/PhongVX/taskmanagement/internal/pkg/log"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(srv Service) *Handler {
	return &Handler{
		srv: srv,
	}
}

//===================================Move To Common Package=================================================
func IntFromQuery(k string, req url.Values, def int) int {
	v := req.Get(k)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return i
}

//====================================================================================================

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	req := FindingRequestObject{
		Offset: IntFromQuery("offset", queries, 0),
		Limit:  IntFromQuery("limit", queries, 15),
		SortBy: queries["sort_by"],
	}
	tasks, err := h.srv.repo.FindAll(r.Context(), req)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, tasks)
}

func (h *Handler) FindByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.WithContext(r.Context()).Infof("FindByID/ id %s", id)
	if id == "" {
		response.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	t, err := h.srv.repo.FindByID(r.Context(), id)
	log.WithContext(r.Context()).Infof("FindByID/ t %v", t)
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

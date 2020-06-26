package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/PhongVX/taskmanagement/internal/pkg/http/response"
	"github.com/PhongVX/taskmanagement/internal/pkg/log"
	"github.com/PhongVX/taskmanagement/internal/pkg/types/responsetype"

	"github.com/gorilla/mux"
)

func NewHTTPHandler(srv ServiceInterface) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	//queries := r.URL.Query()
	req := FindingRequestObject{
		// Offset: handlerutil.IntFromQuery("offset", queries, 0),
		// Limit:  handlerutil.IntFromQuery("limit", queries, 15),
		// SortBy: queries["sort_by"],
	}
	users, err := h.srv.FindAll(r.Context(), req)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, responsetype.Base{
		Result: users,
	})
}

func (h *Handler) Insert(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.WithContext(r.Context()).Infof("Failed to decode body, err: %v", err)
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.Insert(r.Context(), &u); err != nil {
		log.WithContext(r.Context()).Errorf("Could not create user, err: %v", err)
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, responsetype.Base{
		Result: u,
	})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		response.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	err := h.srv.Delete(r.Context(), id)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, responsetype.Base{
		ID: id,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.WithContext(r.Context()).Infof("Failed to decode body, err: %v", err)
		response.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.Update(r.Context(), &u); err != nil {
		log.WithContext(r.Context()).Errorf("Could not update user, err: %v", err)
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, responsetype.Base{
		ID: u.ID.Hex(),
	})
}

func (h *Handler) FindByIdentity(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		response.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	u, err := h.srv.FindByIdentity(r.Context(), id)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.JSON(w, http.StatusOK, responsetype.Base{
		Result: u,
	})
}

package oauth

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

func (h *Handler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	h.srv.GoogleLogin(w, r)
}

func (h *Handler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.WithContext(r.Context()).Error("Invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.WithContext(r.Context()).Error(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	log.WithContext(r.Context()).Infof("Data: %v", data)
	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	var u UserFromRequest
	if err := json.Unmarshal(data, &u); err != nil {
		log.WithContext(r.Context()).Error(err)
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	tokens, err := h.srv.GoogleCallback(r.Context(), &u)
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	log.WithContext(r.Context()).Infof("Token %v", tokens)
	// response.JSON(w, http.StatusOK, tokens)
	urlRedirect := TOKEN_LOGIN_URL + "/" + tokens.AccessToken + "/" + tokens.RefreshToken
	http.Redirect(w, r, urlRedirect, http.StatusFound)
}

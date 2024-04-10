package user

import (
	"github.com/gorilla/mux"
	types "github.com/lelandaure/ecom/types"
	"github.com/lelandaure/ecom/utils"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin()).Methods("POST")
	router.HandleFunc("/register", h.handleRegister()).Methods("POST")
}

func (h *Handler) handleLogin() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (h Handler) handleRegister() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestBody types.RegisterUserRequest
		if err := utils.ParseJson(r, &requestBody); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
		}

	}
}

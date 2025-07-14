package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nishant007-tech/GoDig/internal/service"
	"go.uber.org/dig"
)

// UserHandler handles HTTP routes.
type UserHandler struct {
	svc service.UserService
}

// Params for NewUserHandler.
type handlerParams struct {
	dig.In

	Svc service.UserService
}

// Constructor injected by Dig.
func NewUserHandler(p handlerParams) *UserHandler {
	return &UserHandler{svc: p.Svc}
}

// Handle sends JSON response.
func (h *UserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	users, _ := h.svc.List(r.Context())
	json.NewEncoder(w).Encode(users)
}

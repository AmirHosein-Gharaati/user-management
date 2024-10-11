package handler

import (
	"net/http"

	"github.com/AmirHosein-Gharaati/user-management/internal/adapter/handler/mapper"
	"github.com/AmirHosein-Gharaati/user-management/internal/core/domain"
	"github.com/AmirHosein-Gharaati/user-management/internal/core/port"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	err := mapper.ReadJSON(w, r, &req)
	if err != nil {
		mapper.ErrorJSON(w, err)
		return
	}

	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	registeredUser, err := h.userService.Register(&user)
	if err != nil {
		mapper.ErrorJSON(w, err)
		return
	}

	res := &registerResponse{
		ID:    registeredUser.ID,
		Name:  registeredUser.Name,
		Email: registeredUser.Email,
	}

	mapper.WriteJSON(w, http.StatusCreated, res)
}

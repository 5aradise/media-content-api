package handlers

import (
	"context"
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/5aradise/media-content-api/src/internal/types"
	"github.com/5aradise/media-content-api/src/pkg/api"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage interface {
	CreateUser(ctx context.Context, firstName, lastName, email string, hashedPassword [60]byte) (types.User, error)
	ListUsers(ctx context.Context) ([]types.User, error)
	GetUserById(ctx context.Context, id int32) (types.User, error)
	UpdateUserById(ctx context.Context, id int32, firstName, lastName, email string, hashedPassword [60]byte) (types.User, error)
	DeleteUserById(ctx context.Context, id int32) error
}

type UserService struct {
	db UserStorage
}

func NewUserService(db UserStorage) *UserService {
	return &UserService{
		db: db,
	}
}

type CreateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// @Summary      Create an user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input body CreateUserRequest true "User info"
// @Success      201  {object}  types.User
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users [post]
func (s UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := api.DecodeJSON(r, &req); err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			api.WriteError(w, http.StatusBadRequest, "password length exceeds 72 bytes")
			return
		}

		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := s.db.CreateUser(r.Context(), req.FirstName, req.LastName, req.Email, [60]byte(hashedPassword))
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusCreated, user)
}

type ListUsersResponse struct {
	Users []types.User `json:"users"`
}

// @Summary      List users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object} ListUsersResponse
// @Failure      500  {object} api.ErrorResponse
// @Router       /users [get]
func (s UserService) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.db.ListUsers(r.Context())
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, ListUsersResponse{users})
}

// @Summary      Get user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      201  {object}  types.User
// @Failure      400  {object}  api.ErrorResponse
// @Router       /users/{id} [get]
func (s UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	idS := r.PathValue("id")
	if idS == "" {
		panic("empty id path value")
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if id > math.MaxInt32 {
		api.WriteError(w, http.StatusBadRequest, "id cannot be greater than 2147483647")
		return
	}

	user, err := s.db.GetUserById(r.Context(), int32(id))
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, user)
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// @Summary      Update user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path int               true "User ID"
// @Param        input body CreateUserRequest true "Updated info"
// @Success      201  {object}  types.User
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/{id} [patch]
func (s UserService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idS := r.PathValue("id")
	if idS == "" {
		panic("empty id path value")
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if id > math.MaxInt32 {
		api.WriteError(w, http.StatusBadRequest, "id cannot be greater than 2147483647")
		return
	}

	var req UpdateUserRequest
	if err := api.DecodeJSON(r, &req); err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := s.db.GetUserById(r.Context(), int32(id))
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if req.FirstName != "" {
		newUser.FirstName = req.FirstName
	}

	if req.LastName != "" {
		newUser.LastName = req.LastName
	}

	if req.Email != "" {
		newUser.Email = req.Email
	}

	newPassword := [60]byte([]byte(newUser.Password))
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			if errors.Is(err, bcrypt.ErrPasswordTooLong) {
				api.WriteError(w, http.StatusBadRequest, "password length exceeds 72 bytes")
				return
			}

			api.WriteError(w, http.StatusInternalServerError, err.Error())
			return
		}

		newPassword = [60]byte(hashedPassword)
	}

	user, err := s.db.UpdateUserById(r.Context(), int32(id), newUser.FirstName, newUser.LastName, newUser.Email, newPassword)
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusCreated, user)
}

// @Summary      Delete user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int32  true  "User ID"
// @Success      204
// @Failure      400  {object}  api.ErrorResponse
// @Router       /users/{id} [delete]
func (s UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idS := r.PathValue("id")
	if idS == "" {
		panic("empty id path value")
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if id > math.MaxInt32 {
		api.WriteError(w, http.StatusBadRequest, "id cannot be greater than 2147483647")
		return
	}

	err = s.db.DeleteUserById(r.Context(), int32(id))
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte{})
}

package handlers

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/5aradise/media-content-api/src/internal/types"
	"github.com/5aradise/media-content-api/src/pkg/api"
	"github.com/5aradise/media-content-api/src/pkg/valid"
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

// @Summary      Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input body      CreateUserRequest true "User info"
// @Success      201   {object}  types.User
// @Failure      400   {object}  api.ErrorResponse
// @Failure      500   {object}  api.ErrorResponse
// @Router       /users [post]
func (s UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := api.DecodeJSON(r, &req); err != nil {
		api.WriteErrorf(w, http.StatusBadRequest, err.Error())
		return
	}

	if req.Email == "" {
		api.WriteErrorf(w, http.StatusBadRequest, "empty email")
		return
	}
	if req.FirstName == "" {
		api.WriteErrorf(w, http.StatusBadRequest, "empty first name")
		return
	}
	if req.LastName == "" {
		api.WriteErrorf(w, http.StatusBadRequest, "empty last name")
		return
	}
	if req.Password == "" {
		api.WriteErrorf(w, http.StatusBadRequest, "empty password")
		return
	}
	if !valid.Email(req.Email) {
		api.WriteErrorf(w, http.StatusBadRequest, "invalid email")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		if errors.Is(err, bcrypt.ErrPasswordTooLong) {
			api.WriteErrorf(w, http.StatusBadRequest, "password length exceeds 72 characters")
			return
		}

		api.WriteErrorf(w, http.StatusInternalServerError, "CreateUser: %v", err)
		return
	}

	user, err := s.db.CreateUser(r.Context(), req.FirstName, req.LastName, req.Email, [60]byte(hashedPassword))
	if err != nil {
		if errors.Is(err, types.ErrNameTooLong) {
			api.WriteErrorf(w, http.StatusBadRequest, "first or last name cannot be longer than %d characters", types.NameMaxLen)
			return
		}
		if errors.Is(err, types.ErrEmailTooLong) {
			api.WriteErrorf(w, http.StatusBadRequest, "email cannot be longer than %d characters", types.EmailMaxLen)
			return
		}
		if errors.Is(err, types.ErrUserEmailExists) {
			api.WriteErrorf(w, http.StatusBadRequest, "user with this email already exists")
			return
		}

		api.WriteErrorf(w, http.StatusInternalServerError, "CreateUser: %v", err)
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
		api.WriteErrorf(w, http.StatusInternalServerError, "ListUsers: %v", err)
		return
	}

	api.WriteJSON(w, http.StatusOK, ListUsersResponse{users})
}

// @Summary      Get user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  types.User
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/{id} [get]
func (s UserService) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := getIdPathValue(r)
	if err != nil {
		api.WriteErrorf(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := s.db.GetUserById(r.Context(), id)
	if err != nil {
		if errors.Is(err, types.ErrUserIdNotExists) {
			api.WriteErrorf(w, http.StatusNotFound, "user with this id was not found")
			return
		}

		api.WriteErrorf(w, http.StatusInternalServerError, "GetUser: %v", err)
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
// @Success      200  {object}  types.User
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/{id} [put]
func (s UserService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := getIdPathValue(r)
	if err != nil {
		api.WriteErrorf(w, http.StatusBadRequest, err.Error())
		return
	}

	dbResCh := make(chan struct {
		user types.User
		err  error
	}, 1)
	go func() {
		newUser, err := s.db.GetUserById(r.Context(), id)
		dbResCh <- struct {
			user types.User
			err  error
		}{newUser, err}
	}()

	var req UpdateUserRequest
	if err := api.DecodeJSON(r, &req); err != nil {
		api.WriteErrorf(w, http.StatusBadRequest, err.Error())
		return
	}

	dbRes := <-dbResCh
	newUser := dbRes.user
	err = dbRes.err
	if err != nil {
		if errors.Is(err, types.ErrUserIdNotExists) {
			api.WriteErrorf(w, http.StatusNotFound, "user with this id was not found")
			return
		}
		api.WriteErrorf(w, http.StatusInternalServerError, "UpdateUser: %v", err)
		return
	}

	if req.Email != "" {
		if !valid.Email(req.Email) {
			api.WriteErrorf(w, http.StatusBadRequest, "invalid email")
			return
		}
		newUser.Email = req.Email
	}
	if req.FirstName != "" {
		newUser.FirstName = req.FirstName
	}
	if req.LastName != "" {
		newUser.LastName = req.LastName
	}
	newPassword := [60]byte([]byte(newUser.Password))
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			if errors.Is(err, bcrypt.ErrPasswordTooLong) {
				api.WriteErrorf(w, http.StatusBadRequest, "password length exceeds 72 bytes")
				return
			}

			api.WriteErrorf(w, http.StatusInternalServerError, "UpdateUser: %v", err)
			return
		}

		newPassword = [60]byte(hashedPassword)
	}

	user, err := s.db.UpdateUserById(r.Context(), int32(id), newUser.FirstName, newUser.LastName, newUser.Email, newPassword)
	if err != nil {
		if errors.Is(err, types.ErrNameTooLong) {
			api.WriteErrorf(w, http.StatusBadRequest, "first or last name cannot be longer than %d characters", types.NameMaxLen)
			return
		}
		if errors.Is(err, types.ErrEmailTooLong) {
			api.WriteErrorf(w, http.StatusBadRequest, "email cannot be longer than %d characters", types.EmailMaxLen)
			return
		}
		if errors.Is(err, types.ErrUserEmailExists) {
			api.WriteErrorf(w, http.StatusBadRequest, "this email is already in use")
			return
		}

		api.WriteErrorf(w, http.StatusInternalServerError, "UpdateUser: %v", err)
		return
	}

	api.WriteJSON(w, http.StatusOK, user)
}

// @Summary      Delete user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int32  true  "User ID"
// @Success      204
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /users/{id} [delete]
func (s UserService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := getIdPathValue(r)
	if err != nil {
		api.WriteErrorf(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.DeleteUserById(r.Context(), id)
	if err != nil {
		api.WriteErrorf(w, http.StatusInternalServerError, "DeleteUser: %v", err)
		return
	}

	api.WriteNoContent(w)
}

func getIdPathValue(r *http.Request) (int32, error) {
	idS := r.PathValue("id")
	if idS == "" {
		panic("empty id path value")
	}

	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, errors.New("invalid number")
	}

	if id > math.MaxInt32 {
		return 0, fmt.Errorf("id cannot be greater than %d", math.MaxInt32)
	}

	return int32(id), nil
}

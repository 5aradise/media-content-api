package handlers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/5aradise/media-content-api/src/internal/types"
	"github.com/5aradise/media-content-api/src/pkg/api"
)

type MediaContentStorage interface {
	CreateMediaContent(ctx context.Context, title string, desc sql.NullString, body string, mcType types.MediaContentType, userId int32) (types.MediaContent, error)
	ListMediaContent(ctx context.Context) ([]types.MediaContent, error)
	GetMediaContentById(ctx context.Context, id int32) (types.MediaContent, error)
	ListMediaContentByUserId(ctx context.Context, id int32) ([]types.MediaContent, error)
	DeleteMediaContentById(ctx context.Context, id int32) error
}

type MediaContentService struct {
	db MediaContentStorage
}

func NewMediaContentService(db MediaContentStorage) *MediaContentService {
	return &MediaContentService{
		db: db,
	}
}

type CreateMediaContentRequest struct {
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Body        string                 `json:"body"`
	Type        types.MediaContentType `json:"type"`
	UserId      int32                  `json:"user_id"`
}

// @Summary      Create media content
// @Tags         media_content
// @Accept       json
// @Produce      json
// @Param        input body      CreateMediaContentRequest true "Media content info"
// @Success      201   {object}  types.MediaContent
// @Failure      400   {object}  api.ErrorResponse
// @Failure      500   {object}  api.ErrorResponse
// @Router       /media_content  [post]
func (s MediaContentService) CreateMediaContent(w http.ResponseWriter, r *http.Request) {
	var req CreateMediaContentRequest
	if err := api.DecodeJSON(r, &req); err != nil {
		if errors.Is(err, types.ErrWrongMediaContentTypeString) {
			api.WriteErrorf(w, http.StatusBadRequest, "wrong media content type, available only: text | image | audio | video")
			return
		}
		api.WriteErrorf(w, http.StatusBadRequest, err.Error())
		return
	}

	if req.Title == "" {
		api.WriteErrorf(w, http.StatusBadRequest, "empty title")
		return
	}
	if req.Body == "" {
		api.WriteErrorf(w, http.StatusBadRequest, "empty body")
		return
	}
	if req.Type == "" {
		api.WriteErrorf(w, http.StatusBadRequest, "empty type")
		return
	}
	var desc sql.NullString
	if req.Description != "" {
		desc.String = req.Description
		desc.Valid = true
	}

	mc, err := s.db.CreateMediaContent(r.Context(), req.Title, desc, req.Body, req.Type, req.UserId)
	if err != nil {
		if errors.Is(err, types.ErrTitleTooLong) {
			api.WriteErrorf(w, http.StatusBadRequest, "title cannot be longer than %d characters", types.TitleMaxLen)
			return
		}
		if errors.Is(err, types.ErrDescriptionTooLong) {
			api.WriteErrorf(w, http.StatusBadRequest, "description cannot be longer than %d characters", types.DescriptionMaxLen)
			return
		}
		if errors.Is(err, types.ErrUserIdNotExists) {
			api.WriteErrorf(w, http.StatusBadRequest, "user with this id not exists")
			return
		}

		api.WriteErrorf(w, http.StatusInternalServerError, "CreateMediaContent: %v", err)
		return
	}
	api.WriteJSON(w, http.StatusCreated, mc)
}

type ListMediaContentResponse struct {
	Users []types.MediaContent `json:"media_content"`
}

// @Summary      List media content
// @Tags         media_content
// @Accept       json
// @Produce      json
// @Param        user_id query     int false "User ID"
// @Success      200     {object}  ListMediaContentResponse
// @Failure      400     {object}  api.ErrorResponse
// @Failure      500     {object}  api.ErrorResponse
// @Router       /media_content [get]
func (s MediaContentService) ListMediaContent(w http.ResponseWriter, r *http.Request) {
	userIdS := r.URL.Query().Get("user_id")

	var mc []types.MediaContent
	var err error
	if userIdS == "" {
		mc, err = s.db.ListMediaContent(r.Context())
		if err != nil {
			api.WriteErrorf(w, http.StatusInternalServerError, "ListMediaContent: %v", err)
			return
		}
	} else {
		userId, err := getId(userIdS)
		if err != nil {
			api.WriteErrorf(w, http.StatusBadRequest, err.Error())
			return
		}

		mc, err = s.db.ListMediaContentByUserId(r.Context(), userId)
		if err != nil {
			api.WriteErrorf(w, http.StatusInternalServerError, "ListMediaContent: %v", err)
			return
		}
	}

	api.WriteJSON(w, http.StatusOK, ListMediaContentResponse{mc})
}

// @Summary      Get media content by id
// @Tags         media_content
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Media content ID"
// @Success      200  {object}  types.MediaContent
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /media_content/{id} [get]
func (s MediaContentService) GetMediaContent(w http.ResponseWriter, r *http.Request) {
	idS := r.PathValue("id")
	if idS == "" {
		panic("empty id path value")
	}

	id, err := getId(idS)
	if err != nil {
		api.WriteErrorf(w, http.StatusBadRequest, err.Error())
		return
	}

	mc, err := s.db.GetMediaContentById(r.Context(), id)
	if err != nil {
		if errors.Is(err, types.ErrMediaContentIdNotExists) {
			api.WriteErrorf(w, http.StatusNotFound, "media content with this id was not found")
			return
		}

		api.WriteErrorf(w, http.StatusInternalServerError, "GetMediaContent: %v", err)
		return
	}

	api.WriteJSON(w, http.StatusOK, mc)
}

// @Summary      Delete media content by id
// @Tags         media_content
// @Accept       json
// @Produce      json
// @Param        id   path      int32  true  "Media content ID"
// @Success      204
// @Failure      400  {object}  api.ErrorResponse
// @Failure      500  {object}  api.ErrorResponse
// @Router       /media_content/{id} [delete]
func (s MediaContentService) DeleteMediaContent(w http.ResponseWriter, r *http.Request) {
	idS := r.PathValue("id")
	if idS == "" {
		panic("empty id path value")
	}

	id, err := getId(idS)
	if err != nil {
		api.WriteErrorf(w, http.StatusBadRequest, err.Error())
		return
	}

	err = s.db.DeleteMediaContentById(r.Context(), id)
	if err != nil {
		api.WriteErrorf(w, http.StatusInternalServerError, "DeleteMediaContent: %v", err)
		return
	}

	api.WriteNoContent(w)
}

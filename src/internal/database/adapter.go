package database

import (
	"context"
	"database/sql"

	"github.com/5aradise/media-content-api/src/internal/types"
	"github.com/lib/pq"
)

func userToTypes(dbUser User) types.User {
	return types.User{
		Id:        dbUser.ID,
		FirstName: dbUser.FirstName,
		LastName:  dbUser.LastName,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
	}
}

func mediaContentToTypes(dbMediaContent MediaContent) types.MediaContent {
	contentType, _ := types.NewMediaContentType(dbMediaContent.ContentType)
	return types.MediaContent{
		Id:          dbMediaContent.ID,
		Title:       dbMediaContent.Title,
		Description: dbMediaContent.Description.String,
		Body:        dbMediaContent.Body,
		ContentType: contentType,
		CreatedAt:   dbMediaContent.CreatedAt,
		UserID:      dbMediaContent.UserID,
	}
}

type DB struct {
	q *Queries
}

func Create(db DBTX) *DB {
	return &DB{New(db)}
}

func (db DB) CreateUser(ctx context.Context, firstName, lastName, email string, hashedPassword [types.PasswordMaxLen]byte) (types.User, error) {
	if len(firstName) > types.NameMaxLen {
		return types.User{}, types.ErrNameTooLong
	}
	if len(lastName) > types.NameMaxLen {
		return types.User{}, types.ErrNameTooLong
	}
	if len(email) > types.EmailMaxLen {
		return types.User{}, types.ErrEmailTooLong
	}

	u, err := db.q.CreateUser(ctx, CreateUserParams{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashedPassword[:]),
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" {
				return types.User{}, types.ErrUserEmailExists
			}
		}
		return types.User{}, err
	}
	return userToTypes(u), nil
}

func (db DB) ListUsers(ctx context.Context) ([]types.User, error) {
	dbUsers, err := db.q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]types.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = userToTypes(dbUser)
	}
	return users, nil
}

func (db DB) GetUserById(ctx context.Context, id int32) (types.User, error) {
	u, err := db.q.GetUserById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, types.ErrUserIdNotExists
		}
		return types.User{}, err
	}
	return userToTypes(u), nil
}

func (db DB) UpdateUserById(ctx context.Context, id int32, firstName, lastName, email string, hashedPassword [60]byte) (types.User, error) {
	if len(firstName) > types.NameMaxLen {
		return types.User{}, types.ErrNameTooLong
	}
	if len(lastName) > types.NameMaxLen {
		return types.User{}, types.ErrNameTooLong
	}
	if len(email) > types.EmailMaxLen {
		return types.User{}, types.ErrEmailTooLong
	}

	u, err := db.q.UpdateUserById(ctx, UpdateUserByIdParams{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashedPassword[:]),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, types.ErrUserIdNotExists
		}
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" {
				return types.User{}, types.ErrUserEmailExists
			}
		}
		return types.User{}, err
	}
	return userToTypes(u), nil
}

func (db DB) DeleteUserById(ctx context.Context, id int32) error {
	err := db.q.DeleteUserById(ctx, id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "foreign_key_violation" {
				return types.ErrUserFKConstraint
			}
		}
	}
	return err
}

func (db DB) CreateMediaContent(
	ctx context.Context, title string, description sql.NullString, body string, mcType types.MediaContentType, userId int32,
) (types.MediaContent, error) {
	if len(title) > types.TitleMaxLen {
		return types.MediaContent{}, types.ErrTitleTooLong
	}
	if description.Valid {
		if len(description.String) > types.DescriptionMaxLen {
			return types.MediaContent{}, types.ErrDescriptionTooLong
		}
	}

	mc, err := db.q.CreateMediaContent(ctx, CreateMediaContentParams{
		Title:       title,
		Description: description,
		Body:        body,
		ContentType: mcType.String(),
		UserID:      userId,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "foreign_key_violation" {
				return types.MediaContent{}, types.ErrUserIdNotExists
			}
		}
		return types.MediaContent{}, err
	}
	return mediaContentToTypes(mc), nil
}

func (db DB) ListMediaContent(ctx context.Context) ([]types.MediaContent, error) {
	dbMediaContent, err := db.q.ListMediaContent(ctx)
	if err != nil {
		return nil, err
	}
	mediaContent := make([]types.MediaContent, len(dbMediaContent))
	for i, dbMediaContent := range dbMediaContent {
		mediaContent[i] = mediaContentToTypes(dbMediaContent)
	}
	return mediaContent, nil
}

func (db DB) ListMediaContentByUserId(ctx context.Context, userID int32) ([]types.MediaContent, error) {
	dbMediaContent, err := db.q.ListMediaContentByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}
	mediaContent := make([]types.MediaContent, len(dbMediaContent))
	for i, dbMediaContent := range dbMediaContent {
		mediaContent[i] = mediaContentToTypes(dbMediaContent)
	}
	return mediaContent, nil
}

func (db DB) GetMediaContentById(ctx context.Context, id int32) (types.MediaContent, error) {
	mc, err := db.q.GetMediaContentById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.MediaContent{}, types.ErrMediaContentIdNotExists
		}
		return types.MediaContent{}, err
	}
	return mediaContentToTypes(mc), nil
}

func (db DB) DeleteMediaContentById(ctx context.Context, id int32) error {
	return db.q.DeleteMediaContentById(ctx, id)
}

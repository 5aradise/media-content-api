package database

import (
	"context"

	"github.com/5aradise/media-content-api/src/internal/types"
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

// func mediaContentToTypes(dbMediaContent MediaContent) types.MediaContent {
// 	contentType, _ := types.NewMediaContentType(dbMediaContent.ContentType)
// 	return types.MediaContent{
// 		Id:          dbMediaContent.ID,
// 		Title:       dbMediaContent.Title,
// 		Description: dbMediaContent.Description.String,
// 		Body:        dbMediaContent.Body,
// 		ContentType: contentType,
// 		CreatedAt:   dbMediaContent.CreatedAt,
// 		UserID:      dbMediaContent.UserID,
// 	}
// }

type DB struct {
	q *Queries
}

func Create(db DBTX) *DB {
	return &DB{New(db)}
}

func (db DB) CreateUser(ctx context.Context, firstName, lastName, email string, hashedPassword [60]byte) (types.User, error) {
	u, err := db.q.CreateUser(ctx, CreateUserParams{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashedPassword[:]),
	})
	if err != nil {
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
		return types.User{}, err
	}
	return userToTypes(u), nil
}

func (db DB) UpdateUserById(ctx context.Context, id int32, firstName, lastName, email string, hashedPassword [60]byte) (types.User, error) {
	u, err := db.q.UpdateUserById(ctx, UpdateUserByIdParams{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashedPassword[:]),
	})
	if err != nil {
		return types.User{}, err
	}
	return userToTypes(u), nil
}

func (db DB) DeleteUserById(ctx context.Context, id int32) error {
	return db.q.DeleteUserById(ctx, id)
}

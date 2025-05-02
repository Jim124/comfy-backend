package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jim124/comfy-backend/database"
	"github.com/jim124/comfy-backend/models"
	"github.com/jim124/comfy-backend/utils"
)

type UserRepository interface {
	Register(ctx context.Context, identifier string, password string) (*models.User, error)
	GetUser(ctx context.Context, identifier string, password string) (string, error)
}

func NewUserRepository() {}

type UserImpl struct{}

func (u *UserImpl) Register(ctx context.Context, user models.User) error {
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	uuid := uuid.New()
	_, err = database.DB.ExecContext(ctx, "insert into users(id,identifier,password)values($1,$2,$3)", uuid.String(), user.Identifier, hashPassword)
	return err
}

func (u *UserImpl) GetUser(ctx context.Context, identifier string) (*models.User, error) {
	row := database.DB.QueryRowContext(ctx, "select id,identifier from users where identifier=$1", identifier)
	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Identifier); err != nil {
		return nil, err
	}
	return user, nil
}

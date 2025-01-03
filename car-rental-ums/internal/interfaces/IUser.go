package interfaces

import (
	"car-rental-ums/internal/models"
	"context"
)

type IUserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetUserByToken(ctx context.Context, token string) (*models.User, error)
	GetUserByRefreshToken(ctx context.Context, refreshToken string) (*models.User, error)
	GetUserByEmailVerifyToken(ctx context.Context, emailVerifyToken string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateProfile(ctx context.Context, user *models.User) error
	InsertNewUserSession(ctx context.Context, user *models.UserSession) error
}

package interfaces

import "car-rental-ums/internal/models"

type IUserRepository interface {
	Create(user *models.User) error
	GetUserByToken(token string) (*models.User, error)
	GetUserByRefreshToken(refreshToken string) (*models.User, error)
	GetUserByEmailVerifyToken(emailVerifyToken string) (*models.User, error)
	UpdateProfile(user *models.User) error
}

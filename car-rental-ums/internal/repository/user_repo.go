package repository

import (
	"car-rental-ums/internal/models"
	"context"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) Create(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) GetUserByToken(ctx context.Context, token string) (*models.User, error) {
	var (
		user = &models.User{}
	)
	r.DB.WithContext(ctx).Where("token = ?", token).Last(user)
	return user, nil
}

func (r *UserRepo) GetUserByRefreshToken(ctx context.Context, refreshToken string) (*models.User, error) {
	var (
		user = &models.User{}
	)
	r.DB.WithContext(ctx).Where("token = ?", refreshToken).Last(user)
	return user, nil
}

func (r *UserRepo) GetUserByEmailVerifyToken(ctx context.Context, emailVerifyToken string) (*models.User, error) {
	var (
		user = &models.User{}
	)
	r.DB.WithContext(ctx).Where("verification_token = ?", emailVerifyToken).Last(user)
	return user, nil
}

func (r *UserRepo) GetUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error) {
	var (
		userSession = &models.UserSession{}
	)
	r.DB.WithContext(ctx).Where("token = ?", token).Last(userSession)
	return userSession, nil
}

func (r *UserRepo) UpdateProfile(ctx context.Context, user *models.User) error {
	return r.DB.WithContext(ctx).Save(user).Error
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var (
		user = &models.User{}
	)
	r.DB.WithContext(ctx).Where("email = ?", email).Last(user)
	return user, nil
}

func (r *UserRepo) InsertNewUserSession(ctx context.Context, user *models.UserSession) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r *UserRepo) DeleteUserSession(ctx context.Context, token string) error {
	return r.DB.WithContext(ctx).Where("token = ?", token).Delete(&models.UserSession{}).Error
}

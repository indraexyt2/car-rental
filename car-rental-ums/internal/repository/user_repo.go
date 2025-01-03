package repository

import (
	"car-rental-ums/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (r *UserRepo) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepo) GetUserByToken(token string) (*models.User, error) {
	var (
		user = &models.User{}
	)
	r.DB.Where("token = ?", token).Last(user)
	return user, nil
}

func (r *UserRepo) GetUserByRefreshToken(refreshToken string) (*models.User, error) {
	var (
		user = &models.User{}
	)
	r.DB.Where("token = ?", refreshToken).Last(user)
	return user, nil
}

func (r *UserRepo) GetUserByEmailVerifyToken(emailVerifyToken string) (*models.User, error) {
	var (
		user = &models.User{}
	)
	r.DB.Where("verification_token = ?", emailVerifyToken).Last(user)
	return user, nil
}

func (r *UserRepo) UpdateProfile(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var (
		user = &models.User{}
	)
	r.DB.Where("email = ?", email).Last(user)
	return user, nil
}

func (r *UserRepo) InsertNewUserSession(user *models.UserSession) error {
	return r.DB.Create(user).Error
}

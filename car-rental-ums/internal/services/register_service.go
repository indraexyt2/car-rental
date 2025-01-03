package services

import (
	"car-rental-ums/internal/interfaces"
	"car-rental-ums/internal/models"
	"crypto/rand"
	"encoding/base64"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepository interfaces.IUserRepository
}

func (s *RegisterService) Register(request *models.User) (*models.User, error) {
	// hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, " failed to hash password")
	}

	// generate token email verify
	tokenByte := make([]byte, 32)
	_, err = rand.Read(tokenByte)
	if err != nil {
		return nil, errors.Wrap(err, " failed to generate token email verify")
	}
	tokenEmailVerify := base64.URLEncoding.EncodeToString(tokenByte)

	// set role
	if request.Role == "" {
		request.Role = "user"
	}

	// save user
	request.VerificationToken = tokenEmailVerify
	request.Password = string(hashPassword)

	err = s.UserRepository.Create(request)
	if err != nil {
		return nil, errors.Wrap(err, " failed to create user")
	}

	// response
	resp := request
	resp.Password = ""
	resp.VerificationToken = ""

	return resp, nil
}

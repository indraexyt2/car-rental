package services

import (
	"car-rental-ums/helpers"
	"car-rental-ums/internal/interfaces"
	"car-rental-ums/internal/models"
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, request *models.LoginRequest) (*models.LoginResponse, error) {
	var (
		resp = &models.LoginResponse{}
	)

	// get user
	user, err := s.UserRepo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by email")
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.Wrap(err, "password not match")
	}

	// validate user is verified
	if !user.IsVerified {
		return nil, errors.New("user is not verified")
	}

	// generate token
	token, err := helpers.GenerateJWTToken(ctx, user.FirstName, user.LastName, user.Email, user.Role, "token")
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateJWTToken(ctx, user.FirstName, user.LastName, user.Email, user.Role, "refresh_token")
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate refresh token")
	}

	// save user session
	userSession := &models.UserSession{
		UserID:       user.ID,
		Token:        token,
		RefreshToken: refreshToken,
	}

	err = s.UserRepo.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert new user session")
	}

	// response
	resp.UserID = user.ID
	resp.Email = user.Email
	resp.Token = token
	resp.RefreshToken = refreshToken

	return resp, nil
}

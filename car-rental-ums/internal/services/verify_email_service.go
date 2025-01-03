package services

import (
	"car-rental-ums/internal/interfaces"
	"context"
	"github.com/pkg/errors"
)

type EmailVerifyService struct {
	UserRepo interfaces.IUserRepository
}

func (s *EmailVerifyService) EmailVerify(ctx context.Context, emailVerifyToken string) error {
	// get user
	user, err := s.UserRepo.GetUserByEmailVerifyToken(ctx, emailVerifyToken)
	if err != nil {
		return errors.Wrap(err, " failed to get user by email verify token")
	}

	// update user
	user.IsVerified = true
	user.VerificationToken = ""

	//save user
	err = s.UserRepo.UpdateProfile(ctx, user)
	if err != nil {
		return errors.Wrap(err, " failed to update profile")
	}

	return nil
}

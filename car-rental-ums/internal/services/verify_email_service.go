package services

import (
	"car-rental-ums/internal/interfaces"
	"github.com/pkg/errors"
)

type EmailVerifyService struct {
	UserRepo interfaces.IUserRepository
}

func (s *EmailVerifyService) EmailVerify(emailVerifyToken string) error {
	// get user
	user, err := s.UserRepo.GetUserByEmailVerifyToken(emailVerifyToken)
	if err != nil {
		return errors.Wrap(err, " failed to get user by email verify token")
	}

	// update user
	user.IsVerified = true
	user.VerificationToken = ""

	//save user
	err = s.UserRepo.UpdateProfile(user)
	if err != nil {
		return errors.Wrap(err, " failed to update profile")
	}

	return nil
}

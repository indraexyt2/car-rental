package services

import (
	"car-rental-ums/internal/interfaces"
	"car-rental-ums/internal/models"
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/pkg/errors"
)

type ResendEmailVerifyService struct {
	UserRepo interfaces.IUserRepository
}

func (s *ResendEmailVerifyService) ResendEmailVerify(ctx context.Context, email *models.ResendEmailVerifyRequest) error {
	// get user
	user, err := s.UserRepo.GetUserByEmail(ctx, email.Email)
	if err != nil {
		return errors.Wrap(err, " failed to get user by email")
	}

	// check user is verified
	if user.IsVerified {
		return errors.New("user is already verified")
	}

	// generate token email verify
	tokenByte := make([]byte, 32)
	_, err = rand.Read(tokenByte)
	if err != nil {
		return errors.Wrap(err, "failed to generate token email verify")
	}
	tokenEmailVerify := base64.URLEncoding.EncodeToString(tokenByte)

	// update user
	user.VerificationToken = tokenEmailVerify
	err = s.UserRepo.UpdateProfile(ctx, user)
	if err != nil {
		return errors.Wrap(err, " failed to update token email verify")
	}

	return nil
}

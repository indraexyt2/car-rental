package services

import (
	"car-rental-ums/internal/interfaces"
	"context"
	"github.com/pkg/errors"
)

type LogoutService struct {
	UserRepository interfaces.IUserRepository
}

func (s *LogoutService) Logout(ctx context.Context, token string) error {
	// delete user session
	err := s.UserRepository.DeleteUserSession(ctx, token)
	if err != nil {
		return errors.Wrap(err, "failed to delete user session")
	}
	return nil
}

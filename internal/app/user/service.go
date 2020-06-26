package user

import (
	"context"
	"errors"

	"github.com/PhongVX/taskmanagement/internal/pkg/hash/password"
)

// NewService return a new user service
func NewService(r RepositoryInterface) *Service {
	srv := &Service{
		repo: r,
	}
	return srv
}

func (s *Service) Insert(ctx context.Context, u *User) error {
	_, err := s.repo.FindByUserIdentity(ctx, u)
	if err == nil {
		return errors.New("User already existed")
	}
	//TODO Validation password

	//Hash Password
	passwordHashed, err := password.HashPassword(u.Password)
	if err != nil {
		return errors.New("Password invalid")
	}
	u.Password = string(passwordHashed)
	return s.repo.Insert(ctx, u)
}

func (s *Service) FindAll(ctx context.Context, rO FindingRequestObject) ([]*User, error) {
	return s.repo.FindAll(ctx, rO)
}

//Identity: userName, email, _id
func (s *Service) FindByIdentity(ctx context.Context, identity string) (*User, error) {
	return s.repo.FindByIdentity(ctx, identity)
}

func (s *Service) FindByUserIdentity(ctx context.Context, u *User) (*User, error) {
	return s.repo.FindByUserIdentity(ctx, u)
}

// Delete a user by user id
func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// Update a user
func (s *Service) Update(ctx context.Context, u *User) error {
	return s.repo.Update(ctx, u)
}

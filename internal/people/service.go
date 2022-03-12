package people

import (
	"context"
	"fmt"
)

type Service interface {
	List(ctx context.Context) ([]Person, error)
	Add(ctx context.Context, person Person) error
	Update(ctx context.Context, person Person) error
	Remove(ctx context.Context, id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) List(ctx context.Context) ([]Person, error) {
	people, err := s.repo.GetPeople(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get people: %w", err)
	}

	return people, nil
}

func (s *service) Add(ctx context.Context, person Person) error {
	if err := s.repo.InsertPeople(ctx, person); err != nil {
		return fmt.Errorf("failed to insert people: %w", err)
	}

	return nil
}

func (s *service) Update(ctx context.Context, person Person) error {
	if err := s.repo.UpdatePeople(ctx, person); err != nil {
		return fmt.Errorf("failed to update people: %w", err)
	}

	return nil
}

func (s *service) Remove(ctx context.Context, id string) error {
	if err := s.repo.DeletePeople(ctx, id); err != nil {
		return fmt.Errorf("failed to delete people: %w", err)
	}

	return nil
}

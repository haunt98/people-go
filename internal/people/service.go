package people

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/make-go-great/date-go"
	"github.com/segmentio/ksuid"
)

var (
	ErrEmptyID   = errors.New("empty id")
	ErrEmptyName = errors.New("empty name")
)

type Service interface {
	List(ctx context.Context) ([]*Person, error)
	Get(ctx context.Context, id string) (*Person, error)
	Add(ctx context.Context, person *Person) error
	Update(ctx context.Context, person *Person) error
	Remove(ctx context.Context, id string) error
	Drop(ctx context.Context) error
}

type service struct {
	repo     Repository
	location *time.Location
}

func NewService(repo Repository, location *time.Location) Service {
	return &service{
		repo:     repo,
		location: location,
	}
}

func (s *service) List(ctx context.Context) ([]*Person, error) {
	people, err := s.repo.GetPeople(ctx)
	if err != nil {
		return nil, fmt.Errorf("repo: failed to get people: %w", err)
	}

	return people, nil
}

func (s *service) Get(ctx context.Context, id string) (*Person, error) {
	if id == "" {
		return nil, ErrEmptyID
	}

	person, err := s.repo.GetPerson(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repo: failed to get person: %w", err)
	}

	return person, nil
}

func (s *service) Add(ctx context.Context, person *Person) error {
	if person.ID == "" {
		// Be careful, it can panic
		person.ID = ksuid.New().String()
	}

	if err := validatePerson(person); err != nil {
		return err
	}

	if person.Birthday != "" {
		var err error
		person.Birthday, err = date.ToDefaultDate(person.Birthday, s.location)
		if err != nil {
			return fmt.Errorf("date: failed to default date [%s]: %w", person.Birthday, err)
		}
	}

	if err := s.repo.InsertPeople(ctx, person); err != nil {
		return fmt.Errorf("repo: failed to insert people: %w", err)
	}

	return nil
}

func (s *service) Update(ctx context.Context, person *Person) error {
	if err := validatePerson(person); err != nil {
		return err
	}

	if person.Birthday != "" {
		var err error
		person.Birthday, err = date.ToDefaultDate(person.Birthday, s.location)
		if err != nil {
			return fmt.Errorf("date: failed to default date [%s]: %w", person.Birthday, err)
		}
	}

	if err := s.repo.UpdatePeople(ctx, person); err != nil {
		return fmt.Errorf("repo: failed to update people: %w", err)
	}

	return nil
}

func (s *service) Remove(ctx context.Context, id string) error {
	if id == "" {
		return ErrEmptyID
	}

	if err := s.repo.DeletePeople(ctx, id); err != nil {
		return fmt.Errorf("repo: failed to delete people: %w", err)
	}

	return nil
}

func (s *service) Drop(ctx context.Context) error {
	if err := s.repo.DropPeople(ctx); err != nil {
		return fmt.Errorf("repo: failed to drop people: %w", err)
	}

	return nil
}

func validatePerson(person *Person) error {
	if person.ID == "" {
		return ErrEmptyID
	}

	if person.Name == "" {
		return ErrEmptyName
	}

	return nil
}

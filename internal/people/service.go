package people

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	List(ctx context.Context) ([]Person, error)
	Get(ctx context.Context, id string) (Person, error)
	Add(ctx context.Context, person Person) error
	Update(ctx context.Context, person Person) error
	Remove(ctx context.Context, id string) error
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

func (s *service) List(ctx context.Context) ([]Person, error) {
	people, err := s.repo.GetPeople(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get people: %w", err)
	}

	for _, person := range people {
		if person.Birthday == "" {
			continue
		}

		person.Birthday, err = dateToOutput(person.Birthday, s.location)
		if err != nil {
			return nil, fmt.Errorf("failed to output date %s: %w", person.Birthday, err)
		}
	}

	return people, nil
}

func (s *service) Get(ctx context.Context, id string) (Person, error) {
	if id == "" {
		return Person{}, errors.New("empty id")
	}

	person, err := s.repo.GetPerson(ctx, id)
	if err != nil {
		return Person{}, fmt.Errorf("failed to get person: %w", err)
	}

	if person.Birthday != "" {
		person.Birthday, err = dateToOutput(person.Birthday, s.location)
		if err != nil {
			return Person{}, fmt.Errorf("failed to output date %s: %w", person.Birthday, err)
		}
	}

	return person, nil
}

func (s *service) Add(ctx context.Context, person Person) error {
	person.ID = uuid.NewString()

	if err := validatePerson(person); err != nil {
		return err
	}

	if person.Birthday != "" {
		var err error
		person.Birthday, err = dateFromInput(person.Birthday, s.location)
		if err != nil {
			return fmt.Errorf("failed to input date %s: %w", person.Birthday, err)
		}
	}

	if err := s.repo.InsertPeople(ctx, person); err != nil {
		return fmt.Errorf("failed to insert people: %w", err)
	}

	return nil
}

func (s *service) Update(ctx context.Context, person Person) error {
	if err := validatePerson(person); err != nil {
		return err
	}

	if err := s.repo.UpdatePeople(ctx, person); err != nil {
		return fmt.Errorf("failed to update people: %w", err)
	}

	return nil
}

func (s *service) Remove(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("empty id")
	}

	if err := s.repo.DeletePeople(ctx, id); err != nil {
		return fmt.Errorf("failed to delete people: %w", err)
	}

	return nil
}

func validatePerson(person Person) error {
	if person.ID == "" {
		return errors.New("empty id")
	}

	if person.Name == "" {
		return errors.New("empty name")
	}

	return nil
}

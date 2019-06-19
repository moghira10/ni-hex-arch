package checkin

import (
	"errors"
)

type Service interface {
	AddCheckIn(CheckIn) error
}

type Repository interface {
	AddCheckIn(CheckIn) error
}

type service struct {
	chkin Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddCheckIn(c CheckIn) error {
	if err := s.chkin.AddCheckIn(c); err != nil {
		// Error Handler
		return errors.New("Unable to perform checkin")
	}
	return nil
}

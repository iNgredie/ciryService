package services

import city "cityService/internal/app/entities"

type Repository interface {
	GetAllCities() city.Cities
}

type Service struct {
	r Repository
}

func New(r Repository) *Service {
	return &Service{
		r: r,
	}
}

func (s Service) GetCities() city.Cities {
	c := s.r.GetAllCities()
	return c
}

package repositories

import (
	city "cityService/internal/app/entities"
)

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r Repository) GetAllCities() city.Cities {
	cities := city.Cities{
		city.City{
			ID:   1,
			Name: "Test name",
		},
	}
	return cities
}

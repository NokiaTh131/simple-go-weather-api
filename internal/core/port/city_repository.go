package port

import "weather/internal/core/domain"

type CityRepository interface {
	Create(city *domain.City) error
	GetById(id uint) (*domain.City,error)
	Update(city *domain.City) error
	Delete(id uint) error
	List() ([]domain.City,error)
}
package service

import (
	"weather/internal/core/domain"
	"weather/internal/core/port"
)

type CityService struct {
	repo port.CityRepository
}

func NewCityService(repo port.CityRepository) *CityService  {
	return &CityService{repo: repo}
}

func (s *CityService) CreateCity(city *domain.City) error {
	return s.repo.Create(city)
}

func (s *CityService) GetCity(id uint) (*domain.City,error)  {
	return s.repo.GetById(id)
}

func (s *CityService) UpdateCity(city *domain.City) error  {
	return s.repo.Update(city)
}

func (s *CityService) DeleteCity(id uint) error {
	return s.repo.Delete(id)
}

func (s *CityService) ListCities() ([]domain.City,error)  {
	return s.repo.List()
}
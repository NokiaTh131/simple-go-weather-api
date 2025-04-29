package repository

import (
	"weather/internal/core/domain"
	"weather/internal/core/port"

	"gorm.io/gorm"
)

type CityGormRepository struct {
	db *gorm.DB
}

func NewCityGormRepository(db *gorm.DB) port.CityRepository  {
	return &CityGormRepository{db: db}
}

func (r *CityGormRepository) Create(city *domain.City) error  {
	return r.db.Create(city).Error
}

func (r *CityGormRepository) GetById(id uint) (*domain.City,error)  {
	var city domain.City
	err := r.db.First(&city, id).Error
	return &city,err
}

func (r *CityGormRepository) Delete(id uint) error {
	return r.db.Delete(&domain.City{}, id).Error
}

func (r *CityGormRepository) List() ([]domain.City,error) {
	var cities []domain.City
	err := r.db.Find(&cities).Error
	return cities,err
}

func (r *CityGormRepository) Update(city *domain.City) error {
	return r.db.Save(city).Error
}
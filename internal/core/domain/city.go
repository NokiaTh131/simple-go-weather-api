package domain

type City struct {
	ID        uint `gorm:"primaryKey"`
	Latitude  float32
	Longitude float32
	Name      string
	Country   string
}

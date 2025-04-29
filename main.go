package main

import (
	"fmt"
	"weather/internal/adapter/handler"
	"weather/internal/adapter/repository"
	"weather/internal/core/domain"
	"weather/internal/core/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nokiath131"
	password = "12345678"
	dbname   = "mydb"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect to database")
  }
  // Migrate the schema
  db.AutoMigrate(&domain.City{})
  cityRepo := repository.NewCityGormRepository(db)
  cityService := service.NewCityService(cityRepo)
  cityHandler := handler.NewCityHandler(cityService)

  r := gin.Default()

  cityRoutes := r.Group("/cities")
	{
		cityRoutes.POST("/", cityHandler.CreateCity)
		cityRoutes.GET("/", cityHandler.ListCities)
		cityRoutes.GET("/:id", cityHandler.GetCity)
		cityRoutes.PUT("/:id", cityHandler.UpdateCity)
		cityRoutes.DELETE("/:id", cityHandler.DeleteCity)
	}

	r.Run(":8080")
}
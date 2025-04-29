// internal/adapter/handler/city_handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"weather/internal/core/domain"
	"weather/internal/core/service"

	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	cityService *service.CityService
}

type CityResponse struct {
    Name    string
    Country string
}

func NewCityHandler(cityService *service.CityService) *CityHandler {
	return &CityHandler{cityService: cityService}
}

func (h *CityHandler) CreateCity(c *gin.Context) {
	var city domain.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest,err.Error())
		return
	}
	err := h.cityService.CreateCity(&city) 
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error" : "Failed to create city"})
	}
	c.JSON(http.StatusCreated,city)
}

func (h *CityHandler) GetCity(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid ID"})
		return
	}
	city, err := h.cityService.GetCity(uint(id))  
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"error" : "Not Found City"})
		return
	}
	latitude := strconv.FormatFloat(float64(city.Latitude), 'f', 2, 32)
	longitude := strconv.FormatFloat(float64(city.Longitude), 'f', 2, 64)
	weather, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude="+ latitude +"&longitude=" + longitude + "&daily=temperature_2m_max,temperature_2m_min&hourly=temperature_2m&models=bom_access_global&timezone=Asia%2FBangkok")
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error" : "Failed to fetch weather data"})
		return
	}
	defer weather.Body.Close()

    // Decode the weather response into a generic map
    var weatherData map[string]interface{}
    if err := json.NewDecoder(weather.Body).Decode(&weatherData); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse weather data"})
        return
    }

    // Prepare the city response
    response := CityResponse{
        Name:    city.Name,
        Country: city.Country,
    }

    // Return the city and weather data as part of the JSON response
    c.JSON(http.StatusOK, gin.H{
        "city":   response,
        "weather": weatherData, // include entire weather data from the API
    })
}

func (h *CityHandler) ListCities(c *gin.Context) {
	city, err := h.cityService.ListCities()
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error": "Failed to fetch cities"})
		return
	}

	var cityResponse []CityResponse
	for _, city := range city {
		cityResponse = append(cityResponse, CityResponse{
			Name:   city.Name,
			Country: city.Country,
		})
	}
	c.JSON(http.StatusOK,cityResponse)
}

func (h *CityHandler) DeleteCity(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid ID"})
		return
	}
	err = h.cityService.DeleteCity(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete city"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *CityHandler) UpdateCity(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid ID"})
		return
	}
	var city domain.City
	if err := c.ShouldBindJSON(&city); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	city.ID = uint(id)
	err = h.cityService.UpdateCity(&city)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to update city"})
		return
	}
	c.JSON(http.StatusOK,city)
}
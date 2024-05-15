package weather

import (
	"errors"
	temperature "go-calculator-tutorial/internal/repository/temperature"
	"strconv"
	"time"
)

type IWeatherService interface {
	GetTemperature() (map[string]string, error);
	TemperatureWithUnit(unit string) (map[string]string, error);
}

type WeatherService struct {
	weatherRepo temperature.IWeatherRepository
}

func NewWeatherService(weather temperature.IWeatherRepository) IWeatherService {
	return WeatherService {
        weatherRepo: weather,
    }
}

func (w WeatherService) GetTemperature() (map[string]string, error) {
	response, err := w.weatherRepo.GetTemperature()
	// Get the current time
    currentTime := time.Now()

	// Save response and current time to database
	w.weatherRepo.CreateTemperatureDatabase(101, response, currentTime)
    
    // Convert current time to Unix timestamp
    timestamp := currentTime.Unix()
    
    // Convert Unix timestamp to local time
    localTime := time.Unix(timestamp, 0)
    
    // Format time to only include hours and minutes
    timeStr := localTime.Format("15:04")

	temperature := strconv.FormatFloat(response - 273.15, 'f', 2, 64)

	result := make(map[string]string)
	result["Time"] = timeStr
	result["Temperature"] = temperature
	if err != nil {
		return nil, err
	}
    return result, nil
}

func (w WeatherService) TemperatureWithUnit(unit string) (map[string]string, error) {
	response, err := w.weatherRepo.GetTemperature()
	// Get the current time
    currentTime := time.Now()

	// Save response and current time to database
	w.weatherRepo.CreateTemperatureDatabase(101, response, currentTime)
    
    // Convert current time to Unix timestamp
    timestamp := currentTime.Unix()
    
    // Convert Unix timestamp to local time
    localTime := time.Unix(timestamp, 0)
    
    // Format time to only include hours and minutes
    timeStr := localTime.Format("15:04")
	result := make(map[string]string)
	result["Time"] = timeStr
	switch (unit) {
		case "C":
			result["Temperature"] = convertCelsius(response) + " C"
		case "F":
			result["Temperature"] = convertFahrenheit(response) + " F"
		case "K":
			result["Temperature"] = convertKelvin(response) + " K"
		default:
			result["Temperature"] = "Bad Request"
	}
	if result["Temperature"] == "Bad Request" {
		return nil, errors.New("bad request")
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}

func convertCelsius(t float64) string {
	temperature := strconv.FormatFloat(t - 273.15, 'f', 2, 64)
	return temperature
}

func convertFahrenheit(t float64) string {
	temperature := strconv.FormatFloat(t - 459.67, 'f', 2, 64)
	return temperature
}

func convertKelvin(t float64) string {
	temperature := strconv.FormatFloat(t, 'f', 2, 64)
	return temperature
}
package httphandler

import (
	service "go-calculator-tutorial/internal/service"
	model "go-calculator-tutorial/package/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IWeatherHandler interface {
	GetTemperature(c echo.Context) (error);
}

type WeatherHandler struct {
	weatherServ 	service.IWeatherService
}

func NewWeatherHandler (s service.IWeatherService) IWeatherHandler { 
	return &WeatherHandler {
        weatherServ: s,
    }
}

func (h *WeatherHandler) GetTemperature(c echo.Context) error {
	res, err := h.weatherServ.GetTemperature()
	if err != nil {
		return c.JSON(http.StatusOK, model.ResponseGenericError())
	}
	return c.JSON(http.StatusOK, model.ResponseSuccess(res))
}
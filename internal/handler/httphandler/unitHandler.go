package httphandler

import (
	service "go-calculator-tutorial/internal/service"
	model "go-calculator-tutorial/package/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUnitHandler interface {
	TemperatureWithUnit(c echo.Context, unit string) (error);
}

type UnitHandler struct {
	unitServ 	service.IWeatherService
}

func NewUnitHandler (s service.IWeatherService) IUnitHandler { 
	return &UnitHandler {
        unitServ: s,
    }
}

func (h *UnitHandler) TemperatureWithUnit(c echo.Context, unit string) error {
	res, err := h.unitServ.TemperatureWithUnit(unit)
	if err != nil {
		if err.Error() == "bad request" {
			return c.JSON(http.StatusOK, model.ResponseBadRequest())
		}
		return c.JSON(http.StatusOK, model.ResponseGenericError())
	}
	return c.JSON(http.StatusOK, model.ResponseSuccess(res))
}
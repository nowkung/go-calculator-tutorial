package httphandler

import (
	service "go-calculator-tutorial/internal/service"
	model "go-calculator-tutorial/package/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IDBHandler interface {
	GetTemperatureByID(c echo.Context, id string) (error);
}

type DBHandler struct {
	DBServ 	service.IWeatherService
}

func NewDBHandler (s service.IWeatherService) IDBHandler { 
	return &DBHandler {
        DBServ: s,
    }
}

func (h *DBHandler) GetTemperatureByID(c echo.Context, id string) error {
	data,_ := strconv.ParseUint(id, 10, 0)
	res, err := h.DBServ.GetTemperatureByID(data)
	if err != nil {
		if err.Error() == "bad request" {
			return c.JSON(http.StatusOK, model.ResponseBadRequest())
		}
		return c.JSON(http.StatusOK, model.ResponseGenericError())
	}
	return c.JSON(http.StatusOK, model.ResponseSuccess(res))
}
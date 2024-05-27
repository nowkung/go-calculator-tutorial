package httphandler

import (
	"go-calculator-tutorial/internal/config"

	"github.com/labstack/echo/v4"
)

type HttpServer struct {
	conf  *config.AppConfig
	server *echo.Echo
	weatherHandler IWeatherHandler
	unitHandler IUnitHandler
	dbHandler IDBHandler
}

func NewHttpServer(
	conf *config.AppConfig,
	server *echo.Echo,
	weatherHandler IWeatherHandler,
	unitHandler IUnitHandler,
	dbHandler IDBHandler,
) *HttpServer {
	httpServer := &HttpServer{
		conf:           conf,
		server:         server,
		weatherHandler: weatherHandler,
		unitHandler:    unitHandler,
		dbHandler: dbHandler,
	}

	httpServer.InitRoutes()

	return httpServer
}

func (s *HttpServer) InitRoutes() {
	e := s.server

	e.GET("/time-temperature", s.weatherHandler.GetTemperature)
	e.GET("/unit-temperature", func(c echo.Context) error {

		unit := c.QueryParam("unit")
        return s.unitHandler.TemperatureWithUnit(c, unit)

    })
	e.GET("/temperature-by-id", func(c echo.Context) error {

		id := c.QueryParam("id")
        return s.dbHandler.GetTemperatureByID(c, id)

    })
}

func (s *HttpServer) Start(address string) error {
	return s.server.Start(address)
}

func (s *HttpServer) Server() *echo.Echo {
	return s.server
}
package main

import (
	"fmt"
	"go-calculator-tutorial/internal/config"
	handler "go-calculator-tutorial/internal/handler/httphandler"
	repo "go-calculator-tutorial/internal/repository/temperature"
	service "go-calculator-tutorial/internal/service"

	"github.com/redis/go-redis/v9"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {

	// iphoneCalculator := calculator.IphoneCalculator{}
	// digitalCalculator := calculator.DigitalCalculator{}

	// com1 := computer.NewComputer(iphoneCalculator)
	// com2 := computer.NewComputer(digitalCalculator)

	// result1 := com1.MultiplyNumber(2, 4)
	// result2 := com2.MultiplyNumber(2, 4)

	// fmt.Println(result1, result2)

	conf := new(config.AppConfig)
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config")   // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	conf.Url.Key = viper.GetString("url.key")
	conf.Url.Location = viper.GetString("url.location")
	conf.Server.Address = viper.GetString("server.address")
	conf.Database.Driver = viper.GetString("database.driver")
	conf.Database.Host = viper.GetString("database.host")
	conf.Database.Port = viper.GetString("database.port")
	conf.Database.User = viper.GetString("database.user")
	conf.Database.Password =  viper.GetString("database.password")
	conf.Database.DBname = viper.GetString("database.dbname")
	conf.Redis.User = viper.GetString("redis.user")
	conf.Redis.Password = viper.GetString("redis.password")
	conf.Redis.Host = viper.GetString("redis.host")
	conf.Redis.Port = viper.GetString("redis.port")

	redisUrl := "redis://"+conf.Redis.User +":"+conf.Redis.Password+"@"+conf.Redis.Host+":"+conf.Redis.Port+"/0?protocol=3"
	opts, err := redis.ParseURL(redisUrl)
    if err != nil {
        panic(err)
    }
    client := redis.NewClient(opts)

	weatherRepo := repo.NewWeatherRepository(conf, client)
	weatherServ := service.NewWeatherService(weatherRepo)
	weatherHandler := handler.NewWeatherHandler(weatherServ)
	unitHandler := handler.NewUnitHandler(weatherServ)
	dbHandler := handler.NewDBHandler(weatherServ)
	
	e := echo.New()
	e.GET("/time-temperature", weatherHandler.GetTemperature)
	e.GET("/unit-temperature", func(c echo.Context) error {

		unit := c.QueryParam("unit")
		fmt.Print(unit)
        return unitHandler.TemperatureWithUnit(c, unit)

    })
	e.GET("/temperature-by-id", func(c echo.Context) error {

		id := c.QueryParam("id")
        return dbHandler.GetTemperatureByID(c, id)

    })
	e.Logger.Fatal(e.Start(":1323"))
}
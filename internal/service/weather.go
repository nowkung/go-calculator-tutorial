package weather

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	temperature "go-calculator-tutorial/internal/repository/temperature"
	"log/slog"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IWeatherService interface {
	GetTemperature() (map[string]string, error);
	TemperatureWithUnit(unit string) (map[string]string, error);
	GetTemperatureByID(id uint64) (map[string]string, error) ;
}

type WeatherService struct {
	weatherRepo temperature.IWeatherRepository
}

type TemperatureEntity struct {
	ID				uint64		`gorm:"primaryKey"`
	Uid				uint64        `gorm:"column:uid"` 
	Temperature 	float64     `gorm:"column:temperature"` 
	Times 			time.Time     `gorm:"column:datetime"` 
}

type SqlLoggerInterface struct {
    logger.Interface
}

func (l SqlLoggerInterface) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
   sql,_ := fc()
   fmt.Printf("%v\n=========================\n", sql)
}

var slogger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func NewWeatherService(weather temperature.IWeatherRepository) IWeatherService {
	return WeatherService {
        weatherRepo: weather,
    }
}

func (w WeatherService) GetTemperature() (map[string]string, error) {
	response, err := w.weatherRepo.GetTemperature()
	slog.SetDefault(slogger)
	slog.Info("Get Temperature", slog.Float64("Temperature", response))
	// Get the current time
    currentTime := time.Now()

	// Save response and current time to database
	CreateTemperatureDatabase(101, response, currentTime, w.weatherRepo.GetDBLocation())

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
		slog.Error("Can't get Temperature", err)
		return nil, err
	}
    return result, nil
}

func (w WeatherService) TemperatureWithUnit(unit string) (map[string]string, error) {
	response, err := w.weatherRepo.GetTemperature()
	slog.SetDefault(slogger)
	slog.Info("Get Temperature", slog.Float64("Temperature", response))
	// Get the current time
    currentTime := time.Now()

	// Save response and current time to database
	CreateTemperatureDatabase(101, response, currentTime, w.weatherRepo.GetDBLocation())
    
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
		slog.Error("Bad Request", err)
		return nil, errors.New("bad request")
	}
	if err != nil {
		slog.Error("Can't get Temperature", err)
		return nil, err
	}
	return result, nil
}



var ctx = context.Background()
func (w WeatherService) GetTemperatureByID(id uint64) (map[string]string, error) {
	slog.SetDefault(slogger)
    key := strconv.FormatUint(id,10)
    rdb := w.weatherRepo.GetRedisClient()
    data := TemperatureEntity{}
	result, err := rdb.Get(ctx, key).Result()
    fmt.Println(err)
    if err != nil {
        dsn := w.weatherRepo.GetDBLocation()
	    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ Logger: &SqlLoggerInterface{}})
	    if err!= nil {
			slog.Error("Database not found", err)
            panic(err.Error())
        }
        result := db.Where("temp_id=?", id).Find(&data)
        if result.Error!= nil {
			slog.Error("Database not found", err)
            panic(result.Error)
        }
        save,_ := json.Marshal(data)
        err = rdb.Set(ctx, key, save, 30*time.Second).Err()
    	if err != nil {
			slog.Error("Can't save to redis", err)
        	panic(err)
    	}
    } else {
        json.Unmarshal([]byte(result), &data)
    }
	ans := make(map[string]string)
	ans["Time"] = data.Times.Format("15:04")
	ans["Temperature"] = strconv.FormatFloat(data.Temperature, 'f', 2, 64)
	fmt.Println(ans)
    return ans, nil
}

func CreateTemperatureDatabase(uid uint64, temp float64, t time.Time, dsn string) {
	slog.SetDefault(slogger)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ Logger: &SqlLoggerInterface{}})
	if err!= nil {
		slog.Error("Database not found", err)
        panic(err.Error())
    }
    data := TemperatureEntity{Uid:uid, Temperature: temp, Times: t}
	result := db.Create(&data) // pass pointer of data to Create
	if result.Error!= nil {
		slog.Error("Database not found", err)
        panic(result.Error)
    }else {
		group := slog.Group("Data","ID", strconv.FormatUint(data.ID, 10),"UID", strconv.FormatUint(data.Uid, 10), "Temperature", strconv.FormatFloat(data.Temperature, 'f', 2, 64), "Time", data.Times.String())
		slog.Info("Database created", group)
	}
}

func (t TemperatureEntity) TableName() string {
    return "time_temperature"
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
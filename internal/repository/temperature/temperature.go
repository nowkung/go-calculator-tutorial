package temperature

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-calculator-tutorial/internal/config"
	"io/ioutil"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IWeatherRepository interface {
	GetTemperature() (float64, error);
    CreateTemperatureDatabase( uint, float64, time.Time);
}


type WeatherRepository struct {
	conf   *config.AppConfig
	Days struct {
        Temp float64 `json:"temp"`
    } `json:"main"`

}

type SqlLoggerInterface struct {
    logger.Interface
}

func (l SqlLoggerInterface) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
   sql,_ := fc()
   fmt.Printf("%v\n=========================\n", sql)
}

type TemperatureEntity struct {
	ID				uint		`gorm:"primaryKey"`
	Uid				uint        `gorm:"column:uid"` 
	Temperature 	float64     `gorm:"column:temperature"` 
	Times 			time.Time     `gorm:"column:datetime"` 
}

func NewWeatherRepository(conf *config.AppConfig) IWeatherRepository {
	return &WeatherRepository{
		conf: conf,
	}
}

func (temp WeatherRepository) GetTemperature() (float64, error) {

	// const apiKey = "913da61969a3190ba9de6b5c094a0eb4"
	apiUrl := "http://api.openweathermap.org/data/2.5/weather?q="+temp.conf.Url.Location+"&appid=" + temp.conf.Url.Key
    
    // Fetch weather data from the API
    resp, err := http.Get(apiUrl)
    if resp.StatusCode != 200 {
        fmt.Println("Error fetching weather data:", err)
        return 0, errors.New("error fetching weather data")
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return 0,err
    }

	var weatherResponse WeatherRepository
    // Unmarshal JSON response
    if err := json.Unmarshal(body, &weatherResponse); err != nil {
        fmt.Println("Error unmarshalling JSON:", err)
        return 0,err
    }

    // Print the temperature
    return weatherResponse.Days.Temp, nil
}

func (w WeatherRepository) CreateTemperatureDatabase(uid uint, temp float64, t time.Time){
    dsn := w.conf.Database.User+":"+w.conf.Database.Password+"@tcp("+w.conf.Database.Host+":"+w.conf.Database.Port+")/"+w.conf.Database.DBname+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ Logger: &SqlLoggerInterface{}})
	if err!= nil {
        panic(err.Error())
    }
    data := TemperatureEntity{Uid:uid, Temperature: temp, Times: t}
	result := db.Create(&data) // pass pointer of data to Create
	if result.Error!= nil {
        panic(result.Error)
    }
	fmt.Println(data.ID)
	fmt.Println(result.RowsAffected)
}

func (t TemperatureEntity) TableName() string {
    return "time_temperature"
  }
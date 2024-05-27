package temperature

import (
	"encoding/json"
	"errors"
	"go-calculator-tutorial/internal/config"
	"io/ioutil"
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type IWeatherRepository interface {
	GetTemperature() (float64, error);
    GetDBLocation() string;
    GetRedisClient() *redis.Client;
}


type WeatherRepository struct {
	conf   *config.AppConfig
    client *redis.Client
	Days struct {
        Temp float64 `json:"temp"`
    } `json:"main"`
}


func NewWeatherRepository(conf *config.AppConfig,c *redis.Client) IWeatherRepository {
	return &WeatherRepository{
		conf: conf,
        client: c,
	}
}

func (temp WeatherRepository) GetTemperature() (float64, error) {

	// const apiKey = "913da61969a3190ba9de6b5c094a0eb4"
	apiUrl := "http://api.openweathermap.org/data/2.5/weather?q="+temp.conf.Url.Location+"&appid=" + temp.conf.Url.Key
    
    // Fetch weather data from the API
    resp, err := http.Get(apiUrl)
    if resp.StatusCode != 200 {
        slog.Error("Error fetching weather data:", err)
        return 0, errors.New("error fetching weather data")
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        slog.Error("Error reading response body:", err)
        return 0,err
    }

	var weatherResponse WeatherRepository
    // Unmarshal JSON response
    if err := json.Unmarshal(body, &weatherResponse); err != nil {
        slog.Error("Error unmarshalling JSON:", err)
        return 0,err
    }

    // Print the temperature
    return weatherResponse.Days.Temp, nil
}

// Get database connection link
func (w WeatherRepository) GetDBLocation() string {
    db := w.conf.Database.User+":"+w.conf.Database.Password+"@tcp("+w.conf.Database.Host+":"+w.conf.Database.Port+")/"+w.conf.Database.DBname+"?charset=utf8mb4&parseTime=True&loc=Local"
    return db
}

  // Get redis client
  func (w WeatherRepository) GetRedisClient() *redis.Client {
    return w.client
  }
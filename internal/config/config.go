package config

type AppConfig struct {
	Url Url
	Server Server
	Service Service
	Database Database
}

type Url struct {
    Key string
	Location string
}

type Server struct {
	Address  string
}

type Service struct {
	WeatherConfig  string
}

type Database struct {
	Driver string
	Host string
	Port string
	User string
	Password string
	DBname string
}
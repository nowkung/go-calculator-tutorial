package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convertCelsius(t *testing.T) {
	temperature := 273.15

	ans := convertCelsius(temperature)

	assert.Equal(t, "0.00", ans)
}

func Test_convertCelsiusPositive(t *testing.T) {
	temperature := 373.15

	ans := convertCelsius(temperature)

	assert.Equal(t, "100.00", ans)
}

func Test_convertCelsiusNegative(t *testing.T) {
	temperature := 173.15

	ans := convertCelsius(temperature)

	assert.Equal(t, "-100.00", ans)
}

func Test_convertFahrenheit(t *testing.T) {
	temperature := 459.67

	ans := convertFahrenheit(temperature)

	assert.Equal(t, "0.00", ans)
}

func Test_convertFahrenheitPositive(t *testing.T) {
	temperature := 559.67

	ans := convertFahrenheit(temperature)

	assert.Equal(t, "100.00", ans)
}

func Test_convertFahrenheitNegative(t *testing.T) {
	temperature := 359.67

	ans := convertFahrenheit(temperature)

	assert.Equal(t, "-100.00", ans)
}

func Test_convertKelvin(t *testing.T) {
	temperature := 0.00

	ans := convertKelvin(temperature)

	assert.Equal(t, "0.00", ans)
}

func Test_convertKelvinPositive(t *testing.T) {
	temperature := 100.00

	ans := convertKelvin(temperature)

	assert.Equal(t, "100.00", ans)
}

func Test_convertKelvinNegative(t *testing.T) {
	temperature := -100.00

	ans := convertKelvin(temperature)

	assert.Equal(t, "-100.00", ans)
}

// func TestWeatherService_TemperatureWithUnit(t *testing.T) {
// 	type fields struct {
// 		weatherRepo temperature.IWeatherRepository
// 	}
// 	type args struct {
// 		unit string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    map[string]string
// 		wantErr bool
// 	}{
// 		{
// 			name: "Celsius",
//             fields: fields{
//                 weatherRepo: &temperature.WeatherRepository{},
//             },
//             args: args{
//                 unit: "C",
//             },
//             want: map[string]string{
//                 "Temperature": "0.00C",
//             },
//             wantErr: false,
// 		},
// 		{
// 			name: "Fahrenheit",
//             fields: fields{
//                 weatherRepo: &temperature.WeatherRepository{},
//             },
//             args: args{
//                 unit: "F",
//             },
//             want: map[string]string{
//                 "Temperature": "0.00F",
//             },
//             wantErr: false,
// 		},
// 		{
// 			name: "Kelvin",
//             fields: fields{
//                 weatherRepo: &temperature.WeatherRepository{},
//             },
//             args: args{
//                 unit: "K",
//             },
//             want: map[string]string{
//                 "Temperature": "0.00K",
//             },
//             wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			w := w
// 			got, err := w.TemperatureWithUnit(tt.args.unit)
// 			if (err != nil) != tt.wantErr {
// 				assert.Error(t,err)
// 			}
// 			assert.Equal(t, tt.want, got)
// 		})
// 	}
// }

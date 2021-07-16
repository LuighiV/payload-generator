// Package openweather provides interface to openweather
// Based on https://gist.github.com/aodin/9493190
// https://medium.com/@xcoulon/nested-structs-in-golang-2c750403a007
package openweather

import (
	"encoding/json"
	"fmt"
	"github.com/LuighiV/payload-generator/generator/converter"
	"log"
	"net/http"
)

// Data holds the information of the parameters received from the API of
// OpenWeather and the corresponding converter value of the
// payload in bytes
type Data struct {
	latitude    float64
	longitude   float64
	temperature int
	humidity    int
	pressure    int
	wind_speed  int
	payload     []byte
}

// OWConfig has the APIKey to connect to the OW service and the city wher those
// data come from
type OWConfig struct {
	apikey string
	city   string
}

// Message holds the data obtained from the OpenWeather API
// Example response
//{
//    "coord": {
//        "lon": -78.65,
//        "lat": -6.55
//    },
//    "weather": [
//        {
//            "id": 500,
//            "main": "Rain",
//            "description": "light rain",
//            "icon": "10n"
//        }
//    ],
//    "base": "stations",
//    "main": {
//        "temp": 286.23,
//        "feels_like": 286.15,
//        "temp_min": 286.23,
//        "temp_max": 286.23,
//        "pressure": 1017,
//        "humidity": 98,
//        "sea_level": 1017,
//        "grnd_level": 764
//    },
//    "visibility": 6513,
//    "wind": {
//        "speed": 1.73,
//        "deg": 34,
//        "gust": 2.37
//    },
//    "rain": {
//        "1h": 0.41
//    },
//    "clouds": {
//        "all": 99
//    },
//    "dt": 1617508084,
//    "sys": {
//        "country": "PE",
//        "sunrise": 1617448606,
//        "sunset": 1617491906
//    },
//    "timezone": -18000,
//    "id": 3698141,
//    "name": "Chota",
//    "cod": 200
//}
type Message struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string
	Main struct {
		Temp       float64 `json:"temp"`
		Feels_like float64 `json:"feels_like"`
		Temp_min   float64 `json:"temp_min"`
		Temp_max   float64 `json:"temp_max"`
		Pressure   float64 `json:"pressure"`
		Humidity   float64 `json:"humidity"`
		Sea_level  float64 `json:"sea_level"`
		Grnd_level float64 `json:"grnd_level"`
	} `json:"main"`
	Visibility int
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Dt  int64 `json:"dt"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

// DataOption is the kind of option to be applied to the Data structure
type DataOption func(*Data) error

// NewOWConfig creates a new configuration to connect to the OW service
func NewOWConfig(apikey string, city string) (*OWConfig, error) {

	owconf := OWConfig{
		apikey: apikey,
		city:   city,
	}
	return &owconf, nil
}

// GetOpenDataByCityName obtains data taking the option as input to get
// information
func GetOpenDataByCityName(owconf *OWConfig) DataOption {
	return func(d *Data) error {

		r, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=" + owconf.city + "&appid=" + owconf.apikey)
		if err != nil {
			log.Fatal(err)
			return err
		}

		var msg Message
		err = json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			log.Fatal(err)
			return err
		}

		d.temperature = int((msg.Main.Temp - 273) * 100)
		d.humidity = int((msg.Main.Humidity) * 100)
		d.pressure = int((msg.Main.Pressure) * 100)
		d.pressure = int((msg.Main.Pressure) * 100)
		d.latitude = msg.Coord.Lat
		d.longitude = msg.Coord.Lon
		d.wind_speed = int(msg.Wind.Speed * 100)

		return nil
	}

}

// GetPayload obtains the value of payload in bytes
func GetPayload(d *Data) []byte {
	return d.payload
}

// LoadPayload generates the payload based on the data obtained from OW
func LoadPayload() DataOption {
	return func(d *Data) error {
		bs := make([]byte, 24)
		bs = converter.ConvertIntToBytes(d.temperature)
		bs = append(bs, converter.ConvertIntToBytes(d.humidity)...)
		bs = append(bs, converter.ConvertIntToBytes(d.pressure)...)
		bs = append(bs, converter.ConvertIntToBytes(d.wind_speed)...)
		bs = append(bs, converter.ConvertFloatToBytes(float32(d.latitude))...)
		bs = append(bs, converter.ConvertFloatToBytes(float32(d.longitude))...)
		d.payload = bs
		return nil
	}
}

// NewData is a function to create the data structure based on the options
// applied
func NewData(opts ...DataOption) (*Data, error) {

	d := &Data{}

	for _, o := range opts {
		if err := o(d); err != nil {
			return nil, err
		}
	}

	fmt.Println(d.temperature)
	fmt.Println(d.humidity)
	fmt.Println(d.pressure)
	fmt.Println(d.wind_speed)
	fmt.Println(d.latitude)
	fmt.Println(d.longitude)
	LoadPayload()(d)
	return d, nil
}

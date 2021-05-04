// Package random provides functions to generate random values for
// metheorological data
package random

import (
	"fmt"
	"github.com/LuighiV/payload-generator/generator/converter"
	"math/rand"
	"time"
)

// GenerateRandom returns a random value receiving the a base value and
// variation value which determines the range of variation.
func GenerateRandom(basevalue float64, rangevariation float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return basevalue + rand.Float64()*rangevariation - rangevariation/2
}

// A Base_Values to introduce base values and variation values for each
// parameter (temperature, pressure and humidity)
type Base_Values struct {
	temperature_base      float64
	humidity_base         float64
	pressure_base         float64
	temperature_variation float64
	humidity_variation    float64
	pressure_variation    float64
}

func NewBaseValues(
	temperature_base float64,
	humidity_base float64,
	pressure_base float64,
	temperature_variation float64,
	humidity_variation float64,
	pressure_variation float64,
) (*Base_Values, error) {

	base_values := Base_Values{
		temperature_base,
		humidity_base,
		pressure_base,
		temperature_variation,
		humidity_variation,
		pressure_variation,
	}
	return &base_values, nil
}

// Data holds the information of the parameters converted to integers and the
// payload in bytes
type Data struct {
	temperature int
	humidity    int
	pressure    int
	payload     []byte
}

// DataOption is the kind of option to be applied to the Data structure
type DataOption func(*Data) error

// WithTemperature returns a temperature value with a base and variation values
func WithTemperature(base float64, variation float64) DataOption {
	return func(d *Data) error {
		d.temperature = int(GenerateRandom(base, variation) * 100)
		return nil
	}
}

// WithHumidity returns a humidity value with a base and variation values
func WithHumidity(base float64, variation float64) DataOption {
	return func(d *Data) error {
		d.humidity = int(GenerateRandom(base, variation) * 100)
		return nil
	}
}

// WithPressure returns a pressure value with a base and variation values
func WithPressure(base float64, variation float64) DataOption {
	return func(d *Data) error {
		d.pressure = int(GenerateRandom(base, variation) * 100)
		return nil
	}
}

func WithBaseValues(base *Base_Values) DataOption {
	return func(d *Data) error {

		WithTemperature(base.temperature_base, base.temperature_variation)(d)
		WithHumidity(base.humidity_base, base.humidity_variation)(d)
		WithPressure(base.pressure_base, base.pressure_variation)(d)
		return nil
	}
}

// GetPayload returns the payload in bytes
func GetPayload(d *Data) []byte {
	return d.payload
}

// LoadPayload generates the payload based on the temperature,
// humidity and pressure values
func LoadPayload() DataOption {
	return func(d *Data) error {
		bs := make([]byte, 12)
		bs = converter.ConvertIntToBytes(d.temperature)
		bs = append(bs, converter.ConvertIntToBytes(d.humidity)...)
		bs = append(bs, converter.ConvertIntToBytes(d.pressure)...)
		d.payload = bs
		return nil
	}
}

// NewData is a function to create the data structure and apply options to
// generate the parameters of temperature, pressure and humidity
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
	LoadPayload()(d)
	return d, nil
}

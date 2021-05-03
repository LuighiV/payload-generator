// Package random provides functions to generate random values for
// metheorological data
package random

import (
	"fmt"
	"github.com/LuighiV/payload-generator/generator/converter"
	"math/rand"
	"time"
)

func GenerateRandom(basevalue float64, rangevariation float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return basevalue + rand.Float64()*rangevariation - rangevariation/2
}

type Base_Values struct {
	temperature_base      float64
	humidity_base         float64
	pressure_base         float64
	temperature_variation float64
	humidity_variation    float64
	pressure_variation    float64
}

type Data struct {
	temperature int
	humidity    int
	pressure    int
	payload     []byte
}

type DataOption func(*Data) error

func WithTemperature(base float64, variation float64) DataOption {
	return func(d *Data) error {
		d.temperature = int(GenerateRandom(base, variation) * 100)
		return nil
	}
}

func WithHumidity(base float64, variation float64) DataOption {
	return func(d *Data) error {
		d.humidity = int(GenerateRandom(base, variation) * 100)
		return nil
	}
}

func WithPressure(base float64, variation float64) DataOption {
	return func(d *Data) error {
		d.pressure = int(GenerateRandom(base, variation) * 100)
		return nil
	}
}

func GetPayload(d *Data) []byte {
	return d.payload
}

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

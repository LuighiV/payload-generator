// Package generator provides an interface to generate values
package generator

import (
	"github.com/LuighiV/payload-generator/generator/openweather"
	"github.com/LuighiV/payload-generator/generator/random"
)

type Base_Values struct {
	temperature_base      float64
	humidity_base         float64
	pressure_base         float64
	temperature_variation float64
	humidity_variation    float64
	pressure_variation    float64
}
type Generator struct {
	random_base Base_Values
	ow_config   openweather.OWConfig
	payload     []byte
}

type GeneratorType int
type GeneratorOption func(*Generator) error

const (
	Random      GeneratorType = iota
	OpenWheater GeneratorType = iota
)

func NewGenerator(opts ...GeneratorOption) (*Generator, error) {
	gen := &Generator{}

	for _, o := range opts {
		if err := o(gen); err != nil {
			return nil, err
		}
	}

	return gen, nil
}

func WithRandomBase(
	temperature_base float64,
	humidity_base float64,
	pressure_base float64,
	temperature_variation float64,
	humidity_variation float64,
	pressure_variation float64,
) GeneratorOption {
	return func(gen *Generator) error {

		gen.random_base = Base_Values{
			temperature_base,
			humidity_base,
			pressure_base,
			temperature_variation,
			humidity_variation,
			pressure_variation,
		}

		return nil

	}

}

func WithOWeatherConfig(apikey string, city string) GeneratorOption {
	return func(gen *Generator) error {
		owconf, err := openweather.NewOWConfig(apikey, city)
		if err != nil {
			panic(err)
		}
		gen.ow_config = *owconf
		return nil
	}

}

func Generate(t GeneratorType) GeneratorOption {
	return func(gen *Generator) error {

		if t == Random {
			d, err := random.NewData(
				random.WithTemperature(gen.random_base.temperature_base, gen.random_base.temperature_variation),
				random.WithHumidity(gen.random_base.humidity_base, gen.random_base.humidity_variation),
				random.WithPressure(gen.random_base.pressure_base, gen.random_base.pressure_variation),
			)

			if err != nil {
				panic(err)
			}
			gen.payload = random.GetPayload(d)
		} else if t == OpenWheater {
			d, err := openweather.NewData(
				openweather.GetOpenDataByCityName(&gen.ow_config),
			)

			if err != nil {
				panic(err)
			}
			gen.payload = openweather.GetPayload(d)
		}
		return nil
	}
}

func GetPayload(gen *Generator) []byte {
	return gen.payload
}

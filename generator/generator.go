// Package generator provides an interface to generate values
package generator

import (
	"github.com/LuighiV/payload-generator/generator/openweather"
	"github.com/LuighiV/payload-generator/generator/random"
)

// Generator holds the parameters required to generate Data and the paylod
// generated
type Generator struct {
	random_base random.Base_Values
	ow_config   openweather.OWConfig
	payload     []byte
}

// GeneratorType is an alias for the enumeration value of the generator
type GeneratorType int

// GeneratorOption holds the function option for generator
type GeneratorOption func(*Generator) error

const (
	Random      GeneratorType = iota // Random generator type
	OpenWeather GeneratorType = iota // OpenWeather generator type
)

// NewGenerator creates a new generator structure
func NewGenerator(opts ...GeneratorOption) (*Generator, error) {
	gen := &Generator{}

	for _, o := range opts {
		if err := o(gen); err != nil {
			return nil, err
		}
	}

	return gen, nil
}

// WithRandomBase creates the structure for random values
func WithRandomBase(
	temperature_base float64,
	humidity_base float64,
	pressure_base float64,
	temperature_variation float64,
	humidity_variation float64,
	pressure_variation float64,
) GeneratorOption {
	return func(gen *Generator) error {

		basevalues, err := random.NewBaseValues(
			temperature_base,
			humidity_base,
			pressure_base,
			temperature_variation,
			humidity_variation,
			pressure_variation,
		)
		if err != nil {
			panic(err)
		}

		gen.random_base = *basevalues
		return nil

	}

}

// WithOWeatherConfig creates the structure with OW options
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

// Generate the payload depending on the type of generator
func Generate(t GeneratorType) GeneratorOption {
	return func(gen *Generator) error {

		if t == Random {
			d, err := random.NewData(
				random.WithBaseValues(&gen.random_base),
			)

			if err != nil {
				panic(err)
			}
			gen.payload = random.GetPayload(d)
		} else if t == OpenWeather {
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

// GetPayload returns the payload in bytes array format
func GetPayload(gen *Generator) []byte {
	return gen.payload
}

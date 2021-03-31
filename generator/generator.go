// Package generator provides an interface to generate values
package generator

import (
	"github.com/LuighiV/payload-generator/generator/random"
)

type Generator struct {
	temperature_base      float64
	humidity_base         float64
	pressure_base         float64
	temperature_variation float64
	humidity_variation    float64
	pressure_variation    float64
	payload               []byte
}

type GeneratorType int
type GeneratorData func(*Generator) error

const (
	Random      GeneratorType = iota
	OpenWheater GeneratorType = iota
)

func NewGenerator(
	temperature_base float64,
	humidity_base float64,
	pressure_base float64,
	temperature_variation float64,
	humidity_variation float64,
	pressure_variation float64,
) (*Generator, error) {
	gen := &Generator{
		temperature_base,
		humidity_base,
		pressure_base,
		temperature_variation,
		humidity_variation,
		pressure_variation,
		[]byte{},
	}

	return gen, nil
}

func Generate(t GeneratorType) GeneratorData {
	return func(gen *Generator) error {

		if t == Random {
			d, err := random.NewData(
				random.WithTemperature(gen.temperature_base, gen.temperature_variation),
				random.WithHumidity(gen.humidity_base, gen.humidity_variation),
				random.WithPressure(gen.pressure_base, gen.pressure_variation),
			)

			if err != nil {
				panic(err)
			}
			gen.payload = random.GetPayload(d)
		}
		return nil
	}
}

func GetPayload(gen *Generator) []byte {
	return gen.payload
}

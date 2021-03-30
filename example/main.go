package main

import (
	"fmt"
	"github.com/LuighiV/payload-generator/generator/random"
)

func main() {
	fmt.Println("Hello, World!")
	l := random.GenerateRandom(40, 5)

	fmt.Println(l)

	d, err := random.NewData(
		random.WithTemperature(39.0, 0.5),
		random.WithHumidity(80.0, 5),
		random.WithPressure(1000.0, 100),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(random.GetPayload(d))
}

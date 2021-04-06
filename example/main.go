package main

import (
	"fmt"
	"github.com/LuighiV/payload-generator/generator"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	gen, err := generator.NewGenerator(
		generator.WithRandomBase(
			39.0,
			80.0,
			1000.0,
			0.5,
			5,
			100,
		),
		generator.WithOWeatherConfig(
			os.Getenv("OW_API_KEY"), "London",
		),
	)
	if err != nil {
		panic(err)
	}
	generator.Generate(generator.Random)(gen)
	fmt.Println(generator.GetPayload(gen))

	generator.Generate(generator.OpenWheater)(gen)
	fmt.Println(generator.GetPayload(gen))

}

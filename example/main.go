package main

import (
	"fmt"
	"github.com/LuighiV/payload-generator/generator"
)

func main() {
	fmt.Println("Hello, World!")

	gen, err := generator.NewGenerator(
		39.0,
		80.0,
		1000.0,
		0.5,
		5,
		100,
	)
	if err != nil {
		panic(err)
	}
	generator.Generate(generator.Random)(gen)
	fmt.Println(generator.GetPayload(gen))
}

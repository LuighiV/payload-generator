package main

import (
	"fmt"
	"luighiv/generator"
)

func main() {
	fmt.Println("Hello, World!")
	l := generateRandom(40, 5)
	fmt.Println(l)
}

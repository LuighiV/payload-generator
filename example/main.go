package main

import (
	"fmt"
	"github.com/LuighiV/payload-generator"
)

func main() {
	fmt.Println("Hello, World!")
	l := generateRandom(40, 5)
	fmt.Println(l)
}

package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	Matematica()
	fmt.Println("------------------------")
	Cadenas()
	fmt.Println("------------------------")
}

func Matematica() {

	fmt.Println(math.Floor(36.7688))
	fmt.Println(math.Signbit(-0.1))
	fmt.Println(math.Signbit(12))
}

func Cadenas() {
	fmt.Println(strings.ContainsAny("La cadena contiene el caracter", "i"))
	fmt.Println(strings.Trim("---Esto es una cadena suprimiendo los caracteres de inicio y fin----", "-"))

}

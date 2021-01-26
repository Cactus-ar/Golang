package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var intento int

	source := rand.NewSource(time.Now().UnixNano())
	aleatorio := rand.New(source)
	numeroSecreto := aleatorio.Intn(10) //numero generado entre 0 y 10

	fmt.Println("-------------------------------")
	fmt.Println("	Adivina el Numero!!")
	fmt.Println("-------------------------------")
	fmt.Print("Qué Número Estoy Pensando?: ")
	fmt.Scan(&intento)
	fmt.Println("-------------------------------")

	if intento == numeroSecreto {
		fmt.Println("Sep, ese es ")
	} else if intento > numeroSecreto {
		fmt.Println("Menos")
	} else if intento < numeroSecreto {
		fmt.Println("Más arriba")
	}

}

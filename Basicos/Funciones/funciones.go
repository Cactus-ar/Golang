package main

import "fmt"

func main() {

	var mensaje string = "Salida por pantalla"
	Mensaje(mensaje)
	Suma(4, 6)
	fmt.Println(Division(124, 7))

}

// Las funciones se declaran de forma similar a otros lenguajes.
// si la funcion no retorna nada (void), solo se agregan los parametros de entrada
func Mensaje(mensaje string) {
	fmt.Println(mensaje)
}

func Suma(numero1, numero2 int) {
	var resultado int = numero1 + numero2
	fmt.Println(resultado)
}

//Si la funcion retorna algún tipo se declaran luego de los argumentos de entrada
func Division(numero1, numero2 int) int {
	return numero1 / numero2

}

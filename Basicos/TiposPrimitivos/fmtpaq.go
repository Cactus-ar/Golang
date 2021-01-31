package main

import (
	"fmt"
)

//referencia hacia el paquete fmt https://golang.org/pkg/fmt/

var cierovalor = 262

func main() {

	fmt.Println(cierovalor)        //valor asignado
	fmt.Printf("%T\n", cierovalor) //tipo
	fmt.Printf("%b\n", cierovalor) //binario
	fmt.Printf("%x\n", cierovalor) //etc
	fmt.Printf("%c\n", cierovalor) //..
	fmt.Printf("%U\n", cierovalor) //

	//Sprint devuelve la evaluacion de la expresion en cadena
	cadena := fmt.Sprintf("%b\n", cierovalor)
	fmt.Printf("%T\n", cadena)
	fmt.Println("Devuelto como string", cadena)

}

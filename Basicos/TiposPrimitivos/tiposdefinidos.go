package main

import "fmt"

type dado int //un int definido tipo dado

var numero int = 45
var dado_1 dado = 2 //dado_1 (tipo dado) se le asigna un valor de 2

func main() {

	fmt.Println(numero)
	fmt.Printf("%T\n", numero) //tipo int regular
	fmt.Println(dado_1)
	fmt.Printf("%T\n", dado_1) //tipo int regular

}

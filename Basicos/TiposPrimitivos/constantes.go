package main

import (
	"fmt"
)

//declaracion de constantes
//puede hacerse en forma separada o agruparlas
const separada = "constante Separada"

//constante con tipo declarado
const contipo float32 = 19.6363

// el compilador interpreta el tipo
const (
	constante1 = 12
	constante2 = 16.88
	constante3 = "Hola"
	constante4 = 0x11010101b
)

// puede operarse sobre las constantes
const (
	kb = 1024
	mb = 1024 * kb
	gb = 1024 * mb
)

func main() {

	fmt.Println(contipo)
	fmt.Println(constante1)
	fmt.Println(constante2)
	fmt.Println(constante3)
	fmt.Println(constante4)
	fmt.Printf("%T\n", constante1)
	fmt.Printf("%T\n", constante2)
	fmt.Printf("%T\n", constante3)
	fmt.Printf("%T\n", constante4)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}

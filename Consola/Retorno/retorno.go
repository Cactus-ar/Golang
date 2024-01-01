package main

import "fmt"

//retorno simple de funcion
func nombreCompleto(nombre string, apellido string) string {
	return nombre + " " + apellido
}

func main() {
	fmt.Println(nombreCompleto("Roberto", "Perez"))
}

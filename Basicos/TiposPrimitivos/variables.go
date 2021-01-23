package main

import "fmt"

//-- Tipos de Variables

//-- Scope
var variableConAccesoGlobal int = 0

func main() {

	fmt.Println("---------------------------------------")
	Texto()
	fmt.Println("---------------------------------------")
	NumerosInt()
	fmt.Println("---------------------------------------")
	Tipos()
	fmt.Println("---------------------------------------")
	Logicas()
	fmt.Println("---------------------------------------")
	DeclaracionesCortas()
	fmt.Println("---------------------------------------")
	Fin()
	fmt.Println("---------------------------------------")
}

func Texto() {

	var variable_Texto string = "Cadena De Texto" //Podia omitirse el tipo ya que el compilador puede inferir que es del tipo string
	var variable_Texto2 = "Otra Cadena de Texto"
	fmt.Println(variable_Texto)
	fmt.Println(variable_Texto2)
	variableConAccesoGlobal++

}

func NumerosInt() {

	var unEnteroSimple int = 26 //el complilador arrojará un error si la variable no es usada
	fmt.Println(unEnteroSimple)
	variableConAccesoGlobal++

}

func Tipos() { //el paquete implementa entrada y salida formateada, referencia de verbos: https://golang.org/pkg/fmt/
	var doble = 1234.6369
	fmt.Printf("%T\n", doble) //dada la variable en este caso devuelve su tipo
	variableConAccesoGlobal++
}

func Logicas() {

	var logica bool
	logica = true
	fmt.Println(logica)
	variableConAccesoGlobal++
}

func DeclaracionesCortas() { //no pueden utilizarse estos atajos fuera de funciones

	nombre, apellido := "Gabriel", "Ceravolo"
	altura := 1.78

	fmt.Println(nombre, apellido, altura)

}

func Fin() {
	fmt.Println(variableConAccesoGlobal)
}

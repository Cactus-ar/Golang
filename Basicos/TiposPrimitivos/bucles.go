package main

import "fmt"

//bucles en Go

//No existe el while

func main() {

	cuenta := 10

	//la estructura basica es:
	//for inicio; condicion; incremento o condicion
	for i := 0; i < cuenta; i++ {
		fmt.Println(i)
	}

	//puede evaluar condicionales como el while-loop
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	cuenta2 := 0
	for {
		fmt.Println(cuenta2)
		cuenta2++
		if cuenta2 == 3 {
			break
		}
	}

}

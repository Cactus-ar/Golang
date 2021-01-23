// Varias cosas para comentar del ejemplo que salen del libro:

/*

Las herramientas del lenguaje se acceden por la linea de comandos invocando siempre a Go como inicial precedido por sub comandos (Ejemplo: go build xxx.go).
Al instalarse se requieren algunas variables de entorno para utilizar alguno de los comandos, por ejemplo el comando get.

go get ruta -> va a descargar el repo en la carpeta "src" predefinida en el path.
(en el caso de la instalacion de windows que es automática la carpeta de trabajo se encuentra
dentro del directorio del usuario, "C:\Users\usuario\go")

*/

package main

//Go se encuentra organizado en paquetes (packages) que es similar a librerías módulos de otros lenguajes
//en este caso la declaracion de main es especial ya que define un ejecutable solamente (standalone).

import "fmt" //la libreria standard de Go consta de mas de 100 paquetes, en este caso ftm es el paquete que contiene librerías para salida por pantalla y entrada
//un programa no va a compilar si faltan referencias PERO TAMPOCO SI EXISTEN SIN USAR

//main declarada como funcion también es especial ya que es el punto de entrada del programa. la llave de apertura DEBE estar al final de la declaración. no puede estar separada (como por ejemplo en c++) y producirá un error de sintaxis si se omite la regla
func main() {

	fmt.Println("Hola desde Go, مرحبا ديدي اذهب, 打个招呼") //De forma nativa Go maneja Unicode
}

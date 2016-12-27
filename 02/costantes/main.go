// Comentario de línea

/*
Comentario de bloque
*/

package main

import "fmt"

func main() {

	// const nombre = "Pedro"
	// fmt.Println(nombre)

	var a int
	var b int8
	n := "Pedro"
	p := "Gómez"

	a = 121212
	b = 5

	// casting
	c := a + int(b)
	fmt.Println(c)
	// mostrar el tipo de datos, %s string, %d número.
	fmt.Printf("hola %s %s %d cómo estás? \n", n, p, b)
	fmt.Printf("c es de tipo %T y b es de tipo %T \n", c, b) // %T nos dice que tipo de datos es.
	// prioridad aritmética
	// + - * /
	// () % * / * -

	// d := 6 + 6 * 6 - 6
	//	d := 6 + 6*(6-6)

	var nombre string
	var numero int
	var entiendes bool

	fmt.Println(nombre, numero, entiendes)
}

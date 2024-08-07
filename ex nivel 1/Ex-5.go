package main

import "fmt"

type coffee int

var x coffee

var y int

func main() {
	fmt.Print("Valor de x: ")
	fmt.Println(x)
	fmt.Print("Tipo de x: ")
	fmt.Printf("%T\n", x)
	fmt.Print("Valor de y: ")
	fmt.Println(y)

	x = 42
	y = int(x)
	fmt.Print("Valor de y: ")
	fmt.Println(y)
	fmt.Print("Tipo de y: ")
	fmt.Printf("%T\n", y)
}

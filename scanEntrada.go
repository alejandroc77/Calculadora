package main

import (
	"fmt"
)

func menu() {
	var eleccion int

	fmt.Println("Bienvenido!!")
	fmt.Println("Eliga que operacion desea hacer \n" +
		"1: Suma \n" +
		"2: Multiplicaion\n" +
		"3: Division \n" +
		"4: Resta \n" +
		"5: Salir")

	fmt.Scan(&eleccion)

	switch eleccion {
	case 1:
		fmt.Println("Genial!, Haz escogido suma")
		fmt.Println("El resultado de la suma es => ", suma())
	case 2:
		fmt.Println("Genial!, Su eleccion es Multiplicacion")
		fmt.Println("El resultado de la Multiplicacion es => ", multiplicaion())
	case 3:
		fmt.Println("Genial!, Su eleccion es Division")
		fmt.Println("El resultado de la Division es => ", division())
	case 4:
		fmt.Println("Genial!, Su eleccion es Resta")
		fmt.Println("El resultado de la Resta es => ", resta())
	case 5:
		fmt.Println("Exit")
	default:
		menu()
	}
}

func suma() int {
	var num1, num2 int
	fmt.Println("Ingrese primer numero")
	fmt.Scan(&num1)
	fmt.Println("Ingrese segundo numero")
	fmt.Scan(&num2)
	result := num1 + num2
	return result

}
func multiplicaion() int {
	var num1, num2 int
	fmt.Println("Ingrese primer numero")
	fmt.Scan(&num1)
	fmt.Println("Ingrese segundo numero")
	fmt.Scan(&num2)
	result := num1 * num2
	return result
}
func division() int {
	var num1, num2 int
	fmt.Println("Ingrese primer numero")
	fmt.Scan(&num1)
	fmt.Println("Ingrese segundo numero")
	fmt.Scan(&num2)
	result := num1 / num2
	return result
}
func resta() int {
	var num1, num2 int
	fmt.Println("Ingrese primer numero")
	fmt.Scan(&num1)
	fmt.Println("Ingrese segundo numero")
	fmt.Scan(&num2)
	result := num1 - num2
	return result
}

func main() {
	menu()
}

package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)
	fmt.Println("Capacidad: ", cap(slice))
	fmt.Println("Tamaño: ", len(slice))
	fmt.Println("Arreglo: ", slice)
	slice = append(slice, 2)
	slice = append(slice, 4)
	fmt.Println("Capacidad: ", cap(slice))
	fmt.Println("Tamaño: ", len(slice))
	fmt.Println("Arreglo: ", slice)
	slice = append(slice, 6)
	slice = append(slice, 8)
	fmt.Println("Capacidad: ", cap(slice))
	fmt.Println("Tamaño: ", len(slice))
	fmt.Println("Arreglo: ", slice)
}

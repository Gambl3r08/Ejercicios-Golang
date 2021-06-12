package slice

import "fmt"

func main() {
	arreglo := [4]int{1, 2, 3, 4}
	slice := arreglo[2:4]

	fmt.Println(slice)
}

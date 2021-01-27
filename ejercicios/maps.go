package main

import "fmt"

func main() {
	var m map[string]float64
	m = map[string]float64{"id": 5222.2, "id2": 2}
	valor1 := m["id"]
	fmt.Println(m)
	fmt.Println(valor1)

	//Comprobar

	v, err := m["algunacosa"]
	if err == false {
		println(err)
	} else {
		println(v)
	}

	// for
	for key, value := range m {
		fmt.Println(key)
		fmt.Println(value)
	}

}

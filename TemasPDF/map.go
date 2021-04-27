package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["v1"] = 1
	m["v2"] = 2

	fmt.Println("map: ", m)
}

package main

import "fmt"

func main() {
	n := 20
	pn := new(int)
	pn = &n
	fmt.Println(pn)
}

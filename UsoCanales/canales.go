package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	go printMess("Hola 1", c)
	go printMess("Hola 2", c)
	go printMess("Hola 3", c)

	mss := <-c
	fmt.Print(mss)

	close(c)
}

func printMess(s string, c chan string) {
	fmt.Println(s)
	c <- s
}

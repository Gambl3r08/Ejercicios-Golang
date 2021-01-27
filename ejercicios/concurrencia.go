package main

import (
	"fmt"
)

func IsReady(what string) {
	fmt.Println(what, "está listo")
}

func main() {

	go IsReady("Té")
	go IsReady("Cafe")
	fmt.Println("ultima linea del main")

}

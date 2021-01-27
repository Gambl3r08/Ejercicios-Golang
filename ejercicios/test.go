package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	leer("test.go")

}

func leer(name string) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		//fmt.Println(scanner.Bytes())
	}
}

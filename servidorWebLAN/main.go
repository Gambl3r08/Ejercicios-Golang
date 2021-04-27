package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Home page")
	})

	port := ":8080"
	fmt.Println("Escuchando en el puerto: " + port)
	log.Fatal(http.ListenAndServeTLS(port, "server.crt", "server.key", nil))

}

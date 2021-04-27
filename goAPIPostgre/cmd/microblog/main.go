package main

import (
	"apiREST/internal/data"
	"apiREST/internal/server"
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	go serv.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	serv.Close()
	data.Close()
}

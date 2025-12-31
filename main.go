package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/sambakker4/calculator-api/cmd"
)

const port = "8080"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/add", cmd.Add)
	mux.HandleFunc("/subtract", cmd.Subtract)
	mux.HandleFunc("/divide", cmd.Divide)
	mux.HandleFunc("/multiply", cmd.Multiply)

	server := &http.Server{
		Handler: mux,
		Addr: ":" + port,
	}

	fmt.Println("Calculator API listening on port:", port)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

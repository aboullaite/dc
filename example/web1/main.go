package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/", ExampleHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	for {
		time.Sleep(2 * time.Second)
		go log.Println("** Service Started on Port " + port + " **")

	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"toto"}`)

}

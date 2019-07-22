package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	startHTTPServer()
}

func startHTTPServer() {
	var PORT string
	var HOST string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "80"
	}

	if HOST = os.Getenv("HOST"); HOST == "" {
		HOST = ""
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	log.Printf("Starting Server %s", fmt.Sprintf("%s:%s", HOST, PORT))

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", HOST, PORT), nil)
	if err != nil {
		panic(err)
	}

}

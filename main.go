package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var requests int

func main() {
	fmt.Println("Counter Service starting")

	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":"+os.Getenv("LISTEN_PORT"), nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("error starting server: %s", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	requests += 1
	io.WriteString(w, fmt.Sprint(requests))
}

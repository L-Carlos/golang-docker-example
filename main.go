package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloDocker)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloDocker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Docker")
}

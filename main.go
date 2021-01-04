package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.FPrintf("Hello world!")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
  
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":" + port, nil))
}

package main

import (
	"lastgram.xyz/ethereal/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.Handler)
	print("Done.\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

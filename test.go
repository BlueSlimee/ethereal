package main

import(
	"net/http"
	"lastgram.xyz/ethereal/api"
	"log"
)

func main() {
	http.HandleFunc("/", api.Handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	print("Ready\n")
}

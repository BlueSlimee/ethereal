package main

import(
	"net/http"
	"lastgram.xyz/ethereal/handler"
	"log"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	print("Ready\n")
}

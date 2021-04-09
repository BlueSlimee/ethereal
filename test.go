package main

import (
	"lastgram.xyz/ethereal/api"
	"log"
	"net/http"
	"lastgram.xyz/ethereal/utils"
)

func main() {
	print(utils.GetPath() +"\n")
	http.HandleFunc("/", api.Handler)
	print("Done.\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

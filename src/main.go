package main

import (
	"net/http"
)

func main() {
	PORT := setPort() //Read port from config.yml
	if err := http.ListenAndServe(PORT, nil); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	PORT := setPort() //Read port from config.yml

	fmt.Println("starting server at port" + PORT)

	if err := http.ListenAndServe(PORT, nil); err != nil {
		panic(err)
	}
}

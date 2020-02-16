package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
)

var availableConfigPort string

func portAvailable(configPort string) {
	fmt.Println("Setting port to " + configPort)
	ln, err := net.Listen("tcp", ":"+configPort)
	if err != nil {
		fmt.Println("Can't listen on port " + configPort)
		newConfigPort, errStrConv := strconv.Atoi(configPort)
		if errStrConv != nil {
			panic(errStrConv)
		}

		newConfigPort = newConfigPort + 1
		newConfigPortString := strconv.Itoa(newConfigPort)
		portAvailable(newConfigPortString)
	} else {
		_ = ln.Close()
		availableConfigPort = configPort
	}
}

func main() {
	configPort := setPort() //Read port from config.yml

	portAvailable(configPort)

	fmt.Println("Starting Server at port " + availableConfigPort)

	if err := http.ListenAndServe(":"+availableConfigPort, nil); err != nil {
		fmt.Println(err)
	}
}

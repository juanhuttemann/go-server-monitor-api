package main

import (
	"fmt"
	"net"
	"strings"
)

//Address of NetworkDevice
type Address struct {
	IP string `json:"ip"`
}

type Addresses []Address

//NetworkDevice properties
type NetworkDevice struct {
	Name      string    `json:"name"`
	Addresses Addresses `json:"addresses"`
	MAC       string    `json:"mac"`
	UP        bool      `json:"up"`
}

type NetworkDevices []NetworkDevice

func CheckNetworkDevices() NetworkDevices {
	if !available("networkDevices") {
		return NetworkDevices{}
	}
	networkDeviceChan := make(chan NetworkDevices)
	go func(c chan NetworkDevices) {
		var networkDevices NetworkDevices
		var mac string
		interfaces, err := net.Interfaces()

		if err != nil {
			fmt.Print(err)
		}

		netChan := make(chan NetworkDevice)
		for _, iface := range interfaces {
			go func(iface net.Interface) {
				byNameInterface, err := net.InterfaceByName(iface.Name)
				if err != nil {
					fmt.Println(err)
				}

				if iface.HardwareAddr.String() == "" {
					mac = "00:00:00:00:00:00"
				} else {
					mac = iface.HardwareAddr.String()
				}

				networkDevice := NetworkDevice{
					Name: iface.Name,
					MAC:  mac,
				}

				if (iface.Flags & net.FlagUp) == 0 {
					networkDevice.UP = false
				} else {
					networkDevice.UP = true
				}

				addresses, err := byNameInterface.Addrs()

				if err != nil {
					fmt.Println(err)
				}

				var addr Addresses

				for _, address := range addresses {
					addr = append(addr, Address{
						IP: strings.Split(address.String(), "/")[0],
					})
				}

				networkDevice.Addresses = addr

				netChan <- networkDevice
			}(iface)
		}
		taskList := len(interfaces)

		for i := range netChan {
			networkDevices = append(networkDevices, i)
			taskList--
			if taskList == 0 {
				break
			}
		}
		c <- networkDevices
	}(networkDeviceChan)
	return <-networkDeviceChan
}

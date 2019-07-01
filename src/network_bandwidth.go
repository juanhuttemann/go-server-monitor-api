package main

import (
	"runtime"
	"time"

	network "./network"
)

//NetworkDeviceBandwidth properties
type NetworkDeviceBandwidth struct {
	Name string `json:"name"`
	Rx   uint64 `json:"rx"`
	Tx   uint64 `json:"tx"`
}

func CheckNetworkBandwidth() []NetworkDeviceBandwidth {
	if !available("networkBandwidth") {
		return []NetworkDeviceBandwidth{}
	}
	NetworkDeviceBandwidthChan := make(chan []NetworkDeviceBandwidth)

	go func(c chan []NetworkDeviceBandwidth) {

		if runtime.GOOS == "windows" {
			stats := network.GetNetworkStats()
			var networkDevices []NetworkDeviceBandwidth
			for _, dev := range stats.NetDevStats {
				n := NetworkDeviceBandwidth{
					Name: dev.InterfaceName,
					Rx:   dev.RxBytes,
					Tx:   dev.TxBytes,
				}
				networkDevices = append(networkDevices, n)

			}
			c <- networkDevices
		}
		var networkDevices []NetworkDeviceBandwidth

		stat0, err := network.Get()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 1)
		stat1, err := network.Get()
		if err != nil {
			panic(err)
		}

		for i, s := range stat0 {
			n := NetworkDeviceBandwidth{
				Name: s.Name,
				Rx:   stat1[i].RxBytes - s.RxBytes,
				Tx:   stat1[i].TxBytes - s.TxBytes,
			}
			networkDevices = append(networkDevices, n)
			c <- networkDevices

		}
	}(NetworkDeviceBandwidthChan)
	return <-NetworkDeviceBandwidthChan
}

package main

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/host"
)

//HostInfo machine data
type HostInfo struct {
	Name     string `json:"name"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Platform string `json:"platform"`
	Uptime   uint64 `json:"uptime"`
}

func arch() string {
	archChan := make(chan string)
	go func(c chan string) {
		c <- runtime.GOARCH
	}(archChan)
	return <-archChan
}

func CheckHostInfo() HostInfo {
	if !available("hostInfo") {
		return HostInfo{}
	}

	hostInfoChan := make(chan HostInfo)
	go func(c chan HostInfo) {
		hostInfo, err := host.Info()
		if err != nil {
			fmt.Print(err)
		}

		c <- HostInfo{
			Name:     hostInfo.Hostname,
			OS:       hostInfo.OS,
			Arch:     arch(),
			Platform: hostInfo.Platform + " " + hostInfo.PlatformVersion,
			Uptime:   hostInfo.Uptime,
		}
	}(hostInfoChan)
	return <-hostInfoChan
}

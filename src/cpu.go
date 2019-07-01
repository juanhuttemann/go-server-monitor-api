package main

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CPU struct {
	Name         string    `json:"name"`
	Total        int       `json:"total"`
	Usage        float64   `json:"usage"`
	UsagePerCore []float64 `json:"usagePerCore"`
}

func cpuUsageGeneral(c chan float64) {
	duration := 500 * time.Millisecond
	cpuUsage, err := cpu.Percent(duration, false)
	if err != nil {
		panic(err)
	}
	c <- cpuUsage[0]
}

func cpuUsagePerCore(c chan []float64) {
	duration := 500 * time.Millisecond
	cpuUsage, err := cpu.Percent(duration, true)
	if err != nil {
		panic(err)
	}
	c <- cpuUsage
}

func CheckCPU() CPU {
	if !available("cpu") {
		return CPU{}
	}
	cpuChan := make(chan CPU)

	go func(c chan CPU) {
		cpuUsageGeneralChan := make(chan float64)
		cpuUsagePerCoreChan := make(chan []float64)

		go cpuUsageGeneral(cpuUsageGeneralChan)
		go cpuUsagePerCore(cpuUsagePerCoreChan)

		var cpuName string
		if runtime.GOOS == "windows" {
			out, err := exec.Command("wmic", "cpu", "get", "name").Output()
			if err != nil {
				log.Fatal(err)
			}

			cpuName = strings.TrimSpace(strings.Trim(string(out), "Name"))
		} else {
			command := []string{"/proc/cpuinfo", "|", "grep", "model", "|", "head", "-1"}
			out, err := exec.Command("cat", command...).Output()
			if err != nil {
				log.Fatal(err)
			}
			cpuName = strings.TrimSpace(strings.Trim(string(out), "model name"))
		}

		c <- CPU{
			Total:        runtime.NumCPU(),
			Name:         cpuName,
			Usage:        <-cpuUsageGeneralChan,
			UsagePerCore: <-cpuUsagePerCoreChan,
		}

	}(cpuChan)

	return <-cpuChan
}

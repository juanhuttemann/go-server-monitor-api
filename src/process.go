package main

import "github.com/mitchellh/go-ps"

//Process of the OS
type Process struct {
	Pid  int    `json:"pid"`
	Name string `json:"name"`
}

type Processes []Process

func CheckProcesses() Processes {
	if !available("processes") {
		return Processes{}
	}
	processChan := make(chan Processes)
	go func(c chan Processes) {
		processes, err := ps.Processes()
		if err != nil {
			panic(err)
		}
		var processList Processes
		procChan := make(chan Process)
		for _, p := range processes {
			go func(p ps.Process) {
				proc := Process{Pid: p.Pid(), Name: p.Executable()}
				procChan <- proc
			}(p)
		}

		taskList := len(processes)

		for p := range procChan {
			processList = append(processList, p)
			taskList--
			if taskList == 0 {
				break
			}
		}
		c <- processList
	}(processChan)
	return <-processChan
}

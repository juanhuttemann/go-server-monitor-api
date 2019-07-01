package main

import (
	"encoding/json"
	"net/http"
)

type data struct {
	HostInfo         HostInfo                 `json:"hostInfo"`
	CPU              CPU                      `json:"cpu"`
	RAM              RAM                      `json:"ram"`
	Disks            []Disk                   `json:"disks"`
	NetworkDevices   []NetworkDevice          `json:"networkDevices"`
	NetworkBandwidth []NetworkDeviceBandwidth `json:"networkBandwidth"`
	Processes        []Process                `json:"processes"`
}

func ignoreFavicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func checkData(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	data := data{
		HostInfo:         CheckHostInfo(),
		RAM:              CheckRAM(),
		CPU:              CheckCPU(),
		NetworkDevices:   CheckNetworkDevices(),
		NetworkBandwidth: CheckNetworkBandwidth(),
		Disks:            CheckDisks(),
		Processes:        CheckProcesses(),
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func CreateEndPoint(module string, e interface{}) {
	if available(module) {
		endPoint := func(w http.ResponseWriter, r *http.Request) {
			enableCors(&w)
			data := e
			js, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		}

		http.HandleFunc("/"+module, endPoint)
	} else {
		endPoint := func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error 500 - set " + module + " to 'true' in config.yml file"))
		}
		http.HandleFunc("/"+module, endPoint)
	}
}

func init() {
	http.HandleFunc("/", checkData)
	http.HandleFunc("/favicon.ico", ignoreFavicon)
	http.HandleFunc("/host", hostIndex)
	http.HandleFunc("/cpu", cpuIndex)
	http.HandleFunc("/ram", ramIndex)
	http.HandleFunc("/disks", disksIndex)
	http.HandleFunc("/networks", networkIndex)
	http.HandleFunc("/bandwidth", bandwidthIndex)
	http.HandleFunc("/processes", processesIndex)
}

func moduleServer(w http.ResponseWriter, checker interface{}, module string) {
	enableCors(&w)
	if available(module) {
		data := checker
		js, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500 - set " + module + " to 'true' in config.yml file"))
	}
}

func hostIndex(w http.ResponseWriter, r *http.Request) {
	moduleServer(w, CheckHostInfo(), "hostInfo")
}

func ramIndex(w http.ResponseWriter, r *http.Request) {
	moduleServer(w, CheckRAM(), "ram")
}
func cpuIndex(w http.ResponseWriter, r *http.Request) {
	moduleServer(w, CheckCPU(), "cpu")
}

func disksIndex(w http.ResponseWriter, r *http.Request) {
	moduleServer(w, CheckDisks(), "disks")
}

func networkIndex(w http.ResponseWriter, r *http.Request) {
	moduleServer(w, CheckDisks(), "networkDevices")
}

func bandwidthIndex(w http.ResponseWriter, r *http.Request) {
	moduleServer(w, CheckNetworkBandwidth(), "networkBandwidth")
}

func processesIndex(w http.ResponseWriter, r *http.Request) {
	moduleServer(w, CheckProcesses(), "processes")
}

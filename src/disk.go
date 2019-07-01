package main

import "github.com/shirou/gopsutil/disk"

//Disk properties
type Disk struct {
	Mountpoint string  `json:"mountPoint"`
	Free       uint64  `json:"free"`
	Size       uint64  `json:"size"`
	Used       uint64  `json:"used"`
	Percent    float64 `json:"percent"`
}

type Disks []Disk

func CheckDisks() Disks {
	if !available("disks") {
		return Disks{}
	}
	disksChan := make(chan Disks)
	go func(c chan Disks) {
		disks, _ := disk.Partitions(false)

		var totalDisks Disks

		diskChan := make(chan Disk)
		for _, d := range disks {
			go func(d disk.PartitionStat) {
				diskUsageOf, _ := disk.Usage(d.Mountpoint)
				diskChan <- Disk{
					Free:       diskUsageOf.Free,
					Mountpoint: d.Mountpoint,
					Percent:    diskUsageOf.UsedPercent,
					Size:       diskUsageOf.Total,
					Used:       diskUsageOf.Used,
				}

			}(d)

		}
		taskList := len(disks)

		for d := range diskChan {
			totalDisks = append(totalDisks, d)
			taskList--
			if taskList == 0 {
				break
			}
		}
		c <- totalDisks
	}(disksChan)

	return <-disksChan
}

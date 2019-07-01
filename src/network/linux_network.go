// +build linux

package network

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//FIXME
// AVOID LINUX COMPLAINING ABOUT SYSCALL
type EthrNetDevStat struct {
	InterfaceName string
	RxBytes       uint64
	TxBytes       uint64
	RxPkts        uint64
	TxPkts        uint64
}

type EthrNetStat struct {
	NetDevStats []EthrNetDevStat
}

func GetNetworkStats() EthrNetStat {
	return EthrNetStat{}
}

//FIXME
// AVOID LINUX COMPLAINING ABOUT SYSCALL

func Get() ([]Stats, error) {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return collectNetworkStats(file)
}

type Stats struct {
	Name             string
	RxBytes, TxBytes uint64
}

func collectNetworkStats(out io.Reader) ([]Stats, error) {
	scanner := bufio.NewScanner(out)
	var networks []Stats
	for scanner.Scan() {
		kv := strings.SplitN(scanner.Text(), ":", 2)
		if len(kv) != 2 {
			continue
		}
		fields := strings.Fields(kv[1])
		if len(fields) < 16 {
			continue
		}
		name := strings.TrimSpace(kv[0])
		if name == "lo" {
			continue
		}
		rxBytes, err := strconv.ParseUint(fields[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse rxBytes of %s", name)
		}
		txBytes, err := strconv.ParseUint(fields[8], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse txBytes of %s", name)
		}
		networks = append(networks, Stats{Name: name, RxBytes: rxBytes, TxBytes: txBytes})
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan error for /proc/net/dev: %s", err)
	}
	return networks, nil
}

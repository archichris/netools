package dev

import (
	"encoding/json"
	// "fmt"
	"net"
)

type HardwareAddr net.HardwareAddr

func (n HardwareAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal((*net.HardwareAddr)(&n).String())
}

func (n *HardwareAddr) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	tmp, err := net.ParseMAC(s)
	if err != nil {
		return err
	}

	*n = HardwareAddr(tmp)
	return nil
}

func (n *HardwareAddr) String() string {
	return (*net.HardwareAddr)(n).String()
}

func (n *HardwareAddr) ToNetHardwareAddr() *net.HardwareAddr {
	return (*net.HardwareAddr)(n)
}

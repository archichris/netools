package ipaddr

import (
	"math"
	"math/big"
	"net"
	"strconv"
)

// IP4ToUint32 convert net.IP to uint32
func IP4ToUint32(ip net.IP) uint32 {
	if v := ip.To4(); v != nil {
		return uint32(big.NewInt(0).SetBytes(v).Uint64())
	}
	return 0
}

// Uint32ToIP4 convert uint32 to net.IP
func Uint32ToIP4(i uint32) net.IP {
	return net.IP(big.NewInt(0).SetUint64(uint64(i)).Bytes())
}

// Uint32AddSeg return the result of input IP plus an input ip segment
func Uint32AddSeg(ip uint32, s uint32) uint32 {
	return ip + uint32(2<<(s-1))
}

// Net4ToStartEndUint32 convert net to start ip and end ip in unit32 format
func Net4ToStartEndUint32(n *net.IPNet) (uint32, uint32) {
	ipStart := IP4ToUint32(n.IP.Mask(n.Mask))
	ones, bits := n.Mask.Size()
	ipEnd := Uint32AddSeg(ipStart, uint32(bits-ones)) - 1
	return ipStart, ipEnd
}

// StrToUint32 convert string format ip to uint32 format
func StrToUint32(s string) uint32 {
	tmpU64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(tmpU64)
}

// GetHostUint32 get the host part from net in uint32 format
func GetHostUint32(n *net.IPNet) (uint32, uint32) {
	ones, bits := n.Mask.Size()
	return IP4ToUint32(n.IP) & uint32(math.Pow(float64(2), float64(bits-ones))-1), uint32(bits - ones)
}

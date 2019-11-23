package dev_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net"

	. "github.com/archichris/netools/dev"
)

var _ = Describe("Device", func() {
	var (
		localIP   = "192.168.56.10"
		localMask = "ffffff00"
	)
	It("get net of the given interface", func() {

		iface, _ := net.InterfaceByName("eth0")
		n, _ := GetIfaceIP4Addr(iface)
		Expect(n.IP.String()).To(Equal(localIP))
		Expect(n.Mask.String()).To(Equal(localMask))
	})
})

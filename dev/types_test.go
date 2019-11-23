package dev_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	// "fmt"
	. "github.com/archichris/netools/dev"
)

var _ = Describe("Types", func() {
	type Info struct {
		MAC HardwareAddr `json:"mac"`
	}
	var example = []byte(`
{
	"mac": "0a:0a:0a:0b:0c:0d"
}
`)

	It("marshal mac string", func() {
		mac := HardwareAddr([]byte{0xa, 0xa, 0xa, 0xb, 0xc, 0xd})
		str, _ := mac.MarshalJSON()
		Expect(str).To(ContainSubstring("0a:0a:0a:0b:0c:0d"))
		info := Info{MAC: mac}
		str, _ = json.Marshal(info)
		Expect(str).To(ContainSubstring(`"mac":"0a:0a:0a:0b:0c:0d"`))
	})

	It("unmarshal string to mac", func() {
		var mac HardwareAddr
		err := mac.UnmarshalJSON([]byte(`"0a:0a:0a:0b:0c:0d"`))
		Expect(err).To(BeNil())
		Expect(mac.String()).To(Equal("0a:0a:0a:0b:0c:0d"))
		info := Info{}
		json.Unmarshal(example, &info)
		Expect(info.MAC.String()).To(Equal("0a:0a:0a:0b:0c:0d"))
	})
})

// pcap包提供一些对于pcap常用操作的封装。
package pcap

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

// 用pcap包寻找网络适配器。
func FindDevices() (device string, err error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return
	}
	fmt.Println("All network adapters:")
	for i := 0; i < len(devices); i++ {
		fmt.Print(i, ".", devices[i].Name, " ", devices[i].Description, "\n")
	}
	fmt.Println("Please select a network adapter(Enter number):")
	index := 0
	for {
		_, err = fmt.Scanln(&index)
		if err != nil || index < 0 || (index+1) > len(devices) {
			fmt.Println("Please enter the correct number:")
			continue
		}
		break
	}

	device = devices[index].Name
	return
}

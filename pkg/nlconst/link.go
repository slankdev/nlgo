package nlconst

import (
	"fmt"
	"syscall"
)

type IfinfoType uint16
type IfinfoFlags uint32

func (t IfinfoType) String() string {
	switch t {
	case 1:
		return "Ethernet"
	default:
		return fmt.Sprintf("Unknown(%d)", t)
	}
}

func (f IfinfoFlags) String() string {
	str := ""
	if f&syscall.IFF_UP != 0 {
		str += "UP|"
	}
	if f&syscall.IFF_BROADCAST != 0 {
		str += "BROADCAST|"
	}
	if f&syscall.IFF_MASTER != 0 {
		str += "MASTER|"
	}
	if f&syscall.IFF_SLAVE != 0 {
		str += "SLAVE|"
	}
	if f&syscall.IFF_MULTICAST != 0 {
		str += "MULTICAST|"
	}
	if f&syscall.IFF_RUNNING != 0 {
		str += "RUNNING|"
	}
	if f&syscall.IFF_LOOPBACK != 0 {
		str += "LOOPBACK|"
	}
	if f&syscall.IFF_NOARP != 0 {
		str += "NOARP|"
	}
	if f&syscall.IFF_PROMISC != 0 {
		str += "PROMISC|"
	}

	if str != "" {
		str = str[:len(str)-1]
	}
	return fmt.Sprintf("<(%#x)%s>", int(f), str)
}

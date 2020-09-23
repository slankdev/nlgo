package nlmsg

import (
	"fmt"
	"syscall"

	"github.com/slankdev/nlgo/pkg/util"
)

type callback func([]byte) string

type Context struct {
	nm string
	fn callback
}

var TableIfinfo = map[uint16]Context{
	syscall.IFLA_MTU:       {nm: "MTU", fn: util.ByteToU32},
	syscall.IFLA_IFNAME:    {nm: "IFNAME", fn: util.ByteToString},
	syscall.IFLA_OPERSTATE: {nm: "OPERSTATE", fn: util.ByteToU8},
	syscall.IFLA_TXQLEN:    {nm: "TXQLEN", fn: util.ByteToU32},
}

func (a Attr) String(tab map[uint16]Context) string {
	ctx, ok := tab[a.Type]
	if ok {
		return fmt.Sprintf("%s(%s)", ctx.nm, ctx.fn(a.Data))
	}
	return "unknown"
}

package nlmsg

import (
	"fmt"

	"github.com/slankdev/nlgo/pkg/util"
)

type callback func([]byte) string

type Context struct {
	nm string
	fn callback
}

var TableIfinfo = map[uint16]Context{
	IFLA_MTU:           {nm: "MTU", fn: util.ByteToU32},
	IFLA_IFNAME:        {nm: "IFNAME", fn: util.ByteToString},
	IFLA_OPERSTATE:     {nm: "OPERSTATE", fn: util.ByteToU8},
	IFLA_TXQLEN:        {nm: "TXQLEN", fn: util.ByteToU32},
	IFLA_LINKMODE:      {nm: "LINKMODE", fn: util.ByteToU8},
	IFLA_MIN_MTU:       {nm: "MIN_MTU", fn: util.ByteToU32},
	IFLA_MAX_MTU:       {nm: "MAX_MTU", fn: util.ByteToU32},
	IFLA_GROUP:         {nm: "GROUP", fn: util.ByteToU32},
	IFLA_PROMISCUITY:   {nm: "PROMISCUITY", fn: util.ByteToU32},
	IFLA_NUM_TX_QUEUES: {nm: "NUM_TX_QUEUES", fn: util.ByteToU32},
}

func (a Attr) String(tab map[uint16]Context) string {
	ctx, ok := tab[a.Type]
	if ok {
		return fmt.Sprintf("%s(%s)", ctx.nm, ctx.fn(a.Data))
	}
	return "unknown"
}

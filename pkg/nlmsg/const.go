package nlmsg

const (
	IFLA_UNSPEC             = 0x0000
	IFLA_ADDRESS            = 0x0001
	IFLA_BROADCAST          = 0x0002
	IFLA_IFNAME             = 0x0003
	IFLA_MTU                = 0x0004
	IFLA_LINK               = 0x0005
	IFLA_QDISC              = 0x0006
	IFLA_STATS              = 0x0007
	IFLA_COST               = 0x0008
	IFLA_PRIORITY           = 0x0009
	IFLA_MASTER             = 0x000a
	IFLA_WIRELESS           = 0x000b
	IFLA_PROTINFO           = 0x000c
	IFLA_TXQLEN             = 0x000d
	IFLA_MAP                = 0x000e
	IFLA_WEIGHT             = 0x000f
	IFLA_OPERSTATE          = 0x0010
	IFLA_LINKMODE           = 0x0011
	IFLA_LINKINFO           = 0x0012
	IFLA_NET_NS_PID         = 0x0013
	IFLA_IFALIAS            = 0x0014
	IFLA_NUM_VF             = 0x0015
	IFLA_VFINFO_LIST        = 0x0016
	IFLA_STATS64            = 0x0017
	IFLA_VF_PORTS           = 0x0018
	IFLA_PORT_SELF          = 0x0019
	IFLA_AF_SPEC            = 0x001a
	IFLA_GROUP              = 0x001b
	IFLA_NET_NS_FD          = 0x001c
	IFLA_EXT_MASK           = 0x001d
	IFLA_PROMISCUITY        = 0x001e
	IFLA_NUM_TX_QUEUES      = 0x001f
	IFLA_NUM_RX_QUEUES      = 0x0020
	IFLA_CARRIER            = 0x0021
	IFLA_PHYS_PORT_ID       = 0x0022
	IFLA_CARRIER_CHANGES    = 0x0023
	IFLA_PHYS_SWITCH_ID     = 0x0024
	IFLA_LINK_NETNSID       = 0x0025
	IFLA_PHYS_PORT_NAME     = 0x0026
	IFLA_PROTO_DOWN         = 0x0027
	IFLA_GSO_MAX_SEGS       = 0x0028
	IFLA_GSO_MAX_SIZE       = 0x0029
	IFLA_PAD                = 0x002a
	IFLA_XDP                = 0x002b
	IFLA_EVENT              = 0x002c
	IFLA_NEW_NETNSID        = 0x002d
	IFLA_IF_NETNSID         = 0x002e
	IFLA_CARRIER_UP_COUNT   = 0x002f
	IFLA_CARRIER_DOWN_COUNT = 0x0030
	IFLA_NEW_IFINDEX        = 0x0031
	IFLA_MIN_MTU            = 0x0032
	IFLA_MAX_MTU            = 0x0033
	IFLA_PROP_LIST          = 0x0034
	IFLA_ALT_IFNAME         = 0x0035
	IFLA_PERM_ADDRESS       = 0x0036
	IFLA_PROTO_DOWN_REASON  = 0x0037
	__IFLA_MAX              = 0x0038
)
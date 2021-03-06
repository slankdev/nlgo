package nlconst

import (
	"fmt"
	"syscall"
)

type NlmsgType uint16
type NlmsgFlags uint16

const (
	RTM_BASE = 16

	RTM_NEWLINK = 16
	RTM_DELLINK = 17
	RTM_GETLINK = 18
	RTM_SETLINK = 19

	RTM_NEWADDR = 20
	RTM_DELADDR = 21
	RTM_GETADDR = 22

	RTM_NEWROUTE = 24
	RTM_DELROUTE = 25
	RTM_GETROUTE = 26

	RTM_NEWNEIGH = 28
	RTM_DELNEIGH = 29
	RTM_GETNEIGH = 30

	RTM_NEWRULE = 32
	RTM_DELRULE = 33
	RTM_GETRULE = 34

	RTM_NEWQDISC = 36
	RTM_DELQDISC = 37
	RTM_GETQDISC = 38

	RTM_NEWTCLASS = 40
	RTM_DELTCLASS = 41
	RTM_GETTCLASS = 42

	RTM_NEWTFILTER = 44
	RTM_DELTFILTER = 45
	RTM_GETTFILTER = 46

	RTM_NEWACTION = 48
	RTM_DELACTION = 49
	RTM_GETACTION = 50

	RTM_NEWPREFIX = 52

	RTM_GETMULTICAST = 58
	RTM_GETANYCAST   = 62

	RTM_NEWNEIGHTBL = 64
	RTM_GETNEIGHTBL = 66
	RTM_SETNEIGHTBL = 67

	RTM_NEWNDUSEROPT = 68

	RTM_NEWADDRLABEL = 72
	RTM_DELADDRLABEL = 73
	RTM_GETADDRLABEL = 74

	RTM_GETDCB = 78
	RTM_SETDCB = 79

	RTM_NEWNETCONF = 80
	RTM_DELNETCONF = 81
	RTM_GETNETCONF = 82

	RTM_NEWMDB = 84
	RTM_DELMDB = 85
	RTM_GETMDB = 86

	RTM_NEWNSID = 88
	RTM_DELNSID = 89
	RTM_GETNSID = 90

	RTM_NEWSTATS = 92
	RTM_GETSTATS = 94

	RTM_NEWCACHEREPORT = 96

	RTM_NEWCHAIN = 100
	RTM_DELCHAIN = 101
	RTM_GETCHAIN = 102

	RTM_NEWNEXTHOP = 104
	RTM_DELNEXTHOP = 105
	RTM_GETNEXTHOP = 106

	RTM_NEWLINKPROP = 108
	RTM_DELLINKPROP = 109
	RTM_GETLINKPROP = 110

	RTM_NEWVLAN = 112
	RTM_DELVLAN = 113
	RTM_GETVLAN = 114

	__RTM_MAX
)

func (t NlmsgType) String() string {
	switch t {
	case RTM_NEWLINK:
		return "RTM_NEWLINK"
	case RTM_DELLINK:
		return "RTM_DELLINK"
	case RTM_GETLINK:
		return "RTM_GETLINK"
	case RTM_SETLINK:
		return "RTM_SETLINK"
	case RTM_NEWADDR:
		return "RTM_NEWADDR"
	case RTM_DELADDR:
		return "RTM_DELADDR"
	case RTM_GETADDR:
		return "RTM_GETADDR"
	case RTM_NEWROUTE:
		return "RTM_NEWROUTE"
	case RTM_DELROUTE:
		return "RTM_DELROUTE"
	case RTM_GETROUTE:
		return "RTM_GETROUTE"
	case RTM_NEWNEIGH:
		return "RTM_NEWNEIGH"
	case RTM_DELNEIGH:
		return "RTM_DELNEIGH"
	case RTM_GETNEIGH:
		return "RTM_GETNEIGH"
	case RTM_NEWRULE:
		return "RTM_NEWRULE"
	case RTM_DELRULE:
		return "RTM_DELRULE"
	case RTM_GETRULE:
		return "RTM_GETRULE"
	case RTM_NEWQDISC:
		return "RTM_NEWQDISC"
	case RTM_DELQDISC:
		return "RTM_DELQDISC"
	case RTM_GETQDISC:
		return "RTM_GETQDISC"
	case RTM_NEWTCLASS:
		return "RTM_NEWTCLASS"
	case RTM_DELTCLASS:
		return "RTM_DELTCLASS"
	case RTM_GETTCLASS:
		return "RTM_GETTCLASS"
	case RTM_NEWTFILTER:
		return "RTM_NEWTFILTER"
	case RTM_DELTFILTER:
		return "RTM_DELTFILTER"
	case RTM_GETTFILTER:
		return "RTM_GETTFILTER"
	case RTM_NEWACTION:
		return "RTM_NEWACTION"
	case RTM_DELACTION:
		return "RTM_DELACTION"
	case RTM_GETACTION:
		return "RTM_GETACTION"
	case RTM_NEWPREFIX:
		return "RTM_NEWPREFIX"
	case RTM_GETMULTICAST:
		return "RTM_GETMULTICAST"
	case RTM_GETANYCAST:
		return "RTM_GETANYCAST"
	case RTM_NEWNEIGHTBL:
		return "RTM_NEWNEIGHTBL"
	case RTM_GETNEIGHTBL:
		return "RTM_GETNEIGHTBL"
	case RTM_SETNEIGHTBL:
		return "RTM_SETNEIGHTBL"
	case RTM_NEWNDUSEROPT:
		return "RTM_NEWNDUSEROPT"
	case RTM_NEWADDRLABEL:
		return "RTM_NEWADDRLABEL"
	case RTM_DELADDRLABEL:
		return "RTM_DELADDRLABEL"
	case RTM_GETADDRLABEL:
		return "RTM_GETADDRLABEL"
	case RTM_GETDCB:
		return "RTM_GETDCB"
	case RTM_SETDCB:
		return "RTM_SETDCB"
	case RTM_NEWNETCONF:
		return "RTM_NEWNETCONF"
	case RTM_DELNETCONF:
		return "RTM_DELNETCONF"
	case RTM_GETNETCONF:
		return "RTM_GETNETCONF"
	case RTM_NEWMDB:
		return "RTM_NEWMDB"
	case RTM_DELMDB:
		return "RTM_DELMDB"
	case RTM_GETMDB:
		return "RTM_GETMDB"
	case RTM_NEWNSID:
		return "RTM_NEWNSID"
	case RTM_DELNSID:
		return "RTM_DELNSID"
	case RTM_GETNSID:
		return "RTM_GETNSID"
	case RTM_NEWSTATS:
		return "RTM_NEWSTATS"
	case RTM_GETSTATS:
		return "RTM_GETSTATS"
	case RTM_NEWCACHEREPORT:
		return "RTM_NEWCACHEREPORT"
	case RTM_NEWCHAIN:
		return "RTM_NEWCHAIN"
	case RTM_DELCHAIN:
		return "RTM_DELCHAIN"
	case RTM_GETCHAIN:
		return "RTM_GETCHAIN"
	case RTM_NEWNEXTHOP:
		return "RTM_NEWNEXTHOP"
	case RTM_DELNEXTHOP:
		return "RTM_DELNEXTHOP"
	case RTM_GETNEXTHOP:
		return "RTM_GETNEXTHOP"
	case RTM_NEWLINKPROP:
		return "RTM_NEWLINKPROP"
	case RTM_DELLINKPROP:
		return "RTM_DELLINKPROP"
	case RTM_GETLINKPROP:
		return "RTM_GETLINKPROP"
	case RTM_NEWVLAN:
		return "RTM_NEWVLAN"
	case RTM_DELVLAN:
		return "RTM_DELVLAN"
	case RTM_GETVLAN:
		return "RTM_GETVLAN"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", t)
	}
}

func (f NlmsgFlags) String() string {
	str := ""
	if f&syscall.NLM_F_REQUEST != 0 {
		f -= syscall.NLM_F_REQUEST
		str += "REQUEST|"
	}
	if f&syscall.NLM_F_ACK != 0 {
		f -= syscall.NLM_F_ACK
		str += "ACK|"
	}
	if f&syscall.NLM_F_APPEND != 0 {
		f -= syscall.NLM_F_APPEND
		str += "APPEND|"
	}
	if f&syscall.NLM_F_ATOMIC != 0 {
		f -= syscall.NLM_F_ATOMIC
		str += "ATOMIC|"
	}
	if f&syscall.NLM_F_CREATE != 0 {
		f -= syscall.NLM_F_CREATE
		str += "CREATE|"
	}
	if f&syscall.NLM_F_DUMP != 0 {
		f -= syscall.NLM_F_DUMP
		str += "DUMP|"
	}
	if f&syscall.NLM_F_ECHO != 0 {
		f -= syscall.NLM_F_ECHO
		str += "ECHO|"
	}
	if f&syscall.NLM_F_EXCL != 0 {
		f -= syscall.NLM_F_EXCL
		str += "EXCL|"
	}
	if f&syscall.NLM_F_MATCH != 0 {
		f -= syscall.NLM_F_MATCH
		str += "MATCH|"
	}
	if f&syscall.NLM_F_MULTI != 0 {
		f -= syscall.NLM_F_MULTI
		str += "MULTI|"
	}
	if f&syscall.NLM_F_REPLACE != 0 {
		f -= syscall.NLM_F_REPLACE
		str += "REPLACE|"
	}
	if f&syscall.NLM_F_REQUEST != 0 {
		f -= syscall.NLM_F_REQUEST
		str += "REQUEST|"
	}
	if f&syscall.NLM_F_ROOT != 0 {
		f -= syscall.NLM_F_ROOT
		str += "ROOT|"
	}
	if f != 0 {
		str += fmt.Sprintf("UNKNOWN(%#x)|", int(f))
	}

	if str == "" {
		str = fmt.Sprintf("<(%d)>", f)
	} else {
		str = fmt.Sprintf("<(%d)%s>", f, str[:len(str)-1])
	}
	return str
}

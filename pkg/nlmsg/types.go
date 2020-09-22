package nlmsg

import (
	"fmt"
	"syscall"
)

type NlmsgType uint16

func (nt NlmsgType) String() string {
	switch nt {
	case syscall.RTM_DELACTION:
		return "RTM_DELACTION"
	case syscall.RTM_DELADDR:
		return "RTM_DELADDR"
	case syscall.RTM_DELADDRLABEL:
		return "RTM_DELADDRLABEL"
	case syscall.RTM_DELLINK:
		return "RTM_DELLINK"
	case syscall.RTM_DELNEIGH:
		return "RTM_DELNEIGH"
	case syscall.RTM_DELQDISC:
		return "RTM_DELQDISC"
	case syscall.RTM_DELROUTE:
		return "RTM_DELROUTE"
	case syscall.RTM_DELRULE:
		return "RTM_DELRULE"
	case syscall.RTM_DELTCLASS:
		return "RTM_DELTCLASS"
	case syscall.RTM_DELTFILTER:
		return "RTM_DELTFILTER"
	case syscall.RTM_GETACTION:
		return "RTM_GETACTION"
	case syscall.RTM_GETADDR:
		return "RTM_GETADDR"
	case syscall.RTM_GETADDRLABEL:
		return "RTM_GETADDRLABEL"
	case syscall.RTM_GETANYCAST:
		return "RTM_GETANYCAST"
	case syscall.RTM_GETDCB:
		return "RTM_GETDCB"
	case syscall.RTM_GETLINK:
		return "RTM_GETLINK"
	case syscall.RTM_GETMULTICAST:
		return "RTM_GETMULTICAST"
	case syscall.RTM_GETNEIGH:
		return "RTM_GETNEIGH"
	case syscall.RTM_GETNEIGHTBL:
		return "RTM_GETNEIGHTBL"
	case syscall.RTM_GETQDISC:
		return "RTM_GETQDISC"
	case syscall.RTM_GETROUTE:
		return "RTM_GETROUTE"
	case syscall.RTM_GETRULE:
		return "RTM_GETRULE"
	case syscall.RTM_GETTCLASS:
		return "RTM_GETTCLASS"
	case syscall.RTM_GETTFILTER:
		return "RTM_GETTFILTER"
	case syscall.RTM_MAX:
		return "RTM_MAX"
	case syscall.RTM_NEWACTION:
		return "RTM_NEWACTION"
	case syscall.RTM_NEWADDR:
		return "RTM_NEWADDR"
	case syscall.RTM_NEWADDRLABEL:
		return "RTM_NEWADDRLABEL"
	case syscall.RTM_NEWLINK:
		return "RTM_NEWLINK"
	case syscall.RTM_NEWNDUSEROPT:
		return "RTM_NEWNDUSEROPT"
	case syscall.RTM_NEWNEIGH:
		return "RTM_NEWNEIGH"
	case syscall.RTM_NEWNEIGHTBL:
		return "RTM_NEWNEIGHTBL"
	case syscall.RTM_NEWPREFIX:
		return "RTM_NEWPREFIX"
	case syscall.RTM_NEWQDISC:
		return "RTM_NEWQDISC"
	case syscall.RTM_NEWROUTE:
		return "RTM_NEWROUTE"
	case syscall.RTM_NEWRULE:
		return "RTM_NEWRULE"
	case syscall.RTM_NEWTCLASS:
		return "RTM_NEWTCLASS"
	case syscall.RTM_NEWTFILTER:
		return "RTM_NEWTFILTER"
	case syscall.RTM_SETLINK:
		return "RTM_SETLINK"
	case syscall.RTM_SETNEIGHTBL:
		return "RTM_SETNEIGHTBL"
	}
	return fmt.Sprintf("UNKNOWN(%d)", nt)
}

type Header struct {
	Len   uint32
	Type  NlmsgType
	Flags uint16
	Seq   uint32
	Pid   uint32
}

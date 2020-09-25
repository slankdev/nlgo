package nlmsg

import (
	"fmt"
	"syscall"

	"github.com/slankdev/nlgo/pkg/nlconst"
)

type NlmsgFlags uint16

type Header struct {
	Len   uint32
	Type  nlconst.NlmsgType
	Flags NlmsgFlags
	Seq   uint32
	Pid   uint32
}

func (f NlmsgFlags) String() string {
	str := ""
	if f&syscall.NLM_F_REQUEST != 0 {
		str += "REQUEST|"
	}
	if f&syscall.NLM_F_ACK != 0 {
		str += "ACK|"
	}
	if f&syscall.NLM_F_APPEND != 0 {
		str += "APPEND|"
	}
	if f&syscall.NLM_F_ATOMIC != 0 {
		str += "ATOMIC|"
	}
	if f&syscall.NLM_F_CREATE != 0 {
		str += "CREATE|"
	}
	if f&syscall.NLM_F_DUMP != 0 {
		str += "DUMP|"
	}
	if f&syscall.NLM_F_ECHO != 0 {
		str += "ECHO|"
	}
	if f&syscall.NLM_F_EXCL != 0 {
		str += "EXCL|"
	}
	if f&syscall.NLM_F_MATCH != 0 {
		str += "MATCH|"
	}
	if f&syscall.NLM_F_MULTI != 0 {
		str += "MULTI|"
	}
	if f&syscall.NLM_F_REPLACE != 0 {
		str += "REPLACE|"
	}
	if f&syscall.NLM_F_REQUEST != 0 {
		str += "REQUEST|"
	}
	if f&syscall.NLM_F_ROOT != 0 {
		str += "ROOT|"
	}

	if str == "" {
		str = fmt.Sprintf("<(%d)>", f)
	} else {
		str = fmt.Sprintf("<(%d)%s>", f, str[:len(str)-2])
	}
	return str
}

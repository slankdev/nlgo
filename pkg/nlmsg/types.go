package nlmsg

import (
	"github.com/slankdev/nlgo/pkg/nlconst"
)

type Header struct {
	Len   uint32
	Type  nlconst.NlmsgType
	Flags nlconst.NlmsgFlags
	Seq   uint32
	Pid   uint32
}

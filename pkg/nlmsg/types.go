package nlmsg

type Header struct {
	Len   uint32
	Type  uint32
	Flags uint16
	Seq   uint32
	Pid   uint32
}

package nlsock

import (
	"golang.org/x/sys/unix"
)

type Socket struct {
	fd int
}

func NewSocket(family int, group uint32) (*Socket, error) {
	var err error
	s := &Socket{}

	s.fd, err = unix.Socket(
		unix.AF_NETLINK,
		unix.SOCK_RAW,
		family,
	)
	if err != nil {
		return nil, err
	}

	addr := &unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
		Groups: group,
	}
	err = unix.Bind(s.fd, addr)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Socket) Close() error {
	return unix.Close(s.fd)
}

func (s *Socket) Receive() ([]byte, error) {
	return nil, nil
}

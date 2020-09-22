package nldump

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"syscall"

	"github.com/slankdev/nlgo/pkg/nlmsg"
	"github.com/slankdev/nlgo/pkg/nlsock"
	"github.com/spf13/cobra"
)

var config struct {
	ifname string

	sock *nlsock.Socket
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "nldump",
		RunE: appMain,
	}
	cmd.Flags().StringVarP(&config.ifname, "interface", "i", "lo", "interface to capture")
	return cmd
}

func appMain(cmd *cobra.Command, args []string) error {
	var err error
	config.sock, err = nlsock.NewSocket(0x00, 0xffffffff)
	if err != nil {
		return err
	}

	for {
		b, err := config.sock.Receive()
		if err != nil {
			return err
		}

		dump(b)
	}
	return nil
}

func dump(b []byte) {
	hdr := nlmsg.Header{}
	reader := bytes.NewReader(b)
	binary.Read(reader, binary.LittleEndian, &hdr)
	fmt.Println("-----")
	fmt.Printf("HDR: %+v\n", hdr)

	switch hdr.Type {
	case syscall.RTM_NEWLINK, syscall.RTM_DELLINK:
		ifi := nlmsg.Ifinfo{}
		binary.Read(reader, binary.LittleEndian, &ifi)
		fmt.Printf("  LINK: %+v\n", ifi)
	case syscall.RTM_NEWROUTE, syscall.RTM_DELROUTE:
		fmt.Printf("  ROUTE: \n")
	}

	fmt.Printf(hex.Dump(b[0:64]))
}

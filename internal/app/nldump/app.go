package nldump

import (
	"github.com/k0kubun/pp"
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
	config.sock, err = nlsock.NewSocket(0x00, 0x0F)
	if err != nil {
		return err
	}

	pp.Println(config.sock)
	return nil
}

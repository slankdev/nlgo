package nldump

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
)

var config struct {
	ifname string
}

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "cmd",
		RunE: appMain,
	}
	cmd.Flags().StringVarP(&config.ifname, "interface", "i", "lo", "interface to capture")
	return cmd
}

func appMain(cmd *cobra.Command, args []string) error {
	fmt.Printf("ifname=%s\n", config.ifname)

	handle, err := pcap.OpenLive(config.ifname, 1600, true, pcap.BlockForever)
	if err != nil {
		return err
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}

	return nil
}

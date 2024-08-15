package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vndr/jv/pkg/ip"
)

// NewIPCmd creates the IP command.
func NewIPCmd() *cobra.Command {
	ipCmd := &cobra.Command{
		Use:   "ip",
		Short: "IP related commands",
	}

	// Add child commands to ipCmd
	ipCmd.AddCommand(NewLocalIPCmd(ip.GetLocalIPWithInterfaceCheck))
	ipCmd.AddCommand(NewPublicIPCmd(ip.GetPublicIP))

	return ipCmd
}

package cmd

import (
	"github.com/vndr/jv/pkg/ip"

	"github.com/spf13/cobra"
)

// NewLocalIPCmd creates the local IP command.
func NewLocalIPCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "local",
		Short: "Get the local IP address",
		Run: func(cmd *cobra.Command, args []string) {
			ip.GetLocalIPWithInterfaceCheck()
		},
	}
}

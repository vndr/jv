package cmd

import (
	"github.com/vndr/jv/pkg/ip"

	"github.com/spf13/cobra"
)

// NewPublicIPCmd creates the public IP command.
func NewPublicIPCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "public",
		Short: "Get the public IP address",
		Run: func(cmd *cobra.Command, args []string) {
			ip.GetPublicIP()
		},
	}
}

package cmd

import "github.com/spf13/cobra"

// NewIPCmd creates the IP command.
func NewIPCmd() *cobra.Command {
	ipCmd := &cobra.Command{
		Use:   "ip",
		Short: "IP related commands",
	}

	// Add child commands to ipCmd
	ipCmd.AddCommand(NewLocalIPCmd())
	ipCmd.AddCommand(NewPublicIPCmd())

	return ipCmd
}

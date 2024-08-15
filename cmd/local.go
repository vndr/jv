package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// NewLocalIPCmd creates the local IP command.
func NewLocalIPCmd(getIPFunc func()) *cobra.Command {
	return &cobra.Command{
		Use:   "local",
		Short: "Get the local IP address",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return errors.New("unexpected arguments provided")
			}
			getIPFunc()
			return nil
		},
	}
}

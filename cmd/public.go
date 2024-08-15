package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// NewPublicIPCmd creates the public IP command.
func NewPublicIPCmd(getIPFunc func()) *cobra.Command {
	return &cobra.Command{
		Use:   "public",
		Short: "Get the public IP address",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return errors.New("unexpected arguments provided")
			}
			getIPFunc()
			return nil
		},
	}
}

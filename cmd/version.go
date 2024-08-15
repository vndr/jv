package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// NewVersionCmd creates the version command.
func NewVersionCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of JV Tool",
		Long:  `All software has versions. This is JV Tool's version.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				return fmt.Errorf("unexpected arguments provided")
			}

			// If version already starts with "v", don't prepend another "v"
			if strings.HasPrefix(version, "v") {
				fmt.Fprintf(cmd.OutOrStdout(), "JV Tool %s\n", version)
			} else {
				fmt.Fprintf(cmd.OutOrStdout(), "JV Tool v%s\n", version)
			}
			return nil
		},
	}
}

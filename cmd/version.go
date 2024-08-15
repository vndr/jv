package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewVersionCmd creates the version command.
func NewVersionCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of JV Tool",
		Long:  `All software has versions. This is JV Tool's version.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("JV Tool v%s\n", version)
		},
	}
}

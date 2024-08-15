package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCmd creates the root command for the JV tool.
func NewRootCmd(version string) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "jv",
		Short: "JV Command Line Tool",
		Long:  `A CLI tool to fetch IP addresses, manage network interfaces, and more.`,
	}

	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version number of JV Tool")

	// Add child commands
	rootCmd.AddCommand(NewVersionCmd(version))
	rootCmd.AddCommand(NewIPCmd())

	return rootCmd
}

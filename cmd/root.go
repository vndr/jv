package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// NewRootCmd creates the root command for the JV tool.
func NewRootCmd(version string) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "jv",
		Short: "JV Command Line Tool",
		Long:  `A CLI tool to fetch IP addresses, manage network interfaces, and more.`,
	}

	// Handle the --version flag explicitly
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if versionFlag, _ := cmd.Flags().GetBool("version"); versionFlag {
			fmt.Printf("JV Tool v%s\n", version)
			os.Exit(0)
		}
		return nil
	}

	// Define the --version flag
	//	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version number of JV Tool")

	// Add child commands
	rootCmd.AddCommand(NewVersionCmd(version))
	rootCmd.AddCommand(NewIPCmd())

	return rootCmd
}

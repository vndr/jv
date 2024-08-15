package main

import (
	"fmt"
	"os"

	"github.com/vndr/jv/cmd"

	"github.com/spf13/viper"
)

var version = "0.1.9"

func main() {
	rootCmd := cmd.NewRootCmd(version)

	// Check if --version or -v flag is passed
	if versionFlag, _ := rootCmd.PersistentFlags().GetBool("version"); versionFlag {
		fmt.Printf("JV Tool v%s\n", version)
		os.Exit(0)
	}

	// Handle the error from viper.BindPFlags
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding flags: %v\n", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

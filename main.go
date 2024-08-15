package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/vndr/jv/cmd"
)

var version = "0.1.16"

func main() {
	rootCmd := cmd.NewRootCmd(version)

	// Bind flags after defining them
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		fmt.Fprintf(os.Stderr, "Error binding flags: %v\n", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

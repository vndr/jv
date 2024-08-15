package cmd_test

import (
	"testing"

	"github.com/vndr/jv/cmd"
)

func TestIPCmd(t *testing.T) {
	ipCmd := cmd.NewIPCmd()

	// Check if the command is properly initialized
	if ipCmd.Use != "ip" {
		t.Errorf("Expected command use 'ip', got %s", ipCmd.Use)
	}

	if ipCmd.Short != "IP related commands" {
		t.Errorf("Expected command short 'IP related commands', got %s", ipCmd.Short)
	}

	// Check if the correct number of subcommands are added
	expectedSubCommands := 2
	if len(ipCmd.Commands()) != expectedSubCommands {
		t.Errorf("Expected %d subcommands, got %d", expectedSubCommands, len(ipCmd.Commands()))
	}

	// Check if the subcommands are correctly added
	subCommandNames := []string{"local", "public"}
	for _, name := range subCommandNames {
		found := false
		for _, cmd := range ipCmd.Commands() {
			if cmd.Use == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected subcommand '%s' to be added", name)
		}
	}
}

func TestIPCmdExecute(t *testing.T) {
	ipCmd := cmd.NewIPCmd()

	// Test the command execution without any subcommand
	ipCmd.SetArgs([]string{}) // No arguments provided
	err := ipCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestIPCmdWithInvalidArgs(t *testing.T) {
	ipCmd := cmd.NewIPCmd()

	// Test the command execution with an invalid argument
	ipCmd.SetArgs([]string{"invalid"})
	err := ipCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error for invalid argument, got nil")
	}
}

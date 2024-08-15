package cmd_test

import (
	"testing"

	"github.com/vndr/jv/cmd"
)

func mockGetLocalIPWithInterfaceCheck() {
	// Mock implementation: do nothing
}

func TestLocalIPCmd(t *testing.T) {
	localCmd := cmd.NewLocalIPCmd(mockGetLocalIPWithInterfaceCheck)

	// Check if the command is properly initialized
	if localCmd.Use != "local" {
		t.Errorf("Expected command use 'local', got %s", localCmd.Use)
	}

	if localCmd.Short != "Get the local IP address" {
		t.Errorf("Expected command short 'Get the local IP address', got %s", localCmd.Short)
	}

	// Execute the command and ensure it runs without error
	err := localCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestLocalIPCmdWithArgs(t *testing.T) {
	localCmd := cmd.NewLocalIPCmd(mockGetLocalIPWithInterfaceCheck)

	// Provide an unexpected argument
	localCmd.SetArgs([]string{"unexpected-arg"})

	err := localCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error for unexpected argument, got nil")
	}
}

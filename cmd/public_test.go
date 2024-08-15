package cmd_test

import (
	"testing"

	"github.com/vndr/jv/cmd"
)

// Mocking the GetPublicIP function
func mockGetPublicIP() {
	// Mock implementation: do nothing or simulate public IP retrieval
}

func TestPublicIPCmd(t *testing.T) {
	publicCmd := cmd.NewPublicIPCmd(mockGetPublicIP)

	// Check if the command is properly initialized
	if publicCmd.Use != "public" {
		t.Errorf("Expected command use 'public', got %s", publicCmd.Use)
	}

	if publicCmd.Short != "Get the public IP address" {
		t.Errorf("Expected command short 'Get the public IP address', got %s", publicCmd.Short)
	}

	// Execute the command and ensure it runs without error
	err := publicCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestPublicIPCmdWithArgs(t *testing.T) {
	publicCmd := cmd.NewPublicIPCmd(mockGetPublicIP)

	// Provide an unexpected argument
	publicCmd.SetArgs([]string{"unexpected-arg"})

	err := publicCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error for unexpected argument, got nil")
	}
}

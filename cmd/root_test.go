package cmd_test

import (
	"bytes"
	"testing"

	"github.com/vndr/jv/cmd"
)

func TestRootCmdInitialization(t *testing.T) {
	rootCmd := cmd.NewRootCmd("v0.1.16")

	// Check if the correct number of subcommands are added
	if len(rootCmd.Commands()) != 2 {
		t.Errorf("Expected 2 subcommands, got %d", len(rootCmd.Commands()))
	}
}

func TestRootCmdInvalidCommand(t *testing.T) {
	rootCmd := cmd.NewRootCmd("v0.1.16")

	// Capture the output to suppress it in the test output
	output := &bytes.Buffer{}
	rootCmd.SetOut(output)
	rootCmd.SetErr(output)

	// Set an invalid command
	rootCmd.SetArgs([]string{"invalid-command"})

	// Execute the command and check if it returns an error
	err := rootCmd.Execute()
	if err == nil {
		t.Fatalf("Expected an error for invalid command")
	}

	// Optionally, you can check the output here if needed
	expectedError := "Error: unknown command \"invalid-command\" for \"jv\"\n"
	if !bytes.Contains(output.Bytes(), []byte(expectedError)) {
		t.Errorf("Expected error output to contain %q, got %q", expectedError, output.String())
	}
}

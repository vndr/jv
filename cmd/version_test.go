package cmd_test

import (
	"bytes"
	"testing"

	"github.com/vndr/jv/cmd"
)

func TestVersionCmd(t *testing.T) {
	version := "v0.1.16"
	versionCmd := cmd.NewVersionCmd(version)

	// Capture the output
	output := &bytes.Buffer{}
	versionCmd.SetOut(output)

	// Execute the command
	err := versionCmd.Execute()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Print the captured output to see what it actually is
	actualOutput := output.String()
	t.Logf("Captured output: %q", actualOutput)

	// Check the output
	expectedOutput := "JV Tool " + version + "\n"
	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %q, got %q", expectedOutput, actualOutput)
	}
}

func TestVersionCmdWithArgs(t *testing.T) {
	version := "v0.1.16"
	versionCmd := cmd.NewVersionCmd(version)

	// Provide an unexpected argument
	versionCmd.SetArgs([]string{"unexpected-arg"})

	// Capture the output to ensure error is printed
	output := &bytes.Buffer{}
	versionCmd.SetOut(output)
	versionCmd.SetErr(output)

	err := versionCmd.Execute()
	if err == nil {
		t.Fatalf("Expected error for unexpected argument, got nil")
	}

	// Print the captured output to see what it actually is
	actualErrorOutput := output.String()
	t.Logf("Captured error output: %q", actualErrorOutput)

	// Check if the correct error message is output
	expectedErrorOutput := "Error: unexpected arguments provided"
	if !bytes.Contains([]byte(actualErrorOutput), []byte(expectedErrorOutput)) {
		t.Errorf("Expected error output to contain %q, got %q", expectedErrorOutput, actualErrorOutput)
	}

	// Optionally, check that the usage information is included
	if !bytes.Contains([]byte(actualErrorOutput), []byte("Usage:")) {
		t.Errorf("Expected usage information in error output, got %q", actualErrorOutput)
	}
}

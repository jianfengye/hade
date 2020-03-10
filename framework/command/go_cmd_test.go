package command

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestGoCommand(t *testing.T) {
	var rootCmdArgs []string
	rootCmd := &cobra.Command{
		Use: "root",
		Run: func(_ *cobra.Command, args []string) { rootCmdArgs = args },
	}
	rootCmd.AddCommand(goCommand)

	output, err := executeCommand(rootCmd, "go", "test")
	if output != "" {
		t.Errorf("Unexpected output: %v", output)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	got := strings.Join(rootCmdArgs, " ")
	expected := "one two"
	if got != expected {
		t.Errorf("rootCmdArgs expected: %q, got: %q", expected, got)
	}
}

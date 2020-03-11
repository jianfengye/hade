package command

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// npm just run local go bin
var npmCommand = &cobra.Command{
	Use:   "npm",
	Short: "run PATH/npm for npm action",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("npm")
		if err != nil {
			log.Fatalln("hade npm: should install npm in your PATH")
		}

		cmd := exec.Command(path, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return nil
	},
}

package command

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// go just run local go bin
var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "build hade itself",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("go")
		if err != nil {
			log.Fatalln("hade go: should install npm in your PATH")
		}

		cmd := exec.Command(path, "build", "-o", "hade", "./")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return nil
	},
}

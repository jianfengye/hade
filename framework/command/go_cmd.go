package command

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// go just run local go bin
var goCommand = &cobra.Command{
	Use:   "go",
	Short: "run PATH/go for go action",
	RunE: func(c *cobra.Command, args []string) error {
		path, err := exec.LookPath("go")
		if err != nil {
			log.Fatalln("hade go: should install go in your PATH")
		}

		cmd := exec.Command(path, args...)
		out, _ := cmd.CombinedOutput()
		println(string(out))
		return nil
	},
}

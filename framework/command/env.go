package command

import (
	"fmt"

	"github.com/jianfengye/hade/framework/contract"
	"github.com/spf13/cobra"
)

// envCommand show current envionment
var envCommand = &cobra.Command{
	Use:   "env",
	Short: "Get current environment",
	Run: func(c *cobra.Command, args []string) {
		container := GetContainer(c.Root())
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		fmt.Println(envService.AppEnv())
	},
}

package console

import (
	"github.com/jianfengye/hade/framework"
	"github.com/jianfengye/hade/framework/command"
	"github.com/jianfengye/hade/framework/provider/app"
	"github.com/jianfengye/hade/framework/provider/config"
	"github.com/jianfengye/hade/framework/provider/env"
	"github.com/jianfengye/hade/framework/provider/log"
	"github.com/spf13/cobra"
)

// RunCommand is command
func RunCommand(container framework.Container) error {
	container.Singleton(&app.HadeAppProvider{})
	container.Singleton(&env.HadeEnvProvider{})
	container.Singleton(&config.HadeConfigProvider{})
	container.Singleton(&log.HadeLogServiceProvider{})

	var rootCmd = &cobra.Command{
		Use:   "hade",
		Short: "main",
		Long:  "show all commands",
	}

	ctx := command.RegiestContainer(container, rootCmd)
	command.AddKernelCommands(rootCmd)
	return rootCmd.ExecuteContext(ctx)
}

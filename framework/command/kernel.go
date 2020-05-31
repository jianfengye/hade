package command

import (
	"context"

	"github.com/jianfengye/hade/framework"
	"github.com/spf13/cobra"
)

type ContainerKey string

const containerKey = ContainerKey("container")

func RegiestContainer(c framework.Container, cmd *cobra.Command) context.Context {
	return context.WithValue(context.Background(), containerKey, c)
}

func GetContainer(cmd *cobra.Command) framework.Container {
	val := cmd.Context().Value(containerKey)
	if val == nil {
		container := framework.NewHadeContainer()
		RegiestContainer(container, cmd)
		return container
	}
	return val.(framework.Container)
}

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(cmd *cobra.Command) {
	cmd.AddCommand(envCommand)
	cmd.AddCommand(serveCommand)
	cmd.AddCommand(goCommand)
	cmd.AddCommand(npmCommand)
	cmd.AddCommand(buildCommand)
}

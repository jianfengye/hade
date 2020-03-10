package command

import (
	"github.com/jianfengye/hade/app/http"
	"github.com/jianfengye/hade/framework/contract"
	"github.com/spf13/cobra"
)

// serveCommand start a app serve
var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "Start app serve",
	RunE: func(c *cobra.Command, args []string) error {
		container := GetContainer(c.Root())
		envService := container.MustMake(contract.EnvKey).(contract.Env)

		r, err := http.RunHttp(container)
		if err != nil {
			return err
		}

		url := envService.AppURL()
		if url == "" {
			url = "localhost:8080"
		}
		return r.Run(url)
	},
}

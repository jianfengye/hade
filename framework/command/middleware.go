package command

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/jianfengye/collection"
	"github.com/jianfengye/hade/framework/contract"
	"github.com/spf13/cobra"
)

// func middlewareInit() {
// 	middlewareCommand.AddCommand(middlewareAllCommand)
// }

// middlewareCommand show all installed middleware
var middlewareCommand = &cobra.Command{
	Use:   "middleware",
	Short: "hade middleware action",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) == 0 {
			c.Help()
		}
		return nil
	},
}

var middlewareAllCommand = &cobra.Command{
	Use:   "all",
	Short: "show all installed middleware",
	RunE: func(c *cobra.Command, args []string) error {
		container := GetContainer(c.Root())
		appService := container.MustMake(contract.AppKey).(contract.App)

		middlewarePath := path.Join(appService.BasePath(), "framework", "middleware")
		// check folder
		files, err := ioutil.ReadDir(middlewarePath)
		if err != nil {
			return err
		}
		for _, f := range files {
			if f.IsDir() {
				fmt.Print(f.Name())
			}
		}
		return nil
	},
}

var middlewareAddCommand = &cobra.Command{
	Use:   "add",
	Short: "add gin middleware to framework",
	RunE: func(c *cobra.Command, args []string) error {
		// step1 : read args
		if len(args) != 1 {
			return errors.New("args lens invalid")
		}
		repo := args[0]
		// step2 : download git to middleware sub directory
		container := GetContainer(c.Root())
		appService := container.MustMake(contract.AppKey).(contract.App)

		middlewarePath := path.Join(appService.BasePath(), "framework", "middleware")
		url := "https://github.com/gin-contrib/" + repo + ".git"
		fmt.Println("download middleware from gin-contrib:")
		fmt.Println(url)
		_, err := git.PlainClone(path.Join(middlewarePath, repo), false, &git.CloneOptions{
			URL:      url,
			Progress: os.Stdout,
		})
		if err != nil {
			return err
		}

		// step4 : remove go.mod and go.sum and .git
		repoFolder := path.Join(middlewarePath, repo)
		fmt.Println("remove " + path.Join(repoFolder, "go.mod"))
		os.Remove(path.Join(repoFolder, "go.mod"))
		fmt.Println("remove " + path.Join(repoFolder, "go.sum"))
		os.Remove(path.Join(repoFolder, "go.sum"))
		fmt.Println("remove " + path.Join(repoFolder, ".git"))
		os.RemoveAll(path.Join(repoFolder, ".git"))

		// step3 : replace key words
		filepath.Walk(repoFolder, func(path string, info os.FileInfo, err error) error {
			fmt.Println("read file:" + path)
			if info.IsDir() {
				return nil
			}

			if filepath.Ext(path) != ".go" {
				return nil
			}

			c, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			isContain := bytes.Contains(c, []byte("github.com/gin-gonic/gin"))
			if isContain {
				fmt.Println("update file:" + path)
				c = bytes.ReplaceAll(c, []byte("github.com/gin-gonic/gin"), []byte("github.com/jianfengye/hade/framework/gin"))
				err = ioutil.WriteFile(path, c, 0644)
				if err != nil {
					return err
				}
			}

			return nil
		})
		return nil
	},
}

var middlewareRemoveCommand = &cobra.Command{
	Use:   "remove",
	Short: "remove gin middleware to framework",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) <= 0 {
			return errors.New("arg is invalid")
		}

		container := GetContainer(c.Root())
		appService := container.MustMake(contract.AppKey).(contract.App)

		middlewarePath := path.Join(appService.BasePath(), "framework", "middleware")

		files, err := ioutil.ReadDir(middlewarePath)
		if err != nil {
			return err
		}

		collection := collection.NewStrCollection(args)

		for _, file := range files {
			if file.IsDir() && collection.Contains(file.Name()) {
				os.RemoveAll(path.Join(middlewarePath, file.Name()))
			}
		}
		return nil
	},
}

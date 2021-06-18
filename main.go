package main

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func main() {
	conf, err := os.UserConfigDir()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	config := filepath.Join(conf, "golias", path.Base(os.Args[0])+".yaml")

	commands := []*cli.Command{
		{
			Name:   "init",
			Action: _init,
		},
		{
			Name: "edit",
			Action: func(context *cli.Context) error {
				cmd := exec.Command(os.Getenv("EDITOR"), config)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				return cmd.Run()
			},
		},
	}

	if Exists(config) {
		buf, err := ioutil.ReadFile(config)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		var subcommands []SubCommand
		err = yaml.Unmarshal(buf, &subcommands)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		for _, command := range subcommands {
			c := &cli.Command{
				Name:  command.Name,
				Usage: command.Usage,
				Action: func(context *cli.Context) error {
					if len(command.Commands) != 0 {
						out, err := pipeline.Output(command.GetCommands()...)
						if err != nil {
							return err
						}
						fmt.Println(string(out))
					} else {
						cmd := exec.Command(command.Command, command.Args...)
						out, _ := cmd.Output()
						if err != nil {
							return err
						}
						fmt.Print(string(out))
					}

					return nil
				},
			}
			commands = append(commands, c)
		}
	}

	app := &cli.App{
		Name:     path.Base(os.Args[0]),
		Usage:    "alias subcommand from file",
		Commands: commands,
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

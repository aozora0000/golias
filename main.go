package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

var (
	version = "local"
	commit  = "none"
	date    = "none"
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
			Usage:  "create config file :" + config,
			Action: _init,
		},
		{
			Name:  "edit",
			Usage: fmt.Sprintf("edit config file %s", os.Getenv("EDITOR")),
			Action: func(context *cli.Context) error {
				cmd := exec.Command(os.Getenv("EDITOR"), config)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				return cmd.Run()
			},
		},
		{
			Name:  "path",
			Usage: "display config path",
			Action: func(context *cli.Context) error {
				fmt.Println(config)
				return nil
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
				Name:   command.Name,
				Usage:  command.Usage,
				Action: Run(command),
			}
			commands = append(commands, c)
		}
	}

	app := &cli.App{
		Name:     path.Base(os.Args[0]),
		Usage:    "alias subcommand from file",
		Commands: commands,
		Version:  version,
		Metadata: map[string]interface{}{
			"commit":     commit,
			"created_at": date,
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

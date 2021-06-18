package main

import (
	"fmt"
	"github.com/mattn/go-pipeline"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

type SubCommand struct {
	Name     string
	Command  string
	Commands []Command
	Args     ListOrString
	Usage    string
}

func (s SubCommand) GetCommands() [][]string {
	var output [][]string
	for _, c := range s.Commands {
		output = append(output, c.Get())
	}
	return output
}

type Command struct {
	Command string
	Args    ListOrString
}

func (c Command) Get() []string {
	output := []string{c.Command}
	for _, arg := range c.Args {
		output = append(output, arg)
	}
	return output
}

func Run(command SubCommand) func(ctx *cli.Context) error {
	return func(context *cli.Context) error {
		if len(command.Commands) != 0 {
			out, err := pipeline.Output(command.GetCommands()...)
			if err != nil {
				return err
			}
			fmt.Println(string(out))
		} else {
			cmd := exec.Command(command.Command, command.Args...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		}

		return nil
	}
}

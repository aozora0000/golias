package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type SubCommand struct {
	Name     string
	Command  string
	Commands []Command
	Args     ListOrString
	Usage    string
}

func (s SubCommand) GetCommand(context *cli.Context) string {
	var output []string
	output = append(output, s.Command)
	for _, c := range s.Args {
		output = append(output, c)
	}
	return transform(context, strings.Join(output, " "))
}

func (s SubCommand) GetCommands(context *cli.Context) string {
	var output []string
	for _, c := range s.Commands {
		output = append(output, c.Get())
	}
	return transform(context, strings.Join(output, " | "))
}

func transform(context *cli.Context, str string) string {
	for _, s := range context.Args().Slice() {
		str = strings.Replace(str, "%s", s, 1)
	}
	return str
}

type Command struct {
	Command string
	Args    ListOrString
}

func (c Command) Get() string {
	output := []string{c.Command}
	for _, arg := range c.Args {
		output = append(output, arg)
	}
	return strings.Join(output, " ")
}

func Run(command SubCommand) func(ctx *cli.Context) error {
	return func(context *cli.Context) error {
		if len(command.Commands) != 0 {
			cmd := exec.Command("/bin/sh", []string{"-c", command.GetCommands(context)}...)
			cmd.Stdin = os.NewFile(uintptr(syscall.Stdin), context.String("input"))
			cmd.Stdout = os.NewFile(uintptr(syscall.Stdout), context.String("output"))
			cmd.Stderr = os.NewFile(uintptr(syscall.Stderr), context.String("error"))

			return cmd.Run()
		} else {
			cmd := exec.Command("/bin/sh", []string{"-c", command.GetCommand(context)}...)
			cmd.Stdin = os.NewFile(uintptr(syscall.Stdin), context.String("input"))
			cmd.Stdout = os.NewFile(uintptr(syscall.Stdout), context.String("output"))
			cmd.Stderr = os.NewFile(uintptr(syscall.Stderr), context.String("error"))
			return cmd.Run()
		}
	}
}

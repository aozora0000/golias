package main

import (
	"fmt"
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
	Envs     map[string]string
	Usage    string
}

func (s SubCommand) GetCommand(context *cli.Context) string {
	var output []string
	output = append(output, s.Command)
	for _, c := range s.Args {
		output = append(output, strings.TrimRight(c, "\r\n"))
	}
	return s.replaceEnvironment(_replaceArguments(context, strings.Join(output, " ")))
}

func (s SubCommand) GetCommands(context *cli.Context) string {
	var output []string
	for _, c := range s.Commands {
		output = append(output, strings.TrimRight(c.Get(), "\r\n"))
	}
	return s.replaceEnvironment(_replaceArguments(context, strings.Join(output, " | ")))
}

func (s SubCommand) replaceEnvironment(str string) string {
	for key, command := range s.Envs {
		cmd := exec.Command("/bin/sh", []string{"-c", command}...)
		b, err := cmd.Output()
		if err != nil {
			fmt.Println(fmt.Errorf("environment error. %s : %s", command, err.Error()).Error())
			os.Exit(1)
		}
		str = strings.Replace(str, fmt.Sprintf("%%%s", key), strings.TrimRight(string(b), "\r\n"), -1)
	}
	return str
}

func _replaceArguments(context *cli.Context, str string) string {
	return str + " " + strings.Join(context.Args().Slice(), " ")
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
			//fmt.Println(cmd.String())
			cmd.Stdin = os.NewFile(uintptr(syscall.Stdin), context.String("input"))
			cmd.Stdout = os.NewFile(uintptr(syscall.Stdout), context.String("output"))
			cmd.Stderr = os.NewFile(uintptr(syscall.Stderr), context.String("error"))

			return cmd.Run()
		} else {
			cmd := exec.Command("/bin/sh", []string{"-c", command.GetCommand(context)}...)
			//fmt.Println(cmd.String())
			cmd.Stdin = os.NewFile(uintptr(syscall.Stdin), context.String("input"))
			cmd.Stdout = os.NewFile(uintptr(syscall.Stdout), context.String("output"))
			cmd.Stderr = os.NewFile(uintptr(syscall.Stderr), context.String("error"))
			return cmd.Run()
		}
	}
}

package main

import (
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func _init(ctx *cli.Context) error {
	conf, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Join(conf, "golias"), 0700)
	if err != nil {
		return err
	}
	configFile := filepath.Join(conf, "golias", ctx.App.Name+".yaml")
	if !Exists(configFile) {
		f, err := os.OpenFile(configFile, os.O_CREATE|os.O_RDWR, 0700)
		if err != nil {
			return err
		}
		defer f.Close()
		example := []SubCommand{
			{
				Name:    "example1",
				Command: "ls",
				Args:    []string{"-la"},
				Usage:   "list file display",
			},
			{
				Name: "example2",
				Commands: []Command{
					{Command: "ls", Args: []string{"-la"}},
					{Command: "wc", Args: []string{"-l"}},
				},
				Usage: "list file count",
			},
		}
		b, err := yaml.Marshal(example)
		if err != nil {
			return err
		}
		f.Write(b)
	}

	return nil
}

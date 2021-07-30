package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Pool struct {
	completed map[string]string
	tasks     map[string]string
	args      []string
}

func NewPool(params map[string]string, envs map[string]string, args []string) Pool {
	if params == nil {
		params = map[string]string{}
	}
	if envs == nil {
		params = map[string]string{}
	}
	if args == nil {
		args = []string{}
	}
	return Pool{completed: mergeHashmap(createArgMap(args), params), tasks: envs}
}

func mergeHashmap(args ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range args {
		for key, val := range m {
			result[key] = val
		}
	}
	return result
}

func createArgMap(args []string) map[string]string {
	result := make(map[string]string)
	for i, value := range args {
		result[strconv.Itoa(i)] = value
	}
	result["ARGS"] = strings.Join(args, " ")
	return result
}

func (p Pool) Init() Pool {
	for {
		p.replaces()
		if p.executeAll() {
			return p
		}
	}
}

func (p Pool) Replace(command string) string {
	for key, value := range p.completed {
		command = strings.ReplaceAll(command, "%"+key, value)
	}
	return command
}

func (p Pool) Get() map[string]string {
	return p.completed
}

func (p Pool) replaces() {
	for key, command := range p.tasks {
		for k, val := range p.completed {
			p.tasks[key] = strings.Replace(command, fmt.Sprintf("%%%s", k), val, -1)
		}
	}
}

func (p Pool) isComplete() bool {
	return len(p.tasks) == 0
}

func (p Pool) getExecutableEnv() map[string]string {
	result := make(map[string]string)
	for key, command := range p.tasks {
		if !strings.Contains(command, "%") {
			result[key] = command
		}
	}
	return result
}

func (p Pool) executeAll() bool {
	for key := range p.getExecutableEnv() {
		p.executeEnv(key)
	}
	return p.isComplete()
}

func (p Pool) executeEnv(key string) {
	cmd := exec.Command("/bin/sh", []string{"-c", p.tasks[key]}...)
	b, err := cmd.Output()
	if err != nil {
		fmt.Println(fmt.Errorf("environment error. %s : %s", p.tasks[key], err.Error()).Error())
		os.Exit(1)
	}
	p.completed[key] = strings.TrimSpace(string(b))
	delete(p.tasks, key)
}

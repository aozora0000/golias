package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Pool struct {
	params    map[string]string
	envs      map[string]string
	env_flags map[string]bool
}

func NewPool(params map[string]string, envs map[string]string) Pool {
	return Pool{params: params, envs: envs, env_flags: createHashBoolMap(envs)}
}

func (p Pool) Init() Pool {
	for key, param := range p.params {
		for k, command := range p.envs {
			if strings.Contains(command, "%"+key) {
				p.replaceEnv(k, strings.ReplaceAll(command, "%"+key, param))
			}
		}
	}
	p.replaceEnvAll()
	return p
}

func mergePool(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
func createHashBoolMap(maps map[string]string) map[string]bool {
	result := make(map[string]bool)
	for k := range maps {
		result[k] = false
	}
	return result
}

func (p Pool) Get() map[string]string {
	return mergePool(p.params, p.envs)
}

func (p Pool) replaceEnvAll() {
	for {
		if p.isAllEnvCompleted() {
			return
		}
		for key, command := range p.envs {
			if p.isAllUnwrapperedEnvCompleted() {

			}
			fmt.Printf("KEY: %s, COMMAND: %s, COMPLETED: %v\n", key, command, p.env_flags[key])
			if p.env_flags[key] == false && !strings.Contains(command, "%") {
				p.executeEnv(key)
				p.env_flags[key] = true
				break
			}
		}
		if p.isAllUnwrapperedEnvCompleted() {

		}
	}
}

func (p Pool) replaceEnv(key string, command string) {
	p.envs[key] = command
}

func (p Pool) executeEnv(key string) {
	cmd := exec.Command("/bin/sh", []string{"-c", p.envs[key]}...)
	fmt.Println(key, cmd.String())
	b, err := cmd.Output()
	if err != nil {
		fmt.Println(fmt.Errorf("environment error. %s : %s", p.envs[key], err.Error()).Error())
		os.Exit(1)
	}
	p.envs[key] = strings.TrimSpace(string(b))
}

func (p Pool) isAllEnvCompleted() bool {
	for _, result := range p.env_flags {
		if !result {
			return false
		}
	}
	return true
}
func (p Pool) isAllUnwrapperedEnvCompleted() bool {
	for key := range p.env_flags {
		if p.env_flags[key] == false && !strings.Contains(p.envs[key], "%") {
			return false
		}
	}
	return true
}

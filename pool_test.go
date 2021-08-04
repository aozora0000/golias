package main

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestPool_Get(t *testing.T) {
	tests := []struct {
		name   string
		params map[string]string
		envs   map[string]string
		args   []string
		want   map[string]string
	}{
		{
			name: "normal test",
			params: map[string]string{
				"TEST_PARAM": "test",
			},
			envs: map[string]string{
				"TEST_ENV":        `echo "%TEST_PARAM"`,
				"TEST_ENV_RESULT": `echo "%TEST_ENV"`,
			},
			args: []string{},
			want: map[string]string{
				"ARGS":            "",
				"TEST_PARAM":      "test",
				"TEST_ENV":        "test",
				"TEST_ENV_RESULT": "test",
			},
		},
		{
			name:   "empty test",
			params: nil,
			envs:   nil,
			want:   map[string]string{"ARGS": ""},
		},
		{
			name:   "argument test",
			params: nil,
			envs:   nil,
			args:   []string{"test"},
			want:   map[string]string{"0": "test", "ARGS": "test"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewPool(test.params, test.envs, test.args).Init().Get()
			fmt.Println(got)
			assert.Equal(t, test.want, got)
			t.Logf("want: %v", test.want)
			t.Logf("got: %v", got)
		})
	}
}

func TestPool_Replace(t *testing.T) {
	tests := []struct {
		name    string
		command string
		params  map[string]string
		envs    map[string]string
		args    []string
		want    string
	}{
		{
			name:    "normal test",
			command: "echo \"%TEST_PARAM\"",
			params: map[string]string{
				"TEST_PARAM": "test",
			},
			envs: map[string]string{
				"TEST_ENV":        `echo "%TEST_PARAM"`,
				"TEST_ENV_RESULT": `echo "%TEST_ENV"`,
			},
			args: []string{},
			want: "echo \"test\"",
		},
		{
			name:    "empty test",
			command: "echo \"%0\"",
			params:  nil,
			envs:    nil,
			args:    []string{"test"},
			want:    "echo \"test\"",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewPool(test.params, test.envs, test.args).Init().Replace(test.command)
			assert.Equal(t, test.want, got)
			t.Logf("want: %v", test.want)
			t.Logf("got: %v", got)
		})
	}
}

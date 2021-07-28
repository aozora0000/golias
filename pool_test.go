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
		want   map[string]string
	}{
		{
			name: "normal test",
			params: map[string]string{
				"TEST_PARAM": "test",
			},
			envs: map[string]string{
				"TEST_ENV":        `echo "TEST_PARAM: %TEST_PARAM"`,
				"TEST_ENV_RESULT": `echo "TEST_ENV: %TEST_ENV"`,
			},
			want: map[string]string{
				"TEST_PARAM":      "test",
				"TEST_ENV":        "TEST_PARAM: test",
				"TEST_ENV_RESULT": "TEST_ENV: test",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NewPool(test.params, test.envs).Init().Get()
			fmt.Println(got)
			assert.Equal(t, test.want, got)
			t.Logf("want: %v", test.want)
			t.Logf("got: %v", got)
		})
	}
}

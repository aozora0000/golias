package main

import "fmt"

type ListOrString []string

func (e *ListOrString) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {

	var aux interface{}
	if err = unmarshal(&aux); err != nil {
		return
	}

	switch raw := aux.(type) {
	case string:
		*e = []string{raw}

	case []interface{}:
		list := make([]string, len(raw))
		for i, r := range raw {
			v, ok := r.(string)
			if !ok {
				return fmt.Errorf("An item in evn cannot be converted to a string: %v", aux)
			}
			list[i] = v
		}
		*e = list

	}
	return
}

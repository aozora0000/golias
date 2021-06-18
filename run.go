package main

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

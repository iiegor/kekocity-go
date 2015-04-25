package cli

type Command struct {
	// The name of the command
	Name string
	// The function to call when this command is invoked
	Action func(context *Context)
}

func (c Command) Names() []string {
	return []string{c.Name}
}

func (c Command) HasName(name string) bool {
	for _, n := range c.Names() {
		if n == name {
			return true
		}
	}
	return false
}

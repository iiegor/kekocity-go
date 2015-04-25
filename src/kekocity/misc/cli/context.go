package cli

type Context struct {
  Cli            *Cli
  Command        Command
}

// Creates a new context.
func NewContext(cli *Cli) *Context {
	return &Context{Cli: cli}
}

package cli

import (
  "fmt"
  "bufio"
  "os"
  "io"
  "strings"
)

type Cli struct {
  // List of commands
  Commands []Command

  // The action to execute
	Action func(context *Context)

  Writer io.Writer
}

func newCli() *Cli {
	return &Cli{
    Writer: os.Stdout,
  }
}

func (c *Cli) parseInput(str string) {
  str = strings.TrimSpace(str)
  context := NewContext(c)

  cm := c.Command(str)
  cm.Action(context)
}

func Listen() {
  var _cli = newCli()
  _cli.Commands = []Command{
    {
      Name: "exit",
      Action: func(c *Context) {
        os.Exit(1)
      },
    },
  }

  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("> ")
    text, _ := reader.ReadString('\n')

    _cli.parseInput(text)
  }
}

func (c *Cli) Command(name string) *Command {
	for _, cm := range c.Commands {
		if cm.HasName(name) {
			return &cm
		}
	}

	return nil
}

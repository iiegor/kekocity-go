package kekocity

import "os"
import "strconv"

type Config struct {
  Buffer int
}

func (u *Config) Emit(str string) {
  os.Stdout.WriteString(str + " (Buffer is " + strconv.Itoa(u.Buffer) + ") \n")
}

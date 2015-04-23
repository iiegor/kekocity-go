package main

import (
	"kekocity"
)

var app = kekocity.Config {
	Buffer: 1200,
}

func main() {
	app.Emit("hello")
}

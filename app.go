package main

import (
	"runtime"
	"kekocity"
)

func main() {
	// Always use the maximum available CPU/Cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	kekocity.Prepare()
	kekocity.Boot()
}

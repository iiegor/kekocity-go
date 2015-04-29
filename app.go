package main

import (
	"runtime"
	"kekocity"
)

func main() {
	// Use the maximum available CPU/Cores
	// GOMAXPROCS is unnecessary when the handlers do not do enough work to justify the time lost communicating between processes.
	runtime.GOMAXPROCS(runtime.NumCPU())

	kekocity.Prepare()
	kekocity.Boot()
}

package main

import (
	"runtime"
	"github.com/fukumone/eto_pirka/routes"
)

func init() {
	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	router := routes.Init()

	router.Run(":3000")
}

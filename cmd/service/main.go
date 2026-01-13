package main

import (
	"fmt"
	"rxplore/internals/daemon"
)

func main() {
	fmt.Println("Starting server!")
	daemon.StartServer()
}

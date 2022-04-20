package main

import (
	"fmt"
	dependencyinjection "learn-go-with-tests/9_dependency-injection"
)

func main() {
	fmt.Println("This is main.")
	fmt.Println("It's not connected to all the child packages because the meat of this is in running the test within each package.")
	fmt.Println("Try running this command:")
	fmt.Println("    go test ./...")

	dependencyinjection.ServeGreeter()
}

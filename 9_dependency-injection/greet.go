package dependencyinjection

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GreetToStdOut() {
	Greet(os.Stdout, "Alexander")
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

/*
	This would work if this were package main
*/
// func main() {
// 	http.ListenAndServe(":5000", http.HandlerFunc(GreeterHandler))
// }

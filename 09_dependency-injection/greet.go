package dependencyinjection

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetToStdOut() {
	Greet(os.Stdout, "Alexander")
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func ServeGreeter() {
	http.ListenAndServe(":5000", http.HandlerFunc(GreeterHandler))
}

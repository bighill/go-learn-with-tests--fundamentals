package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Lex")

	got := buffer.String()
	want := "Hello, Lex"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

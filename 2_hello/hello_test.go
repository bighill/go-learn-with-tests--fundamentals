package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectResult := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to person", func(t *testing.T) {
		got := Hello("lex", "")
		want := "hello lex"
		assertCorrectResult(t, got, want)
	})

	t.Run("say hello to person in Spanish", func(t *testing.T) {
		got := Hello("lex", "Spanish")
		want := "hola lex"
		assertCorrectResult(t, got, want)
	})

	t.Run("say hello to person in French", func(t *testing.T) {
		got := Hello("lex", "French")
		want := "bonjour lex"
		assertCorrectResult(t, got, want)
	})

	t.Run("say hello default", func(t *testing.T) {
		got := Hello("", "")
		want := "hello world"
		assertCorrectResult(t, got, want)
	})
}

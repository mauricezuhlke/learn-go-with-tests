package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Bob", "")
		want := "Hiya, Bob"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying 'Hiya, peeps' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hiya, peeps"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jean Claude", "French")
		want := "Bonjour, Jean Claude"
		assertCorrectMessage(t, got, want)
	})

}

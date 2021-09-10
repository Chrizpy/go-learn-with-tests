package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to person", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello world! when empty argument", func(t *testing.T) {
		got := Hello("")
		want := "Hello world!"

		assertCorrectMessage(t, got, want)
	})
}

package main

import "testing"

func TestHello(t *testing.T){
	
	assertCorrectMessage := func(t testing.TB, got, want string){
		t.Helper()
		
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Herchelle", "")
		want := "Hello, Herchelle"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
                got := Hello("", "")
                want := "Hello, World"

                assertCorrectMessage(t, got, want)
        })

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Herchelle", "french")
		want := "Bonjour, Herchelle"
		assertCorrectMessage(t, got, want)
	})
}

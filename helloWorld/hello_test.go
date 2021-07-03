package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Herchelle")
	want := "Hello, Herchelle"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

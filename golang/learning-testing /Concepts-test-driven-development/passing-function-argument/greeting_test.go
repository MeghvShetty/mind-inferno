package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

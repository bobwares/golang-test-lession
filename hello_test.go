package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bob")
		want := "Hello, Bob"
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("say Hello world when empty string is supplied.", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello, World"
		if actual != expected {
			t.Errorf("got '%s' want '%s'", actual, expected)
		}
	})
}

package main

import "testing"

func Test(t *testing.T) {
	want := "Hello, you"
	got := Hello("you", "")

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}

}

func TestHello(t *testing.T) {
	assertCorrectMessage := func(want, got string, t *testing.T) {
		t.Helper()
		if want != got {
			t.Errorf("Want %s, got %s", want, got)
		}
	}

	t.Run("empty input", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(want, got, t)
	})

	t.Run("empty input", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(want, got, t)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(want, got, t)
	})

}

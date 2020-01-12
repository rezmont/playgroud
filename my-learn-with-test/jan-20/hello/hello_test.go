package main

import "testing"

func Test(t *testing.T) {
	want := "Hello you"
	got := Hello("you")

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}

}

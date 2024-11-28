package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("sagar")
	want := "Hello, sagar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
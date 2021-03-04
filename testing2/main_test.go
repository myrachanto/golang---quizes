package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expected := "Hello, Gopher!\n"
	if got != expected {
		t.Errorf("Did not get the expected results wanted %q, got: %q\n", expected, got)
	}
}
func TestDepart(t *testing.T) {
	got := Depart("Gopher")
	expected := "Goodby, Gopher\n"
	if got != expected {
		t.Errorf("Did not get the expected results wanted %q, got: %q\n", expected, got)
	}
}

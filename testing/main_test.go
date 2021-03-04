package main_test

import (
	"testing"
	
)
func TestAddition(t *testing.T){
		got := 2 + 2
		expected := 4
		if got != 4 {
			t.Errorf("did not get expected results got: '%v' wanted: '%v'", got, expected)
		}
}
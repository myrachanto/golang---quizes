package main

import (
	"sort"
	"testing"

)
func TestSort(t *testing.T){
	ints := []int{9, 20, 4,76,56}
	want := false
 sort.Ints(ints)
 if ints[0]<ints[1]{
	t.Errorf("The results were invalid, got: %t, want: %t.", want, true)
 }

}
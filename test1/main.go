package main

import (
	"fmt"
	"strconv"
)

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

func main() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println ("test")
	fmt.Println(input.GetValue())
}
func (i *NumericInput)GetValue() string{
	return i.input
}
func (i *NumericInput)Add(r rune){
	s:=string(r)
	t, err := strconv.Atoi(s)
	if err != nil {
			return 
	} 
	x := strconv.Itoa(t) 
	// fmt.Println(x)
		i.input = i.input + x
}
// func isNumeric(s string) string {
// 	_, err := strconv.Atoi(s)
// 	if err != nil {
// 			return ""
// 	} 
// 	return s
// }
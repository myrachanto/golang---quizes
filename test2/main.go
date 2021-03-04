package main

import (
	"fmt"
	"math"
	"strconv"

)
var workers int
var resa int
type result struct{
	name string
	weeks int
	days int
} 
func main() {
	c:= make(chan result)
	m := map[string]int { "Robert" : 30, "John" : 475, "Elly" : 1022, "Bob" : 99 }
	results,w := converter(m,c,workers,resa)
	for msg := range results{
		fmt.Println(msg.name + " has worked here for "+ strconv.Itoa(msg.weeks) + " and " + strconv.Itoa(msg.days)+ " days in the company")

	}
	fmt.Println("Workers Count : " + strconv.Itoa(w))
}
func converter(m map[string]int, c chan result, w, resa int) (chan result, int){
	w += 1
 go func(){
	 z :=0
	res := result{} 
		for k,v := range m {
			res.weeks = v/7
			res.name = k
			res.days = v%7

			f := res.weeks
			e := float64(f)
			t := math.Floor(e)
			y := int(t)
			res.weeks = y
			z++

	c <- res
	}
	close(c)
 }()
	return c,w
}

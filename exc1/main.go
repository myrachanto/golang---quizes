package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	c := fanin(boring("joe"), boring("ann"))
	for i :=0; i<10; i++{
		fmt.Println(<-c)
	}
	// jane := boring("booring!")
	// john := boring("booring!")
	// for i := 0; i< 5; i++ {
	// 	fmt.Println(<-jane)
	// 	fmt.Println(<-john)

	// }
	fmt.Println("You are Boring ; am leaving")
}
func fanin(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {c <- <-input2}
	}()
	return c
}

func boring(msg string) <-chan string{
	c := make(chan string)
	go func() {
		for i := 0; i <5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}
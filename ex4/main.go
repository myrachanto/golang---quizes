package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	c := boring("joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("you talk toom much")
			return
		}
	}
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
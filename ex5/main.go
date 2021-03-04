package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	quit := make(chan string)
	c := boring("joe", quit)
	for i := rand.Intn(10); i >= 0; i--{
		fmt.Println(<-c)
		quit <- "bye!"
	}
	select {
	case c <- fmt.Println("blabla"):
		fmt.Println("am working")
	case <- quit:
		fmt.Println("I quit")
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

func boring(msg string,quit <-chan string) <-chan string{
	c := make(chan string)
	go func() {
		for i := 0; i <5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	c<-quit
	return c
}
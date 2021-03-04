package main
import (
	"fmt"
	"time"
)
func main() {
c := producer()
go consumer(c)
time.Sleep(1*time.Millisecond)

}
func producer () <- chan int{
	c := make(chan int)
	go func(){
	fmt.Println("hello i am a producer")
	for i := 0; i<20; i++{
		c<-i*2
	}
	close(c)
	}()
	return c
}
func consumer(in <-chan int){
	fmt.Println("hello am a consumer")
		for msg := range in{
				fmt.Println(msg)
		}
}
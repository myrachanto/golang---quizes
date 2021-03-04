package main
import (
	"fmt"
	"time"
)
func main() {
	c:= make(chan string)
	go count("dogs",c)
	for msg := range c{
	fmt.Println(msg)
	}

}
func count(thing string, c chan string){
		for i := 1; i<= 7; i++{
			c <- thing
			time.Sleep(time.Millisecond * 500)
		}
		close(c)
}
package main
import(
	"fmt"
)

func Greet(name string) string{
	return fmt.Sprintf("Hello, %v!\n", name)
}
func Depart(name string) string{
	return fmt.Sprintf("Goodby, %v\n", name)
}
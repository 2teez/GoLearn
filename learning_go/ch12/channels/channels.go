package channels

import "fmt"

func Greet(ch chan string) {
	msg := <-ch
	ch <- fmt.Sprintf("Thanks for %s", msg)
	ch <- "Hello, " + "Master!"
}

//Channel type is specified goroutine
package main

import (
	"fmt"
)

func receicer(ch <-chan int)  {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func main()  {
	ch := make(chan int)
	go receicer(ch)
	i := 0
	for i < 1000 {
		ch <- i
		i++
	}
}
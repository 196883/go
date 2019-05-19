package main

import (
	"fmt"
)

func runDefer() {
	defer fmt.Println("defer")
	fmt.Println("done")
}

func main() {
	runDefer()
}
//runDefer()
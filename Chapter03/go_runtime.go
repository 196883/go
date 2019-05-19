package main

import (
	"fmt"
	"runtime"
)

func main()  {
	go fmt.Println("Yeah!")
	fmt.Println("NumCPU: %d\n", runtime.NumCPU())
	fmt.Println("NumGproutine: %d\n",runtime.NumGoroutine())
	fmt.Println("Version: %s\n",runtime.Version())
}
package main

import (
	"fmt"
)

func main()  {
	s := []string{"Apple", "Banana", "Cherry"}

	for i := 0; i < len(s); i++ {
		fmt.Println("[%d] => %s\n", i, s[i])
		s = append(s, "Melon")
	}
}
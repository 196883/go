//Slice is dynamic array type in golang
package main

import (
	"fmt"
)

func main()  {
	s := make([]int,10) //int type slice # of contents = 10
	fmt.Println(s)

	var a [10]int //without make
	fmt.Println(a)

	//Substitution
	b := make([]float64,3)
	fmt.Println(b)
	b[0] = 3.14
	fmt.Println(b)
	b[1] = 6.28
	fmt.Println(b)
	fmt.Println(b[0])
	//Run time panic
	//fmt.Println(b[4])

	//See the length of slice
	t := make([]int,8)
	fmt.Println(len(t))

	//See the capacity of slice
	s1 := make([]int,5)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	s2 := make([]int, 5, 10)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	//Append contents for silce
	u := []int{1,2,3}
	u = append(u,4)
	fmt.Println(u)
	u = append(u,5,6,7,)
	fmt.Println(u)

	//Copy slice
	t1 := []int{1,2,3,4,5}
	t2 := []int{10,11}
	n := copy(t1,t2)
	fmt.Println(n)
	fmt.Println(t1)
}
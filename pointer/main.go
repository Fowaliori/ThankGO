package main

import "fmt"

func inc(a *int) {
	*a++
}

func main() {
	var iptr *int
	i := 43
	iptr = &i
	fmt.Println(*iptr)
	fmt.Println(iptr)
	a := 1
	inc(&a)
	fmt.Println(a)
}

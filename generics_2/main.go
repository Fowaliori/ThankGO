package main

import "fmt"

type Slice[T any] []T

type Pair[T any] struct {
	first, second T
}

func (p *Pair[T]) Swap() {
	p.first, p.second = p.second, p.first
}

func (s Slice[T]) Reverse() {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func main() {
	intSlice := Slice[int]{1, 2, 3}
	stringSlice := Slice[string]{"one", "two", "three"}
	fmt.Println(intSlice)
	fmt.Println(stringSlice)
	intSlice.Reverse()
	stringSlice.Reverse()
	fmt.Println(intSlice)
	fmt.Println(stringSlice)
	intPair := Pair[int]{5, 3}
	intPair.Swap()
	fmt.Println(intPair)
}

package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{"one": 1, "two": 2, "thr": 3, "four": 4}
	m2 := maps.Clone(m1)
	fmt.Println(m1, m2)
	maps.DeleteFunc(m1, func(key string, val int) bool {
		return val%2 == 0
	})
	fmt.Println(m1)
	fmt.Println(maps.Equal(m1, m2))
}
